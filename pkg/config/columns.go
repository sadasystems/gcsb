package config

import "github.com/hashicorp/go-multierror"

// Assert that Column implements Validate
var _ Validate = (*Column)(nil)

type (
	Column struct {
		Name      string    `mapstructure:"name"`
		Type      string    `mapstructure:"type"`
		Generator Generator `mapstructure:"generator"`
	}
)

func (c *Column) Validate() error {
	var result *multierror.Error

	// TODO: Validate table config

	return result.ErrorOrNil()
}
