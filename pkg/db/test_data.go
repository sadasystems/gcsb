package db

import (
	. "github.com/sadasystems/gcsb/pkg/config"
	"reflect"
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

func (f FakeRow) GetValuesString() string {
	return "'" + f.Name + "', '" + f.City + "', '" + f.Genre + "'"
}

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
	},
}}
