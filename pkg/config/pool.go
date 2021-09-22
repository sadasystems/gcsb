package config

import (
	"time"

	"github.com/hashicorp/go-multierror"
)

// Assert that Connection implements Validate
var _ Validate = (*Pool)(nil)

type (
	Pool struct {
		MaxOpened           int           `mapstructure:"max_opened"`
		MinOpened           int           `mapstructure:"min_opened"`
		MaxIdle             int           `mapstructure:"max_idle"`
		WriteSessions       float64       `mapstructure:"write_sessions"`
		HealthcheckWorkers  int           `mapstructure:"healthcheck_workers"`
		HealthcheckInterval time.Duration `mapstructure:"healthcheck_interval"`
		TrackSessionHandles bool          `mapstructure:"track_session_handles"`
	}
)

func (c *Pool) Validate() error {
	var result *multierror.Error

	// TODO: Validate pool config

	return result.ErrorOrNil()
}
