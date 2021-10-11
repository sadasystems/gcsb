package config

import "github.com/hashicorp/go-multierror"

// Assert that Generator implements Validate
var _ Validate = (*Generator)(nil)

type (
	Generator struct {
		Type         string `mapstructure:"type"`
		Length       int    `mapstructure:"length"`
		PrefixLength int    `mapstructure:"prefix_length"`
		Range        Range  `mapstructure:"range"`
	}
)

func (g *Generator) Validate() error {
	var result *multierror.Error

	// TODO: Validate pool config

	return result.ErrorOrNil()
}
