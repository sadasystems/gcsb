package db

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQuery(t *testing.T) {
	Convey("NewInsertQuery", t, func() {
		fakeRow := FAKE_ROW_CONFIG
		config := FAKE_ROW_TABLE_CONFIG
		query := NewInsertQuery(config, fakeRow)
		So(query, ShouldEqual, "INSERT Singers (Name, City, Genre) VALUES ('Lil Peep', 'Allentown', 'Rap')")
	})
}
