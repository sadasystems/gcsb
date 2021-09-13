package db

import (
	"github.com/sadasystems/gcsb/pkg/config"
	"strings"
)

type DataRow interface {
	Get(string) string
}

func NewInsertQuery (c config.TableConfigTable, row DataRow) string {
	var sb strings.Builder
	sb.WriteString("INSERT " + c.Name + " (")
	var columns []string
	var values []string
	for _, column := range c.Columns {
		columns = append(columns, column.Name)
		values = append(values, row.Get(column.Name))
	}
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(") VALUES (")
	sb.WriteString(strings.Join(values, ", "))
	sb.WriteString(")")
	return sb.String()
}

