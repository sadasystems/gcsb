package db

import (
	"reflect"

	. "github.com/sadasystems/gcsb/pkg/config"
)

type FakeRow struct {
	Name  string
	City  string
	Genre string
}

func (f FakeRow) Get(c TableConfigColumn) string {
	rv := reflect.ValueOf(&f)
	value := reflect.Indirect(rv).FieldByName(c.Name)
	return "'" + value.String() + "'"
}

var FAKE_ROW_CONFIG = FakeRow{Name: "Lil Peep", City: "Allentown", Genre: "Rap"}

var FAKE_ROW_TABLE_CONFIG = TableConfigTable{Name: "Singers", Columns: []TableConfigColumn{
	{
		Name: "Name",
		Type: "STRING(1024)",
	},
	{
		Name: "City",
		Type: "STRING(1024)",
	}, {
		Name: "Genre",
		Type: "STRING(1024)",
	},
}}

var FAKE_DATA_ROW_BUILDER_CONFIG = TableConfigTable{Name: "Singers", Columns: []TableConfigColumn{
	{
		Name: "Name",
		Type: "STRING(1024)",
		Generator: TableConfigGenerator{
			Type:         "combined",
			KeyRange:     TableConfigGeneratorRange{Start: "AAAA", End: "ZZZZ"},
			Length:       10,
			PrefixLength: 4,
		},
	},
	{
		Name: "City",
		Type: "STRING(1024)",
		Generator: TableConfigGenerator{
			Type:     "hexavigesmal",
			KeyRange: TableConfigGeneratorRange{Start: "BBBBB", End: "UUUUU"},
			Length:   5,
		},
	}, {
		Name: "Genre",
		Type: "STRING(1024)",
		Generator: TableConfigGenerator{
			Type:   "string",
			Length: 20,
		},
	},
}}
