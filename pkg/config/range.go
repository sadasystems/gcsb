package config

import "github.com/hashicorp/go-multierror"

// Assert that Generator implements Validate
var _ Validate = (*Range)(nil)

type (
	Range struct {
		Begin   *interface{} `mapstructure:"begin"`   // Begin for ranges like ranged string & date
		End     *interface{} `mapstructure:"end"`     // End of ranges like ranged string & date
		Length  *int         `mapstructure:"length"`  // Length for generators like string or bytes
		Static  *bool        `mapstructure:"static"`  // Static value indicator for bool generator
		Value   *interface{} `mapstructure:"value"`   // Value for static generation
		Minimum *interface{} `mapstructure:"minimum"` // Minimum for numeric generators
		Maximum *interface{} `mapstructure:"maximum"` // Maximum for numeric generators
	}
)

func (r *Range) Validate() error {
	var result *multierror.Error

	// TODO: Validate Range config

	return result.ErrorOrNil()
}
