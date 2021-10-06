package schema

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTable(t *testing.T) {
	Convey("Table", t, func() {
		// TODO: Test single column and no column error return
		Convey("PointInsertStatement", func() {
			t := NewTable()
			t.SetName("test")

			c1 := NewColumn()
			c1.SetName("foo")
			c2 := NewColumn()
			c2.SetName("bar")

			t.AddColumn(c1)
			t.AddColumn(c2)

			stmt, err := t.PointInsertStatement()
			So(err, ShouldBeNil)
			So(stmt, ShouldEqual, "INSERT INTO test(foo, bar) VALUES(@foo, @bar)")
		})

		// TODO: Test single predicates
		// TODO: Test no predicates returns error
		// TODO: Test that predicates that are not valid column names return error
		Convey("PointReadStatement", func() {
			t := NewTable()
			t.SetName("test")

			c1 := NewColumn()
			c1.SetName("foo")
			c2 := NewColumn()
			c2.SetName("bar")
			c3 := NewColumn()
			c3.SetName("baz")

			t.AddColumn(c1)
			t.AddColumn(c2)
			t.AddColumn(c3)

			stmt, err := t.PointReadStatement("foo", "bar")
			So(err, ShouldBeNil)
			So(stmt, ShouldEqual, "SELECT foo, bar, baz FROM test WHERE foo = @foo AND bar = @bar")
		})
	})
}
