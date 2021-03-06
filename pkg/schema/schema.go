package schema

import (
	"context"
	"fmt"

	"github.com/sadasystems/gcsb/pkg/config"
)

type (
	Schema interface {
		SetTable(Table) // TODO: Remove this
		Table() Table
		AddTable(Table)
		Tables() Tables
		GetTable(string) Table
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

	// Load Tables
	err = LoadTables(ctx, client, s)
	if err != nil {
		return nil, err
	}

	// Load Columns
	ts := s.Tables()
	for ts.HasNext() {
		t := ts.GetNext()

		// Load Columns
		err := LoadColumns(ctx, client, t)
		if err != nil {
			return nil, fmt.Errorf("loading columns for table '%s': %s", t.Name(), err.Error())
		}

		// Load Indexes
		// Load Indexes
		err = LoadIndexes(ctx, client, t)
		if err != nil {
			return nil, fmt.Errorf("loading indexes for table '%s': %s", t.Name(), err.Error())
		}
	}

	// reset iterator
	ts.ResetIterator()

	return s, nil
}

// TODO: Get rid of this in milestone 2 (multi-table)
func LoadSingleTableSchema(ctx context.Context, cfg *config.Config, t string) (Schema, error) {
	client, err := cfg.Client(ctx)
	if err != nil {
		return nil, err
	}

	s := NewSchema()

	// Load Table
	err = LoadTable(ctx, client, s, t)
	if err != nil {
		return nil, err
	}

	tab := s.Table()

	// Load Columns
	err = LoadColumns(ctx, client, tab)
	if err != nil {
		return nil, fmt.Errorf("loading columns for table '%s': %s", tab.Name(), err.Error())
	}

	// Load Indexes
	err = LoadIndexes(ctx, client, tab)
	if err != nil {
		return nil, fmt.Errorf("loading indexes for table '%s': %s", tab.Name(), err.Error())
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

func (s *schema) Tables() Tables {
	return s.tables
}

func (s *schema) GetTable(x string) Table {
	defer s.tables.ResetIterator()
	for s.tables.HasNext() {
		t := s.tables.GetNext()
		if t.Name() == x {
			return t
		}
	}

	return nil
}
