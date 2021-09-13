package db

import (
	. "github.com/sadasystems/gcsb/pkg/config"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type FakeRow struct {
	Name  string
	City  string
	Genre string
}

func (f FakeRow) Get(s string) string {
	rv := reflect.ValueOf(&f)
	value := reflect.Indirect(rv).FieldByName(s)
	return "'" + value.String() + "'"
}

func TestQuery(t *testing.T) {
	Convey("NewInsertQuery", t, func() {
		fakeRow := FakeRow{Name: "Lil Peep", City: "Allentown", Genre: "Rap"}
		config := TableConfigTable{Name: "Singers", Columns: []TableConfigColumn{
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
		query := NewInsertQuery(config, fakeRow)
		So(query, ShouldEqual, "INSERT Singers (Name, City, Genre) VALUES ('Lil Peep', 'Allentown', 'Rap')")
	})
}
