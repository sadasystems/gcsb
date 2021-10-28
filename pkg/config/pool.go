package config

import (
	"time"

	"github.com/hashicorp/go-multierror"
)

// Assert that Connection implements Validate
var _ Validate = (*Pool)(nil)

type (
	Pool struct {
		MaxOpened           int           `mapstructure:"max_opened" yaml:"max_opened"`
		MinOpened           int           `mapstructure:"min_opened" yaml:"min_opened"`
		MaxIdle             int           `mapstructure:"max_idle" yaml:"max_idle"`
		WriteSessions       float64       `mapstructure:"write_sessions" yaml:"write_sessions"`
		HealthcheckWorkers  int           `mapstructure:"healthcheck_workers" yaml:"healthcheck_workers"`
		HealthcheckInterval time.Duration `mapstructure:"healthcheck_interval" yaml:"healthcheck_interval"`
		TrackSessionHandles bool          `mapstructure:"track_session_handles" yaml:"track_session_handles"`
	}
)

func (c *Pool) Validate() error {
	var result *multierror.Error

	// TODO: Validate pool config

	return result.ErrorOrNil()
}
