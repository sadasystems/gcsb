package config

import (
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
)

// Assert that Config implements Validate
var _ Validate = (*Config)(nil)

// Set Configuration Defaults
func SetDefaults(v *viper.Viper) {
	// Default Connection
	v.SetDefault("connection.num_conns", 10)
	// v.SetDefault("")
}

type (
	Validate interface {
		Validate() error
	}
	Config struct {
		Connection Connection `mapstructure:"connection"`
	}
)

func (c *Config) Validate() error {
	var result *multierror.Error

	// Validate connection block
	errs := c.Connection.Validate()
	if errs != nil {
		result = multierror.Append(result, errs)
	}

	return result.ErrorOrNil()
}
