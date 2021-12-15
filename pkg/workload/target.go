package workload

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/rcrowley/go-metrics"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/generator/sample"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"github.com/sadasystems/gcsb/pkg/schema"
)

type Target struct {
	Config                   *config.Config
	Context                  context.Context
	Client                   *spanner.Client
	JobType                  JobType                 // Determines if we are in a 'run' phase or a 'load' phase
	Table                    schema.Table            // Which table this target points at
	TableName                string                  // string name of the table
	Operations               int                     // Total number of operations to execute against this target
	ColumnNames              []string                // Col names for reads
	OperationSelector        selector.Selector       // If JobType == JobRun this is used to determine if it should be a read op or a write op
	WriteGenerator           data.GeneratorMap       // Map used for generating row data on inserts
	ReadGenerator            *sample.SampleGenerator // Sample generator for generating point reads
	DataWriteGenerationTimer metrics.Timer           // Used to time data generation
	DataReadGenerationTimer  metrics.Timer           // Used to time data geenration
	DataWriteTimer           metrics.Timer           // Used to time writes
	DataWriteMeter           metrics.Meter           // Used to measure volume of writes
	DataReadTimer            metrics.Timer           // Used to time reads
	DataReadMeter            metrics.Meter           // Used to measure volume of reads
}

func (t *Target) NewJob() *Job {
	return &Job{
		JobType:                  t.JobType,
		Context:                  t.Context,
		Client:                   t.Client,
		Table:                    t.TableName,
		Columns:                  t.ColumnNames,
		StaleReads:               t.Config.Operations.ReadStale,
		Staleness:                t.Config.Operations.Staleness,
		Batched:                  t.Config.Batch,
		BatchSize:                t.Config.BatchSize,
		OperationSelector:        t.OperationSelector,
		WriteGenerator:           t.WriteGenerator,
		ReadGenerator:            t.ReadGenerator,
		DataWriteGenerationTimer: t.DataWriteGenerationTimer,
		DataReadGenerationTimer:  t.DataReadGenerationTimer,
		DataWriteTimer:           t.DataWriteTimer,
		DataWriteMeter:           t.DataWriteMeter,
		DataReadTimer:            t.DataReadTimer,
		DataReadMeter:            t.DataReadMeter,
	}
}

func FindTargetByName(plan []*Target, name string) *Target {
	for _, t := range plan {
		if t.TableName == name {
			return t
		}
	}

	return nil
}
