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
	"github.com/sadasystems/gcsb/pkg/generator/sample"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload/pool"
)

var (
	// Assert that WorkerPool implements Workload
	_ Workload = (*CoreWorkload)(nil)
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
	for _, t := range targets {
		// Fetch table from schema by name
		st := c.Schema.GetTable(t)
		if st == nil {
			return fmt.Errorf("table '%s' missing from information schema", t)
		}

		// Create target
		target := &Target{
			JobType:   pt,
			Table:     st,
			TableName: t,
		}

		c.plan = append(c.plan, target)
	}

	return nil
}

func (c *CoreWorkload) Load(string) error {
	return nil
}

func (c *CoreWorkload) Run(string) error {
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

// glueMetrics is a helper to attach metrics to a job
func (c *CoreWorkload) glueMetrics(j *Job) {
	// Add metrics to job
	j.DataWriteGenerationTimer = c.DataWriteGenerationTimer
	j.DataReadGenerationTimer = c.DataReadGenerationTimer
	j.DataWriteTimer = c.DataWriteTimer
	j.DataWriteMeter = c.DataWriteMeter
	j.DataReadTimer = c.DataReadTimer
	j.DataReadMeter = c.DataReadMeter
}
