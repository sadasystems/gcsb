package db

import (
	"reflect"

	"github.com/sadasystems/gcsb/pkg/config"
)

type FakeRow struct {
	Name  string
	City  string
	Genre string
}

func (f FakeRow) Get(c config.TableConfigColumn) string {
	rv := reflect.ValueOf(&f)
	value := reflect.Indirect(rv).FieldByName(c.Name)
	return "'" + value.String() + "'"
}

func (f FakeRow) GetValuesString() string {
	return "'" + f.Name + "', '" + f.City + "', '" + f.Genre + "'"
}

var FAKE_DATA_ROW_BUILDER_CONFIG = config.TableConfigTable{Name: "Singers", Columns: []config.TableConfigColumn{
	{
		Name: "Name",
		Type: "STRING(1024)",
		Generator: config.TableConfigGenerator{
			Type:         "combined",
			KeyRange:     config.TableConfigGeneratorRange{Start: "AAAA", End: "ZZZZ"},
			Length:       10,
			PrefixLength: 4,
		},
	},
	{
		Name: "City",
		Type: "STRING(1024)",
		Generator: config.TableConfigGenerator{
			Type:     "hexavigesmal",
			KeyRange: config.TableConfigGeneratorRange{Start: "BBBBB", End: "UUUUU"},
			Length:   5,
		},
	},
}}
