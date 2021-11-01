package config

import (
	"time"

	"github.com/hashicorp/go-multierror"
)

// Assert that Operations implements Validate
var _ Validate = (*Operations)(nil)

type (
	Operations struct {
		Total       int           `mapstructure:"total" yaml:"total"`
		Read        int           `mapstructure:"read" yaml:"read"`
		Write       int           `mapstructure:"write" yaml:"write"`
		SampleSize  float64       `mapstructure:"sample_size" yaml:"sample_size"`
		ReadStale   bool          `mapstructure:"read_stale" yaml:"read_stale"`
		Staleness   time.Duration `mapstructure:"staleness" yaml:"staleness"`
		PartialKeys bool          `mapstructure:"partial_keys" yaml:"partial_keys"`
	}
)

func (o *Operations) Validate() error {
	var result *multierror.Error

	// TODO: Validate table config

	return result.ErrorOrNil()
}
