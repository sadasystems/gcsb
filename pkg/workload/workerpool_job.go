package workload

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/workload/pool"
)

var (
	// Assert that WorkerPoolLoadJob implements pool.Job
	_ pool.Job = (*WorkerPoolLoadJob)(nil)

	// Assert that WorkerPoolRunJob implements pool.Job
	_ pool.Job = (*WorkerPoolRunJob)(nil)
)

type (
	// WorkerPoolLoadJob is responsible for inserting data into a table
	WorkerPoolLoadJob struct {
		Context      context.Context
		Client       *spanner.Client
		TableName    string
		RowCount     int
		Statement    string
		GeneratorMap data.GeneratorMap
		Batch        bool
		BatchSize    int
		WaitGroup    *sync.WaitGroup
	}

	WorkerPoolRunJob struct{} // Implement pool.Job
)

func (j *WorkerPoolLoadJob) Execute() {
	if j.Batch {
		j.InsertMapBatch()
	} else {
		j.InsertMap()
	}

	j.WaitGroup.Done()
}

func (j *WorkerPoolLoadJob) InsertMapBatch() {
	batch := make([]*spanner.Mutation, 0, j.BatchSize)

	for i := 1; i <= j.RowCount; i++ {
		m := make(map[string]interface{}, len(j.GeneratorMap))
		for k, v := range j.GeneratorMap {
			m[k] = v.Next()
		}

		batch = append(batch, spanner.InsertMap(j.TableName, m))

		if len(batch) == j.BatchSize {
			_, err := j.Client.Apply(j.Context, batch)
			if err != nil {
				log.Printf("error in write transaction: %s", err.Error())
			}

			batch = nil
			batch = make([]*spanner.Mutation, 0, j.BatchSize)
		}
	}

	// Flush the buffer at the end
	if len(batch) > 0 {
		_, err := j.Client.Apply(j.Context, batch)
		if err != nil {
			log.Printf("error in write transaction: %s", err.Error())
		}
	}
}

func (j *WorkerPoolLoadJob) InsertMap() {
	for i := 0; i <= j.RowCount; i++ {
		m := make(map[string]interface{}, len(j.GeneratorMap))
		for k, v := range j.GeneratorMap {
			m[k] = v.Next()
		}

		_, err := j.Client.Apply(j.Context, []*spanner.Mutation{spanner.InsertMap(j.TableName, m)})
		if err != nil {
			log.Printf("error in write transaction: %s", err.Error())
		}
	}
}

func (j *WorkerPoolLoadJob) InsertDML() {
	for i := 0; i <= j.RowCount; i++ {
		stmt := spanner.NewStatement(j.Statement)
		for k, v := range j.GeneratorMap {
			stmt.Params[k] = v.Next()
		}

		_, err := j.Client.ReadWriteTransaction(j.Context, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			_ = txn.Query(ctx, stmt)
			return nil
		})

		if err != nil {
			log.Printf("error in write transaction: %s", err.Error())
		}
	}
}

func (j *WorkerPoolRunJob) Execute() {}
