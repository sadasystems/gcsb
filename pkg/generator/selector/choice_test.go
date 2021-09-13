package selector

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChoice(t *testing.T) {
	Convey("Choice", t, func() {
		Convey("NewChoice", func() {
			c := NewChoice("test")
			So(c, ShouldNotBeNil)

			_, ok := c.(*choice)
			So(ok, ShouldBeTrue)
		})

		Convey("Item", func() {
			c := NewChoice("test")
			So(c, ShouldNotBeNil)

			v := c.Item()
			vv, ok := v.(string)
			So(ok, ShouldBeTrue)
			So(vv, ShouldEqual, "test")
		})
	})

	Convey("Weighted Choice", t, func() {
		Convey("NewWeightedChoice", func() {
			c := NewWeightedChoice("test", 100)
			So(c, ShouldNotBeNil)

			_, ok := c.(*weightedChoice)
			So(ok, ShouldBeTrue)
		})

		Convey("Item", func() {
			c := NewWeightedChoice("test", 100)
			So(c, ShouldNotBeNil)

			v := c.Item()
			vv, ok := v.(string)
			So(ok, ShouldBeTrue)
			So(vv, ShouldEqual, "test")
		})

		Convey("Weight", func() {
			c := NewWeightedChoice("test", 100)
			So(c, ShouldNotBeNil)

			v := c.Weight()
			So(v, ShouldNotBeNil)
			So(v, ShouldHaveSameTypeAs, uint(1))
		})
	})
}
