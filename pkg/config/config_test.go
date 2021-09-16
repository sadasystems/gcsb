package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConfig(t *testing.T) {
	Convey("TableConfigTable", t, func() {
		Convey("GetColumnNamesString", func() {
			config := TableConfigTable{Columns: []TableConfigColumn{
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
			columns := config.GetColumnNamesString()
			So(columns, ShouldEqual, "Name, City, Genre")
		})

	})
}
