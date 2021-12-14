package workload

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"cloud.google.com/go/spanner"
	"github.com/rcrowley/go-metrics"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/generator/operation"
	"github.com/sadasystems/gcsb/pkg/generator/sample"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload/pool"
)

var (
	// Assert that WorkerPool implements Workload
	_ Workload = (*CoreWorkload)(nil)
)

const (
	defaultBufferLen = 5000
)

type (
	CoreWorkload struct {
		Context         context.Context
		Config          *config.Config
		Schema          schema.Schema
		MetricsRegistry metrics.Registry

		// Internals
		pool   *pool.PipedPool
		wg     sync.WaitGroup
		client *spanner.Client

		DataWriteGenerationTimer metrics.Timer // Used to time data generation
		DataReadGenerationTimer  metrics.Timer // Used to time data geenration
		DataWriteTimer           metrics.Timer // Used to time writes
		DataWriteMeter           metrics.Meter // Used to measure volume of writes
		DataReadTimer            metrics.Timer // Used to time reads
		DataReadMeter            metrics.Meter // Used to measure volume of reads

		// Plans and targets
		plan []*Target // The entire run plan. 1 target per table
	}
)

// NewCoreWorkload initializes a "worker pool" type workload
func NewCoreWorkload(cfg WorkloadConfig) (Workload, error) {
	wl := &CoreWorkload{
		Context:         cfg.Context,
		Config:          cfg.Config,
		Schema:          cfg.Schema,
		MetricsRegistry: cfg.MetricRegistry,
		plan:            make([]*Target, 0),
		pool: pool.NewPipedPool(pool.PipedPoolConfig{
			Workers:         cfg.Config.Threads,
			EnableOutput:    true,
			BufferOutput:    true,
			OutputBufferLen: defaultBufferLen,
			BufferInput:     true,
			InputBufferLen:  defaultBufferLen,
		}),
	}

	// Validat that metrics registry is not nil
	if wl.MetricsRegistry == nil {
		return nil, errors.New("missing metrics registry")
	}

	// Validate that schema is not nil
	if wl.Schema == nil {
		return nil, errors.New("missing schema")
	}

	// Validate that config is not nil
	if wl.Config == nil {
		return nil, errors.New("missing config")
	}

	err := wl.Initialize()
	if err != nil {
		return nil, fmt.Errorf("initializing CoreWorkload: %s", err.Error())
	}

	return wl, nil
}

func (c *CoreWorkload) Initialize() error {
	var err error
	// Initialize spanner client
	c.client, err = c.Config.Client(c.Context)
	if err != nil {
		return err
	}

	// Start the thread pool
	c.pool.Start()

	// Create our job metrics
	c.DataWriteGenerationTimer = metrics.GetOrRegisterTimer("operations.write.data", c.MetricsRegistry) // Used to time data generation
	c.DataReadGenerationTimer = metrics.GetOrRegisterTimer("operations.read.data", c.MetricsRegistry)   // Used to time data geenration
	c.DataWriteTimer = metrics.GetOrRegisterTimer("operations.write.time", c.MetricsRegistry)           // Used to time writes
	c.DataWriteMeter = metrics.GetOrRegisterMeter("operations.write.rate", c.MetricsRegistry)           // Used to measure volume of writes
	c.DataReadTimer = metrics.GetOrRegisterTimer("operations.read.time", c.MetricsRegistry)             // Used to time reads
	c.DataReadMeter = metrics.GetOrRegisterMeter("operations.read.rate", c.MetricsRegistry)             // Used to measure volume of reads

	return nil
}

