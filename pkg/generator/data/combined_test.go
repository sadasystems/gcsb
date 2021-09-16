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
		Convey("Next 5", func() {
			cg, err := NewCombinedGenerator(CombinedGeneratorConfig{
				Min:          0,
				Max:          1000000,
				PrefixLength: 5,
				StringLength: 10,
			})
			So(err, ShouldBeNil)
			So(cg, ShouldNotBeNil)

			So(cg.Next(), ShouldStartWith, "AAAAA")
			So(cg.Next(), ShouldStartWith, "AAAAB")
			So(cg.Next(), ShouldStartWith, "AAAAC")
			So(cg.Next(), ShouldStartWith, "AAAAD")
			So(cg.Next(), ShouldStartWith, "AAAAE")
			i := 0
			for i < 1000 {
				cg.Next()
				i++
			}
			So(cg.Next(), ShouldStartWith, "AABMR")
			So(cg.Next(), ShouldStartWith, "AABMS")
			So(cg.Next(), ShouldStartWith, "AABMT")
		})
	})
}
