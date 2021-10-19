package config

import "github.com/hashicorp/go-multierror"

// Assert that Operations implements Validate
var _ Validate = (*Operations)(nil)

type (
	Operations struct {
		Total int `mapstructure:"total"`
		Read  int `mapstructure:"read"`
		Write int `mapstructure:"write"`
	}
)

func (o *Operations) Validate() error {
	var result *multierror.Error

	// TODO: Validate table config

	return result.ErrorOrNil()
}
