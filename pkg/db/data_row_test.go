package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDataRow(t *testing.T) {
	Convey("DataRowBuilder", t, func() {
		SkipConvey("DataRowBuilder", func() {

			config := FAKE_DATA_ROW_BUILDER_CONFIG
			drb := NewDataRowBuilder(config)

			c0 := config.Columns[0]
			c1 := config.Columns[1]
			c2 := config.Columns[2]

			So(drb.Get(c0), ShouldStartWith, "'AAAA")
			So(drb.Get(c1), ShouldEqual, "'AAAAA'")
			So(len(drb.Get(c2)), ShouldEqual, 22)

			So(drb.Get(c0), ShouldStartWith, "'AAAB")
			So(drb.Get(c1), ShouldEqual, "'AAAAB'")
			So(len(drb.Get(c2)), ShouldEqual, 22)

			So(drb.Get(c0), ShouldStartWith, "'AAAC")
			So(drb.Get(c1), ShouldEqual, "'AAAAC'")
			So(len(drb.Get(c2)), ShouldEqual, 22)

		})
		SkipConvey("GetValuesString", func() {
			config := FAKE_DATA_ROW_BUILDER_CONFIG
			drb := NewDataRowBuilder(config)
			values := drb.GetValuesString()
			So(values, ShouldStartWith, "AAAAlRczqb, AAAAA")
		})
	})
}
