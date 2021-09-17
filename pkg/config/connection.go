package config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/hashicorp/go-multierror"
	"google.golang.org/api/option"
)

// Assert that Connection implements Validate
var _ Validate = (*Connection)(nil)

type (
	Connection struct {
		Project  string `mapstructure:"project"`
		Instance string `mapstructure:"instance"`
		Database string `mapstructure:"database"`
		NumConns int    `mapstructure:"num_conns"`
		Pool     Pool   `mapstructure:"pool"`
	}

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

func (c *Connection) Validate() error {
	var result *multierror.Error

	if c.Project == "" {
		result = multierror.Append(result, errors.New("connection.project can not be empty"))
	}

	if c.Instance == "" {
		result = multierror.Append(result, errors.New("connection.instance can not be empty"))
	}

	if c.Database == "" {
		result = multierror.Append(result, errors.New("connection.database can not be empty"))
	}

	return result.ErrorOrNil()
}

// Client returns a configured spanner client
func (c *Connection) Client(ctx context.Context) (*spanner.Client, error) {
	client, err := spanner.NewClientWithConfig(ctx, c.DB(), spanner.ClientConfig{
		SessionPoolConfig: spanner.SessionPoolConfig{
			MaxOpened:           uint64(c.Pool.MaxOpened),
			MinOpened:           uint64(c.Pool.MinOpened),
			MaxIdle:             uint64(c.Pool.MaxIdle),
			WriteSessions:       c.Pool.WriteSessions,
			HealthCheckWorkers:  c.Pool.HealthcheckWorkers,
			HealthCheckInterval: c.Pool.HealthcheckInterval,
			TrackSessionHandles: c.Pool.TrackSessionHandles,
		},
	},
		option.WithGRPCConnectionPool(c.NumConns),
	)

	return client, err
}

// DB returns the database DSN
func (c *Connection) DB() string {
	return fmt.Sprintf("projects/%s/instances/%s/database/%s", c.Project, c.Instance, c.Database)
}
