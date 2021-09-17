package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

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
	PrimaryKey string                `yaml:"primary_key"`
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
	Min          int                       `yaml:"min"`
	Max          int                       `yaml:"max"`
}

type TableConfigGeneratorRange struct {
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

type TableConfigOperations struct {
	Read  uint `yaml:"read"`
	Write uint `yaml:"write"`
}

func NewGCSBConfigFromPath(configPath string) (*GCSBConfig, error) {
	c := &GCSBConfig{}
	err := c.ReadConfig(configPath)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *TableConfigTable) GetColumnNamesString() string {
	var columns []string
	for _, column := range c.Columns {
		columns = append(columns, column.Name)
	}
	return strings.Join(columns, ", ")
}

func (c *TableConfigTable) GetCreateStatement() string {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	sb.WriteString(c.Name)
	sb.WriteString("(")
	var columns []string
	for _, column := range c.Columns {
		columns = append(columns, column.Name+" "+column.Type)
	}
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(") PRIMARY KEY (" + c.PrimaryKey + ")")
	return sb.String()
}

func (c *GCSBConfig) ParentName() string {
	return "projects/" + c.Project + "/instances/" + c.Instance
}

func (c *GCSBConfig) DBName() string {
	return c.ParentName() + "/databases/" + c.Database
}

func (c *GCSBConfig) ReadConfig(configPath string) error {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return err
	}
	return nil
}
func (c *GCSBConfig) GetCreateStatements() []string {
	var columns []string
	for _, table := range c.Tables {
		columns = append(columns, table.GetCreateStatement())
	}
	return columns
}
