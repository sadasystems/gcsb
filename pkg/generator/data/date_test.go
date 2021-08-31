package data

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDateGenerator(t *testing.T) {
	Convey("DateGenerator", t, func() {
		Convey("Nil Source", func() {
			dg, err := NewDateGenerator(DateGeneratorConfig{})

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)
			So(dg.src, ShouldNotBeNil)
		})

		Convey("Random", func() {
			dg, err := NewDateGenerator(DateGeneratorConfig{
				Source: rand.NewSource(time.Now().UnixNano()),
			})

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)
			So(dg.src, ShouldNotBeNil)

			min := civil.DateOf(time.Unix(dg.min, 0))
			max := civil.DateOf(time.Unix(dg.max, 0))

			for i := 0; i < 20; i++ {
				v, ok := dg.Next().(civil.Date)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				So(v.Before(max), ShouldBeTrue)
				So(v.After(min), ShouldBeTrue)
			}
		})

		Convey("Range", func() {
			max := time.Now()
			min := max.Add(-(time.Hour * 24) * 7)
			fmt.Println(min, max)

			dg, err := NewDateGenerator(DateGeneratorConfig{
				Source:  rand.NewSource(time.Now().UnixNano()),
				Range:   true,
				Minimum: min,
				Maximum: max,
			})

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)
			So(dg.src, ShouldNotBeNil)

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
	dg, err := NewDateGenerator(DateGeneratorConfig{
		Source: rand.NewSource(time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dg.Next()
	}
}
