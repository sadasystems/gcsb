package config

import (
	"time"

	"github.com/hashicorp/go-multierror"
)

// Assert that Operations implements Validate
var _ Validate = (*Operations)(nil)

type (
	Operations struct {
		Total      int           `mapstructure:"total"`
		Read       int           `mapstructure:"read"`
		Write      int           `mapstructure:"write"`
		SampleSize float64       `mapstructure:"sample_size"`
		ReadStale  bool          `mapstructure:"read_stale"`
		Staleness  time.Duration `mapstructure:"staleness"`
	}
)

func (o *Operations) Validate() error {
	var result *multierror.Error

	// TODO: Validate table config

	return result.ErrorOrNil()
}
