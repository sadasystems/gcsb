package config

type GCSBConfig struct {
	Database string `yaml:"database"`
	Project  string `yaml:"project"`
	Instance string `yaml:"instance"`
	Tables   []TableConfigTable
}

type TableConfigTable struct {
	Name     string `yaml:"name"`
	RowCount int    `yaml:"row_count"`
	Columns  []TableConfigColumn
}

type TableConfigColumn struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Generator TableConfigGenerator
}

type TableConfigGenerator struct {
	Type         string `yaml:"type"`
	Length       int    `yaml:"length"`
	PrefixLength int    `yaml:"prefix_length"`
	Threads      int    `yaml:"threads"`
	Range        struct {
		Start string `yaml:"start"`
		End   string `yaml:"end"`
	}
}
