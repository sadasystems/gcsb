package db

import (
	"fmt"
	. "github.com/sadasystems/gcsb/pkg/config"
	. "github.com/sadasystems/gcsb/pkg/generator"
	. "github.com/sadasystems/gcsb/pkg/generator/data"
	"strings"
)

type (
	DataRow interface {
		Get(column TableConfigColumn) string
	}

	DataRowBuilder struct {
		generators map[string]Generator
		config     TableConfigTable
		Get        func(TableConfigColumn) string
	}
)

func NewDataRowBuilder(c TableConfigTable) *DataRowBuilder {
	drb := &DataRowBuilder{
		generators: make(map[string]Generator),
		config:     c,
	}
	for _, column := range c.Columns {
		drb.generators[column.Name] = GetGenerator(column.Generator)
	}
	drb.Get = func(col TableConfigColumn) string {
		ret := drb.generators[col.Name].Next()
		str := fmt.Sprintf("%v", ret)
		if strings.HasPrefix(str, "STRING") {
			str = "'" + str + "'"
		}
		return str
	}
	return drb
}

