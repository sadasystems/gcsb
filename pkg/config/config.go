package config

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var (
	// Assert that Config implements Validate
	_ Validate = (*Config)(nil)
)

type (
	Validate interface {
		Validate() error
	}
	Config struct {
		Project  string `mapstructure:"project"`
		Instance string `mapstructure:"instance"`
		Database string `mapstructure:"database"`
		NumConns int    `mapstructure:"num_conns"`
		Pool     Pool   `mapstructure:"pool"`
	}
)

// NewConfig will unmarshal a viper instance into *Config and validate it
func NewConfig(v *viper.Viper) (*Config, error) {
	// Bind env vars
	Bind(v)

	// Set Default Values
	SetDefaults(v)

	// Unmarshal the config
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Validate will ensure the configuration is valid for attempting to establish a connection
func (c *Config) Validate() error {
	var result *multierror.Error

	if c.Project == "" {
		result = multierror.Append(result, errors.New("project can not be empty"))
	}

	if c.Instance == "" {
		result = multierror.Append(result, errors.New("instance can not be empty"))
	}

	if c.Database == "" {
		result = multierror.Append(result, errors.New("database can not be empty"))
	}

	// Validate pool block
	errs := c.Pool.Validate()
	if errs != nil {
		result = multierror.Append(result, errs)
	}

	return result.ErrorOrNil()
}

// Client returns a configured spanner client
func (c *Config) Client(ctx context.Context) (*spanner.Client, error) {
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

		// TODO(grpc/grpc-go#1388) using connection pool without WithBlock
		// can cause RPCs to fail randomly. We can delete this after the issue is fixed.
		option.WithGRPCDialOption(grpc.WithBlock()),
	)

	return client, err
}

// DB returns the database DSN
func (c *Config) DB() string {
	return fmt.Sprintf("projects/%s/instances/%s/database/%s", c.Project, c.Instance, c.Database)
}