// Plan will create *Targets for each TargetName
func (c *CoreWorkload) Plan(pt JobType, targets []string) error {
	// search func for looking if targets contains the given string
	// contains := func(s []string, searchterm string) bool {
	// 	i := sort.SearchStrings(s, searchterm)
	// 	return i < len(s) && s[i] == searchterm
	// }

	// Iterate over targets and create Target
	for _, t := range targets {
		// Fetch table from schema by name
		st := c.Schema.GetTable(t)
		if st == nil {
			return fmt.Errorf("table '%s' missing from information schema", t)
		}

		// Create target
		target := &Target{
			Config:                   c.Config,
			Context:                  c.Context,
			Client:                   c.client,
			JobType:                  pt,
			Table:                    st,
			TableName:                t,
			ColumnNames:              st.ColumnNames(),
			DataWriteGenerationTimer: c.DataWriteGenerationTimer,
			DataReadGenerationTimer:  c.DataReadGenerationTimer,
			DataWriteTimer:           c.DataWriteTimer,
			DataWriteMeter:           c.DataWriteMeter,
			DataReadTimer:            c.DataReadTimer,
			DataReadMeter:            c.DataReadMeter,
		}

		// If we are in 'run' context
		if pt == JobRun {
			// Generate an operation selector
			sel, err := c.GetOperationSelector()
			if err != nil {
				return fmt.Errorf("creating operation selector: %s", err.Error())
			}

			target.OperationSelector = sel

			// If our read fraction is > 0,
			// We have faith that the operation selector will not return reads if read fraction is <= 0
			if c.Config.Operations.Read > 0 {
				// Sample the table and create a sample generator
				sg, err := c.GetReadGeneratorMap(target.Table)
				if err != nil {
					return fmt.Errorf("creating sample generator: %s", err.Error())
				}

				target.ReadGenerator = sg
			}
		}

		// Create a generator map for the table
		gm, err := c.GetGeneratorMap(target.Table)
		if err != nil {
			return fmt.Errorf("creating generator map: %s", err.Error())
		}

		target.WriteGenerator = gm

		c.plan = append(c.plan, target)
	}

	return nil
}

func (c *CoreWorkload) Load(string) error {
	return nil
}

// Run will execute a Run phase against the target table.
func (c *CoreWorkload) Run(x string) error {
	// Fetch table from schema
	table := c.Schema.GetTable(x)
	if table == nil {
		return fmt.Errorf("table '%s' missing from schema", x)
	}

	// Check if table is interleaved
	if table.IsInterleaved() {
		// Check if table is apex. If not, return error
		if !table.IsApex() {
			apex := table.GetApex()
			return fmt.Errorf("can only execute run against apex table (try '%s')", apex.Name())
		}
	}

	// Plan our run
	err := c.Plan(JobRun, []string{x})
	if err != nil {
		return fmt.Errorf("planning run: %s", err.Error())
	}

	// Execute our run
	err = c.Execute()
	if err != nil {
		return fmt.Errorf("executing run: %s", err.Error())
	}

	return nil
}

func (c *CoreWorkload) Execute() error {
	////
	// Setup transition threads
	////

	// Create a waitgroup thread. This thread listens to the output of c.pool and decrements
	// the wait group when the job is complete
	waitGroupChan := make(chan pool.Job, defaultBufferLen)
	waitGroupEnd := make(chan bool)
	waitGroupFunc := func() {
		for {
			select {
			case <-waitGroupEnd:
				return
			// case j := <-waitGroupChan:
			case <-waitGroupChan:
				// TODO: type assert j is *Job and check it for errors

				c.wg.Done() // Must release! Otherwise we will deadlock
			}
		}
	}

	c.pool.BindPool(waitGroupChan)
	go waitGroupFunc()

	////
	// Do work. Generate jobs and feed them to the pool
	////
	for _, target := range c.plan {
		// for i:=0;i<=target.Operations;i++ {
		// TODO: figure out operations and batching

		// Get a job from the target
		job := target.NewJob()

		// Submit job to pool
		c.pool.Submit(job)

		// Increment waitgorup
		c.wg.Add(1)
	}

	// Wait for all jobs to flow through the pipeline
	c.wg.Wait()

	return nil
}

func (c *CoreWorkload) Stop() error {
	if c.pool != nil {
		c.pool.Stop()
	}

	return nil
}

// GetReadGeneratorMap will sample rows from the table and create a map structure for creating point reads
func (c *CoreWorkload) GetReadGeneratorMap(t schema.Table) (*sample.SampleGenerator, error) {
	samples, err := c.SampleTable(t)
	if err != nil {
		return nil, fmt.Errorf("sampling table: %s", err.Error())
	}

	return generator.GetReadGeneratorMap(samples, t.PrimaryKeyNames())
}

// SampleTable will return a map[string]interface of values using the tables primary keys
func (c *CoreWorkload) SampleTable(t schema.Table) (map[string]interface{}, error) {
	return generator.SampleTable(c.Config, c.Context, c.client, t)
}

// GetGeneratorMap will return a generator map suitable for creating insert operations against a table
func (c *CoreWorkload) GetGeneratorMap(t schema.Table) (data.GeneratorMap, error) {
	return generator.GetDataGeneratorMapForTable(*c.Config, t)
}

func (c *CoreWorkload) GetOperationSelector() (selector.Selector, error) {
	return operation.NewOperationSelector(c.Config)
}
