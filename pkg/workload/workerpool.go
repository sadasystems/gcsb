package workload

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload/pool"
)

var (
	// Assert that WorkerPool implements Workload
	_ Workload = (*WorkerPool)(nil)
)

type (
	WorkerPool struct { // Implement Workload
		Context     context.Context
		Config      *config.Config
		Schema      schema.Schema
		initialized bool
		Pool        *pool.Pool
		Jobs        []pool.Job
		wg          sync.WaitGroup
		client      *spanner.Client
	}
)

// NewPoolWorkload initializes a "worker pool" type workload
func NewPoolWorkload(cfg WorkloadConfig) (Workload, error) {
	w := &WorkerPool{
		Context: cfg.Context,
		Config:  cfg.Config,
		Schema:  cfg.Schema,
		Jobs:    make([]pool.Job, 0),
		Pool: pool.NewPool(pool.PoolConfig{
			Workers:        cfg.Config.Threads,
			BufferInput:    true,
			InputBufferLen: 100, // TODO: don't hardcode this
		}),
	}

	return w, nil
}

func (w *WorkerPool) Initialize() error {
	var err error
	w.client, err = w.Config.Client(w.Context)
	if err != nil {
		return err
	}

	w.Pool.Start()

	w.initialized = true

	return nil
}

func (w *WorkerPool) Load(tableName string) error {
	if !w.initialized {
		err := w.Initialize()
		if err != nil {
			return fmt.Errorf("failed to initialize workload: %s", err.Error())
		}

	}

	// Construct generator map for table
	table := w.Schema.GetTable(tableName)
	if table == nil {
		return fmt.Errorf("table '%s' missing from schema", tableName)
	}

	opsPerJob := w.Config.Operations.Total / w.Config.Threads
	for i := 1; i <= w.Config.Threads; i++ {
		// Create a unique generator map instance for each job
		genMap, err := generator.GetDataGeneratorMapForTable(*w.Config, table)
		if err != nil {
			return fmt.Errorf("getting generator map: %s", err.Error())
		}

		// For fun lets grab an insert statement just in case we decide to use dml later
		stmt, err := table.PointInsertStatement()
		if err != nil {
			return fmt.Errorf("getting table write statement: %s", err.Error())
		}

		j := &WorkerPoolLoadJob{
			Context:      w.Context,
			Client:       w.client,
			TableName:    tableName,
			RowCount:     opsPerJob,
			Statement:    stmt,
			GeneratorMap: genMap,
			Batch:        true,
			BatchSize:    500,
			WaitGroup:    &w.wg,
		}

		w.Jobs = append(w.Jobs, j)
		w.wg.Add(1)
		w.Pool.Submit(j)
	}

	w.wg.Wait()

	return nil
}

func (w *WorkerPool) Run() error {
	return nil
}
