package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinedGenerator(t *testing.T) {
	Convey("CombinedGenerator", t, func() {
		Convey("Next", func() {
			cg, err := NewCombinedGenerator(CombinedGeneratorConfig{
				Min:          0,
				Max:          1000000,
				PrefixLength: 8,
				StringLength: 64,
			})
			So(err, ShouldBeNil)
			So(cg, ShouldNotBeNil)

			So(cg.Next(), ShouldStartWith, "AAAAAAAA")
			So(cg.Next(), ShouldStartWith, "AAAAAAAB")
			So(cg.Next(), ShouldStartWith, "AAAAAAAC")
			So(cg.Next(), ShouldStartWith, "AAAAAAAD")
			So(cg.Next(), ShouldStartWith, "AAAAAAAE")
			i := 0
			for i < 1000 {
				cg.Next()
				i++
			}
			So(cg.Next(), ShouldStartWith, "AAAAABMR")
			So(cg.Next(), ShouldStartWith, "AAAAABMS")
			So(cg.Next(), ShouldStartWith, "AAAAABMT")
		})
	})
}
