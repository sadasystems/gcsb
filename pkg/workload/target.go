package workload

import (
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/generator/sample"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"github.com/sadasystems/gcsb/pkg/schema"
)

type Target struct {
	JobType    JobType      // Determines if we are in a 'run' phase or a 'load' phase
	Table      schema.Table // Which table this target points at
	TableName  string       // string name of the table
	Operations int          // Total number of operations to execute against this target

	OperationSelector selector.Selector       // If JobType == JobRun this is used to determine if it should be a read op or a write op
	WriteGenerator    data.GeneratorMap       // Map used for generating row data on inserts
	ReadGenerator     *sample.SampleGenerator // Sample generator for generating point reads
}
