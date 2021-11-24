package config

import "github.com/hashicorp/go-multierror"

// Assert that Table implements Validate
var _ Validate = (*Table)(nil)

type (
	Table struct {
		Name       string      `mapstructure:"name"`
		Operations *Operations `mapstructure:"operations" yaml:"operations"`
		Columns    []Column    `mapstructure:"columns"`
	}
)

func (t *Table) Validate() error {
	var result *multierror.Error

	// TODO: Validate table config

	return result.ErrorOrNil()
}

func (t *Table) Column(name string) *Column {
	for _, c := range t.Columns {
		if c.Name == name {
			return &c
		}
	}

	return nil
}
