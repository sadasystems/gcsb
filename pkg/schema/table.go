package schema

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/schema/information"
)

type (
	Table interface {
		SetName(string)
		Name() string
		SetType(string)
		Type() string
		HasParent() bool
		SetParentName(string)
		ParentName() string
		SetParent(Table)
		Parent() Table
		SetSpanenrState(string)
		SpannerState() string

		AddColumn(Column)
		AddIndex(Index)
		Columns() Columns
		ColumnNames() []string

		PrimaryKeys() Columns
		PrimaryKeyNames() []string
		PointInsertStatement() (string, error)
		PointReadStatement(...string) (string, error)
		TableSample(float64) (string, error)
	}

	table struct {
		n            string
		t            string
		p            string
		parent       Table
		spannerState string
		columns      Columns
		indexes      Indexes
	}
)

func NewTable() Table {
	return &table{
		columns: NewColumns(),
		indexes: NewIndexes(),
	}
}

func LoadTable(ctx context.Context, client *spanner.Client, s Schema, t string) error {
	iter := client.Single().Query(ctx, information.GetTableQuery(t))
	defer iter.Stop()
	err := iter.Do(func(row *spanner.Row) error {
		var ti information.Table
		if err := row.ToStruct(&ti); err != nil {
			return err
		}

		tp := NewTableFromSchema(ti)
		s.SetTable(tp)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func LoadTables(ctx context.Context, client *spanner.Client, s Schema) error {
	iter := client.Single().Query(ctx, information.ListTablesQuery())
	defer iter.Stop()

	err := iter.Do(func(row *spanner.Row) error {
		var ti information.Table
		if err := row.ToStruct(&ti); err != nil {
			return err
		}

		tp := NewTableFromSchema(ti)
		s.AddTable(tp)

		return nil
	})

	if err != nil {
		return fmt.Errorf("iterating tables: %s", err.Error())
	}

	return nil
}

func NewTableFromSchema(x information.Table) Table {
	t := NewTable()

	// TODO: I guess check for nil? This isn't safe
	t.SetName(*x.TableName)
	t.SetType(*x.TableType)
	t.SetSpanenrState(*x.SpannerState)

	if x.ParentTableName != nil {
		t.SetParentName(*x.ParentTableName)
	}

	return t
}

func (t *table) SetName(x string) {
	t.n = x
}

func (t *table) Name() string {
	return t.n
}

func (t *table) SetType(x string) {
	t.t = x
}

func (t *table) Type() string {
	return t.t
}

func (t *table) SetParentName(x string) {
	t.p = x
}

func (t *table) ParentName() string {
	return t.p
}

func (t *table) HasParent() bool {
	return t.p != ""
}

func (t *table) SetParent(x Table) {
	t.parent = x
}

func (t *table) Parent() Table {
	return t.parent
}

func (t *table) SetSpanenrState(x string) {
	t.spannerState = x
}

func (t *table) SpannerState() string {
	return t.spannerState
}

func (t *table) AddColumn(x Column) {
	t.columns.AddColumn(x)
}

func (t *table) AddIndex(x Index) {
	t.indexes.AddIndex(x)
}

func (t *table) Columns() Columns {
	return t.columns
}

func (t *table) PointInsertStatement() (string, error) {
	var b strings.Builder

	cols := t.columns.ColumnNames()
	if len(cols) <= 0 {
		return "", errors.New("no columns associated with table")
	}

	fmt.Fprintf(&b, "INSERT INTO %s(%s) VALUES(", t.Name(), strings.Join(cols, ", "))

	fmt.Fprintf(&b, "@%s", cols[0])
	if len(cols) > 1 {
		for _, s := range cols[1:] {
			fmt.Fprintf(&b, ", @%s", s)
		}
	}
	b.WriteString(")")

	return b.String(), nil
}

func (t *table) PointReadStatement(predicates ...string) (string, error) {
	// TODO: Return error if predicate is not a valid column
	if len(predicates) <= 0 {
		return "", errors.New("can not generate point read without predicates")
	}

	var b strings.Builder

	cols := t.columns.ColumnNames()
	if len(cols) <= 0 {
		return "", errors.New("no columns associated with table")
	}

	fmt.Fprintf(&b, "SELECT %s FROM %s WHERE ", strings.Join(cols, ", "), t.Name())
	fmt.Fprintf(&b, "%s = @%s", predicates[0], predicates[0])
	if len(cols) > 1 {
		for _, s := range predicates[1:] {
			fmt.Fprintf(&b, " AND %s = @%s", s, s)
		}
	}

	return b.String(), nil
}

func (t *table) TableSample(x float64) (string, error) {
	pkeys := t.PrimaryKeyNames()

	if len(pkeys) <= 0 {
		return "", errors.New("no primary keys associated with table")
	}

	var b strings.Builder

	fmt.Fprintf(&b, "SELECT %s FROM %s TABLESAMPLE BERNOULLI (%f PERCENT)", strings.Join(pkeys, ", "), t.Name(), x)

	return b.String(), nil
}

func (t *table) PrimaryKeys() Columns {
	return t.columns.PrimaryKeys()
}

func (t *table) PrimaryKeyNames() []string {
	cols := t.columns.PrimaryKeys()
	ret := make([]string, 0, cols.Len())
	for cols.HasNext() {
		col := cols.GetNext()
		ret = append(ret, col.Name())
	}

	return ret
}

func (t *table) ColumnNames() []string {
	ret := make([]string, 0, t.columns.Len())
	for t.columns.HasNext() {
		c := t.columns.GetNext()
		ret = append(ret, c.Name())
	}

	t.columns.ResetIterator()

	return ret
}
