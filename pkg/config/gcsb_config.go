package config

import "strings"

type GCSBConfig struct {
	Database string             `yaml:"database"`
	Project  string             `yaml:"project"`
	Instance string             `yaml:"instance"`
	Tables   []TableConfigTable `yaml:"tables"`
}

type TableConfigTable struct {
	Name       string                `yaml:"name"`
	RowCount   int                   `yaml:"row_count"`
	Columns    []TableConfigColumn   `yaml:"columns"`
	Operations TableConfigOperations `yaml:"operations"`
}

type TableConfigColumn struct {
	Name      string               `yaml:"name"`
	Type      string               `yaml:"type"`
	Generator TableConfigGenerator `yaml:"generator"`
}

type TableConfigGenerator struct {
	Type         string                    `yaml:"type"`
	Length       int                       `yaml:"length"`
	PrefixLength int                       `yaml:"prefix_length"`
	Threads      int                       `yaml:"threads"`
	KeyRange     TableConfigGeneratorRange `yaml:"key_range"`
	Range        bool                      `yaml:"range"`
	Min          int                       `yaml:"range"`
	Max          int                       `yaml:"range"`
}

type TableConfigGeneratorRange struct {
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

type TableConfigOperations struct {
	Read  uint `yaml:"read"`
	Write uint `yaml:"write"`
}

func (c *TableConfigTable) GetColumnNamesString() string {
	var columns []string
	for _, column := range c.Columns {
		columns = append(columns, column.Name)
	}
	return strings.Join(columns, ", ")
}
