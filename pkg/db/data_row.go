package db

import (
	"fmt"
	"strings"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator"
	"github.com/sadasystems/gcsb/pkg/generator/data"
)

type (
	DataRow interface {
		Get(column config.TableConfigColumn) string
		GetValuesString() string
	}

	DataRowBuilder struct {
			config     config.TableConfigTable
		generators map[string]data.Generator
	}
)

func NewDataRowBuilder(c config.TableConfigTable) *DataRowBuilder {
	drb := &DataRowBuilder{
		generators: make(map[string]data.Generator),
		config:     c,
	}
	for _, column := range c.Columns {
		drb.generators[column.Name], _ = generator.GetGenerator(column.Generator)
	}

	return drb
}

func (drb *DataRowBuilder) Get(col config.TableConfigColumn) string {
	var ret string
	gen := drb.generators[col.Name]
	if strings.HasPrefix(col.Type, "STRING") {
		ret = fmt.Sprintf("%v", gen.Next())
		ret = "'" + ret + "'"
	}
	return ret
}

func (drb *DataRowBuilder) GetValuesString() string {
	var values []string
	for _, col := range drb.config.Columns {
		gen := drb.generators[col.Name]
		val := fmt.Sprintf("%v", gen.Next())
		values = append(values, val)
	}
	return strings.Join(values, ", ")
}

