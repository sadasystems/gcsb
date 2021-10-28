// package workload contains an interface for orchestrating goroutines.
// The idea is that we can experiment with different concurrency models
// without disrupting our call sites.
package workload

import (
	"context"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
)

type (
	Constructor func(WorkloadConfig) (Workload, error)

	Workload interface {
		// Initialize is called once before any operations can proceed
		Initialize() error
		Load(string) error
		Run(string) error
		Stop() error
	}

	WorkloadConfig struct {
		Context context.Context
		Config  *config.Config
		Schema  schema.Schema
	}
)

// GetWorkload adds future support for different concurrency models
func GetWorkloadConstructor(workloadType string) (Constructor, error) {
	switch workloadType {
	default:
		return NewPoolWorkload, nil
	}
}
