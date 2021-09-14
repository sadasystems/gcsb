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
	}

	DataRowBuilder struct {
		generators map[string]data.Generator
		config     config.TableConfigTable
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
