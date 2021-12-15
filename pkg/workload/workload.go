// package workload contains an interface for orchestrating goroutines.
// The idea is that we can experiment with different concurrency models
// without disrupting our call sites.
package workload

import (
	"context"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"

	"github.com/rcrowley/go-metrics"
)

type (
	Constructor func(WorkloadConfig) (Workload, error)

	Workload interface {
		Load([]string) error
		Run(string) error
		Stop() error
	}

	WorkloadConfig struct {
		Context        context.Context
		Config         *config.Config
		Schema         schema.Schema
		MetricRegistry metrics.Registry
	}
)

// GetWorkload adds future support for different concurrency models
func GetWorkloadConstructor(workloadType string) (Constructor, error) {
	switch workloadType {
	default:
		// return NewPoolWorkload, nil
		return NewCoreWorkload, nil
	}
}
