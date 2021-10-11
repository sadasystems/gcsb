package config

import "github.com/hashicorp/go-multierror"

// Assert that Table implements Validate
var _ Validate = (*Table)(nil)

type (
	Table struct {
		Name       string     `mapstructure:"name"`
		RowCount   int        `mapstructure:"row_count"`
		Operations Operations `mapstructure:"operations"`
		Columns    []Column   `mapstructure:"columns"`
	}

	Operations struct {
		Read  int `mapstructure:"read"`
		Write int `mapstructure:"write"`
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
