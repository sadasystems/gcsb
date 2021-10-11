package config

import "github.com/hashicorp/go-multierror"

// Assert that Generator implements Validate
var _ Validate = (*Range)(nil)

type (
	Range struct{}
)

func (r *Range) Validate() error {
	var result *multierror.Error

	// TODO: Validate pool config

	return result.ErrorOrNil()
}
