package db

import (
	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/operation"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"strings"
)

type (
	QueryBuilder struct {
		config         config.TableConfigTable
		selector       selector.Selector
		dataRowBuilder DataRowBuilder
	}
)

func NewQueryBuilder(c config.TableConfigTable) *QueryBuilder {
	qb := &QueryBuilder{
		config:         c,
		selector:       operation.NewOperationGenerator(&c.Operations),
		dataRowBuilder: *NewDataRowBuilder(c),
	}
	return qb
}
func (qb *QueryBuilder) NewInsertQuery() string {
	var sb strings.Builder
	sb.WriteString("INSERT " + qb.config.Name + " (")
	sb.WriteString(qb.config.GetColumnNamesString())
	sb.WriteString(") VALUES (")
	sb.WriteString(qb.dataRowBuilder.GetValuesString())
	sb.WriteString(")")
	return sb.String()
}

func (qb *QueryBuilder) NewReadQuery() string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	sb.WriteString(qb.config.GetColumnNamesString())
	sb.WriteString(" FROM ")
	sb.WriteString(qb.config.Name)
	return sb.String()
}

func (qb *QueryBuilder) Next() spanner.Statement {
	op := qb.selector.Select().Item()
	var sql string
	if op == 1 {
		sql = qb.NewReadQuery()
	} else if op == 2 {
		sql = qb.NewInsertQuery()
	}
	return spanner.Statement{SQL: sql}
}

//func NewInsertQuery(c config.TableConfigTable, row DataRow) string {
//	var sb strings.Builder
//	sb.WriteString("INSERT " + c.Name + " (")
//	sb.WriteString(c.GetColumnNamesString())
//	sb.WriteString(") VALUES (")
//	sb.WriteString(row.GetValuesString())
//	sb.WriteString(")")
//	return sb.String()
//}
//
//func NewReadQuery(c config.TableConfigTable) string {
//	var sb strings.Builder
//	sb.WriteString("SELECT ")
//	sb.WriteString(c.GetColumnNamesString())
//	sb.WriteString(" FROM ")
//	sb.WriteString(c.Name)
//	return sb.String()
//}
