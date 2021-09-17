package config

import (
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
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
		Connection Connection `mapstructure:"connection"`
	}
)

// NewConfig will unmarshal a viper instance into *Config and validate it
func NewConfig(v *viper.Viper) (*Config, error) {
	// Unmarshal the config
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Validate will ensure the configuration is valid
func (c *Config) Validate() error {
	var result *multierror.Error

	// Validate connection block
	errs := c.Connection.Validate()
	if errs != nil {
		result = multierror.Append(result, errs)
	}

	return result.ErrorOrNil()
}
