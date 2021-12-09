package workload

import (
	"context"
	"errors"
	"sync"

	"cloud.google.com/go/spanner"
	"github.com/rcrowley/go-metrics"
	"github.com/sadasystems/gcsb/pkg/config"
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

		pool   *pool.PipedPool
		wg     sync.WaitGroup
		client *spanner.Client
	}
)

// NewCoreWorkload initializes a "worker pool" type workload
func NewCoreWorkload(cfg WorkloadConfig) (Workload, error) {
	wl := &CoreWorkload{
		Context:         cfg.Context,
		Config:          cfg.Config,
		Schema:          cfg.Schema,
		MetricsRegistry: cfg.MetricRegistry,
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

	return wl, nil
}

func (c *CoreWorkload) Load(string) error {
	return nil
}

func (c *CoreWorkload) Run(string) error {
	return nil
}

func (c *CoreWorkload) Stop() error {
	return nil
}
