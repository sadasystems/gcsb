package data

import (
	"testing"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCommitTimestampGenerator(t *testing.T) {
	Convey("CommitTimestampGenerator", t, func() {
		g, err := NewCommitTimestampGenerator(nil)
		So(g, ShouldNotBeNil)
		So(err, ShouldBeNil)

		Convey("Next", func() {
			v := g.Next()
			So(v, ShouldNotBeNil)
			So(v, ShouldHaveSameTypeAs, spanner.CommitTimestamp)
		})

		Convey("TypeBase", func() {
			So(g.Type(), ShouldHaveSameTypeAs, spansql.Timestamp)
		})
	})
}
