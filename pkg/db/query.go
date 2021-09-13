package db

import (
	"github.com/sadasystems/gcsb/pkg/config"
	"strings"
)

func NewInsertQuery (c config.TableConfigTable, row DataRow) string {
	var sb strings.Builder
	sb.WriteString("INSERT " + c.Name + " (")
	var columns []string
	var values []string
	for _, column := range c.Columns {
		columns = append(columns, column.Name)
		values = append(values, row.Get(column))
	}
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(") VALUES (")
	sb.WriteString(strings.Join(values, ", "))
	sb.WriteString(")")
	return sb.String()
}

