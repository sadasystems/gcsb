package schema

import (
	"testing"

	"cloud.google.com/go/spanner/spansql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestColumn(t *testing.T) {
	Convey("Column", t, func() {
		Convey("Type", func() {
			Convey("BOOL", func() {
				c := NewColumn()
				c.SetSpannerType("BOOL")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.Bool)
			})

			Convey("String", func() {
				c := NewColumn()
				c.SetSpannerType("STRING")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.String)
			})

			Convey("String(1024)", func() {
				c := NewColumn()
				c.SetSpannerType("STRING(1024)")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.String)
				So(x.Len, ShouldEqual, 1024)
			})

			Convey("String(MAX)", func() {
				c := NewColumn()
				c.SetSpannerType("STRING(MAX)")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.String)
				So(x.Len, ShouldEqual, spansql.MaxLen)
			})

			Convey("INT64", func() {
				c := NewColumn()
				c.SetSpannerType("INT64")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.Int64)
			})

			Convey("FLOAT64", func() {
				c := NewColumn()
				c.SetSpannerType("FLOAT64")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.Float64)
			})

			Convey("BYTES", func() {
				c := NewColumn()
				c.SetSpannerType("BYTES")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.Bytes)
			})

			Convey("TIMESTAMP", func() {
				c := NewColumn()
				c.SetSpannerType("TIMESTAMP")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.Timestamp)
			})

			Convey("DATE", func() {
				c := NewColumn()
				c.SetSpannerType("DATE")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeFalse)
				So(x.Base, ShouldEqual, spansql.Date)
			})

			Convey("ARRAY<STRING>", func() {
				c := NewColumn()
				c.SetSpannerType("ARRAY<STRING>")

				x := c.Type()
				So(x, ShouldNotBeNil)
				So(x.Array, ShouldBeTrue)
				So(x.Base, ShouldEqual, spansql.String)
			})

			Convey("UNKNOWN", func() {
				c := NewColumn()
				c.SetSpannerType("UNKNOWN")

				So(func() { c.Type() }, ShouldPanic)
			})
		})
	})
}
