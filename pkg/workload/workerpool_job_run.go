package workload

import (
	"context"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/generator/operation"
	"github.com/sadasystems/gcsb/pkg/generator/sample"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload/pool"
)

var (
	// Assert that WorkerPoolRunJob implements pool.Job
	_ pool.Job = (*WorkerPoolRunJob)(nil)
)

type (
	WorkerPoolRunJob struct {
		Context           context.Context
		Client            *spanner.Client
		TableName         string
		ReadMap           data.GeneratorMap // Generate data for point reads
		WriteMap          data.GeneratorMap // Generate data for writes
		OperationSelector selector.Selector // Weghted choice selector (read or write)
		WaitGroup         *sync.WaitGroup
		StaleReads        bool // Should we perform stale reads? If not, strong reads
		Staleness         time.Duration
		Operations        int // How many operations to perform
		readfn            func() error
		Table             schema.Table
		ReadGenerator     *sample.SampleGenerator
		cols              []string
	}
)

// TODO: Maybe add a constructor..

func (j *WorkerPoolRunJob) Execute() {
	j.cols = j.Table.ColumnNames()

	if j.readfn == nil {
		if j.StaleReads {
			j.readfn = j.ReadStale
		} else {
			j.readfn = j.ReadStrong
		}
	}

	for i := 0; i <= j.Operations; i++ {
		// Select an operation to perform
		choice := j.OperationSelector.Select()
		op, ok := choice.Item().(operation.Operation)
		if !ok {
			log.Printf("invalid operation from selector '%s', '%T'", choice.Item(), choice.Item())
			continue
		}

		var err error
		switch op {
		case operation.READ:
			err = j.readfn()
		case operation.WRITE:
			err = j.Insert()
		}

		if err != nil {
			log.Printf("error performing operation: %s", err.Error())
		}
	}

	j.WaitGroup.Done()
}

func (j *WorkerPoolRunJob) ReadStale() error {
	ro := j.Client.ReadOnlyTransaction().WithTimestampBound(spanner.ExactStaleness(j.Staleness))

	row, err := ro.ReadRow(j.Context, j.TableName, j.ReadGenerator.Next(), j.cols)
	if err != nil {
		return err
	}

	dst := make([]interface{}, 0, row.Size())
	err = row.Columns(dst...)
	return err
}

func (j *WorkerPoolRunJob) ReadStrong() error {
	// func (t *ReadOnlyTransaction) ReadRow(ctx context.Context, table string, key Key, columns []string) (*Row, error)
	row, err := j.Client.Single().ReadRow(j.Context, j.TableName, j.ReadGenerator.Next(), j.cols)
	if err != nil {
		return err
	}

	dst := make([]interface{}, row.Size())
	err = row.Columns(dst...)
	return err
}

func (j *WorkerPoolRunJob) Insert() error {
	m := make(map[string]interface{}, len(j.WriteMap))
	for k, v := range j.WriteMap {
		m[k] = v.Next()
	}

	_, err := j.Client.Apply(j.Context, []*spanner.Mutation{spanner.InsertMap(j.TableName, m)})
	return err
}
