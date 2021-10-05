package schema

import (
	"context"
	"fmt"

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
