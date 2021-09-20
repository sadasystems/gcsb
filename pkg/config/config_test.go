package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConfig(t *testing.T) {
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
		config := TableConfigTable{Name: "Singers", PrimaryKey: "Name", Columns: []TableConfigColumn{
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
		Convey("GetColumnNamesString", func() {
			columns := config.GetColumnNamesString()
			So(columns, ShouldEqual, "Name, City, Genre")
		})
		Convey("GetCreateStatement", func() {
			columns := config.GetCreateStatement()
			So(columns, ShouldEqual, "CREATE TABLE Singers(Name STRING(1024), City STRING(1024), Genre STRING(1024)) PRIMARY KEY (Name)")
		})

	})
}
