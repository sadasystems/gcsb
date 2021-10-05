package schema

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema/information"
)

type (
	Schema interface {
		SetTable(Table) // TODO: Remove this
		Table() Table
		AddTable(Table)
	}

	schema struct {
		table  Table // TODO: Remove this
		tables Tables
	}
)

func NewSchema() Schema {
	return &schema{
		tables: NewTables(),
	}
}

func LoadSchema(ctx context.Context, cfg *config.Config) (Schema, error) {
	client, err := cfg.Client(ctx)
	if err != nil {
		return nil, err
	}

	s := NewSchema()

	iter := client.Single().Query(ctx, information.ListTablesQuery())
	defer iter.Stop()

	err = iter.Do(func(row *spanner.Row) error {
		var ti information.Table
		if err := row.ToStruct(&ti); err != nil {
			return err
		}

		tp := NewTableFromSchema(ti)
		s.AddTable(tp)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("iterating tables: %s", err.Error())
	}

	return s, nil
}

// TODO: Get rid of this in milestone 2 (multi-table)
func LoadSingleTableSchema(ctx context.Context, cfg *config.Config, t string) (Schema, error) {
	client, err := cfg.Client(ctx)
	if err != nil {
		return nil, err
	}

	s := NewSchema()

	iter := client.Single().Query(ctx, information.GetTableQuery(t))
	defer iter.Stop()

	err = iter.Do(func(row *spanner.Row) error {
		var ti information.Table
		if err := row.ToStruct(&ti); err != nil {
			return err
		}

		tp := NewTableFromSchema(ti)
		s.SetTable(tp)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *schema) SetTable(x Table) {
	s.table = x
}

func (s *schema) Table() Table {
	return s.table
}

func (s *schema) AddTable(x Table) {
	s.tables.AddTable(x)
}
