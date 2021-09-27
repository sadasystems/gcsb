package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGCSBConfig(t *testing.T) {
	Convey("GCSBConfig", t, func() {
		config := GCSBConfig{
			Project:  "A",
			Instance: "B",
			Database: "C",
		}
		Convey("ParentName", func() {
			dbname := config.ParentName()
			So(dbname, ShouldEqual, "projects/A/instances/B")
		})
		Convey("DBName", func() {
			dbname := config.DBName()
			So(dbname, ShouldEqual, "projects/A/instances/B/databases/C")
		})
	})
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
