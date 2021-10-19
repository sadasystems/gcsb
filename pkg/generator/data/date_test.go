package data

import (
	"testing"
	"time"

	"cloud.google.com/go/civil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDateGenerator(t *testing.T) {
	Convey("DateGenerator", t, func() {
		Convey("Nil Source", func() {
			dg, err := NewDateGenerator(NewConfig())

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)

		})

		Convey("Random", func() {
			dg, err := NewDateGenerator(NewConfig())

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)

			tdg, ok := dg.(*DateGenerator)
			So(ok, ShouldBeTrue)
			So(tdg, ShouldNotBeNil)
			min := civil.DateOf(time.Unix(tdg.min, 0))
			max := civil.DateOf(time.Unix(tdg.max, 0))

			for i := 0; i < 20; i++ {
				v, ok := dg.Next().(civil.Date)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				So(v.Before(max), ShouldBeTrue)
				So(v.After(min), ShouldBeTrue)
			}
		})

		Convey("Range", func() {
			cfg := NewConfig()

			max := time.Now()
			min := max.Add(-(time.Hour * 24) * 7)

			cfg.SetRange(true)
			cfg.SetMinimum(min)
			cfg.SetMaximum(max)
			dg, err := NewDateGenerator(cfg)

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)

			dMin := civil.DateOf(min).AddDays(-1)
			dMax := civil.DateOf(max).AddDays(1)

			for i := 0; i < 20; i++ {
				v, ok := dg.Next().(civil.Date)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				So(v.Before(dMax), ShouldBeTrue)
				So(v.After(dMin), ShouldBeTrue)
			}
		})
	})
}

func BenchmarkDateGenerator(b *testing.B) {
	dg, err := NewDateGenerator(NewConfig())
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dg.Next()
	}
}
