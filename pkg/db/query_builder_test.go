package db

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQuery(t *testing.T) {
	Convey("NewInsertQuery", t, func() {
		config := FAKE_DATA_ROW_BUILDER_CONFIG
		qb := NewQueryBuilder(config)
		query := qb.NewInsertQuery()
		So(query, ShouldEqual, "INSERT Singers (Name, City) VALUES (AAAAlRczqb, AAAAA)")
	})
	Convey("NewReadQuery", t, func() {
		config := FAKE_DATA_ROW_BUILDER_CONFIG
		qb := NewQueryBuilder(config)
		query := qb.NewReadQuery()
		So(query, ShouldEqual, "SELECT Name, City FROM Singers")
	})
}
