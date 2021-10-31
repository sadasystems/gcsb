package workload

import (
	"context"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/rcrowley/go-metrics"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/generator/operation"
	"github.com/sadasystems/gcsb/pkg/generator/sample"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload/pool"
	"google.golang.org/grpc/codes"
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
		MetricsRegistry   metrics.Registry
		initialized       bool

		// Timers
		mOps        metrics.Counter
		mErrs       metrics.Counter
		mReadOps    metrics.Counter
		mWriteOps   metrics.Counter
		mReadTimer  metrics.Timer
		mWriteTimer metrics.Timer
		mReadGen    metrics.Timer
		mWriteGen   metrics.Timer
	}
)

// TODO: Maybe add a constructor..

func (j *WorkerPoolRunJob) Initialize() error {
	if j.initialized {
		return nil
	}

	// Get table column names
	j.cols = j.Table.ColumnNames()

	// Set read method based on config
	if j.readfn == nil {
		if j.StaleReads {
			j.readfn = j.ReadStale
		} else {
			j.readfn = j.ReadStrong
		}
	}

	// initialize timers
	j.mOps = metrics.GetOrRegisterCounter("operations", j.MetricsRegistry)
	j.mErrs = metrics.GetOrRegisterCounter("errors", j.MetricsRegistry)
	j.mReadOps = metrics.GetOrRegisterCounter("operations.read.count", j.MetricsRegistry)
	j.mWriteOps = metrics.GetOrRegisterCounter("operations.write.count", j.MetricsRegistry)
	j.mReadTimer = metrics.GetOrRegisterTimer("operations.read.time", j.MetricsRegistry)
	j.mWriteTimer = metrics.GetOrRegisterTimer("operations.write.time", j.MetricsRegistry)
	j.mReadGen = metrics.GetOrRegisterTimer("operations.read.data", j.MetricsRegistry)
	j.mWriteGen = metrics.GetOrRegisterTimer("operations.write.data", j.MetricsRegistry)

	j.initialized = true

	return nil
}

func (j *WorkerPoolRunJob) Execute() {
	err := j.Initialize()
	if err != nil {
		log.Printf("error initializing run job: %s", err.Error())
	}

	for i := 0; i <= j.Operations; i++ {
		// Select an operation to perform
		choice := j.OperationSelector.Select()
		op, ok := choice.Item().(operation.Operation)
		if !ok {
			log.Printf("invalid operation from selector '%s', '%T'", choice.Item(), choice.Item())
			continue
		}

		j.mOps.Inc(1)

		var err error
		switch op {
		case operation.READ:
			j.mReadOps.Inc(1)
			err = j.readfn()
			j.mWriteOps.Inc(1)
		case operation.WRITE:
			err = j.Insert()
		}

		if err != nil {
			if spanner.ErrCode(err) == codes.Canceled {
				log.Println("context canceled")
				break
			}

			j.mErrs.Inc(1)
			log.Printf("error performing operation: %s", err.Error())
		}
	}

	j.WaitGroup.Done()
}

func (j *WorkerPoolRunJob) ReadStale() error {
	ro := j.Client.ReadOnlyTransaction().WithTimestampBound(spanner.ExactStaleness(j.Staleness))

	var rData spanner.Key
	j.mReadGen.Time(func() {
		rData = j.ReadGenerator.Next()
	})

	// var row *spanner.Row
	var err error
	j.mReadTimer.Time(func() {
		_, err = ro.ReadRow(j.Context, j.TableName, rData, j.cols)
	})
	return err
}

func (j *WorkerPoolRunJob) ReadStrong() error {
	var rData spanner.Key
	j.mReadGen.Time(func() {
		rData = j.ReadGenerator.Next()
	})

	var err error
	j.mReadTimer.Time(func() {
		_, err = j.Client.Single().ReadRow(j.Context, j.TableName, rData, j.cols)
	})
	return err
}

func (j *WorkerPoolRunJob) Insert() error {
	m := make(map[string]interface{}, len(j.WriteMap))
	j.mWriteGen.Time(func() {
		for k, v := range j.WriteMap {
			m[k] = v.Next()
		}
	})

	var err error
	j.mWriteTimer.Time(func() {
		_, err = j.Client.Apply(j.Context, []*spanner.Mutation{spanner.InsertMap(j.TableName, m)})
	})

	return err
}
