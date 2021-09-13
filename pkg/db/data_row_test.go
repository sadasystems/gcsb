package db

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDataRow(t *testing.T) {
	Convey("DataRowBuilder", t, func() {
		Convey("DataRowBuilder", func() {
			config := FAKE_DATA_ROW_BUILDER_CONFIG
			drb := NewDataRowBuilder(config)

			c0 := config.Columns[0]
			fmt.Println(drb.generators[c0.Name].Next())
			fmt.Println(drb.generators[c0.Name].Next())
			fmt.Println(drb.generators[c0.Name].Next())
			fmt.Println(drb.generators[c0.Name].Next())
			fmt.Println(drb.generators[c0.Name].Next())
			//So(drb.Get(c0), ShouldStartWith, "AAAA")
			//So(drb.Get(c0), ShouldStartWith, "AAAB")
			//So(drb.Get(c0), ShouldStartWith, "AAAC")
			//So(drb.Get(c0), ShouldStartWith, "AAAD")
			//So(drb.Get(c0), ShouldStartWith, "AAAE")
			//
			//c1 := config.Columns[0]
			//So(drb.Get(c1), ShouldStartWith, "AAAA")
			//So(drb.Get(c1), ShouldStartWith, "AAAC")
			//So(drb.Get(c1), ShouldStartWith, "AAAD")
			//So(drb.Get(c1), ShouldStartWith, "AAAB")
			//So(drb.Get(c1), ShouldStartWith, "AAAE")
		})
	})
}
