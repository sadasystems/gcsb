package data

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimestampGenerator(t *testing.T) {
	Convey("TimestampGenerator", t, func() {
		Convey("Nil Source", func() {
			dg, err := NewTimestampGenerator(TimestampGeneratorConfig{})

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)
			So(dg.src, ShouldNotBeNil)
		})

		Convey("Random", func() {
			dg, err := NewTimestampGenerator(TimestampGeneratorConfig{
				Source: rand.NewSource(time.Now().UnixNano()),
			})

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)
			So(dg.src, ShouldNotBeNil)

			// min := civil.DateOf(time.Unix(dg.min, 0))
			// max := civil.DateOf(time.Unix(dg.max, 0))

			for i := 0; i < 20; i++ {
				v, ok := dg.Next().(time.Time)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				// So(v.Before(max), ShouldBeTrue)
				// So(v.After(min), ShouldBeTrue)
			}
		})

		Convey("Range", func() {
			max := time.Now()
			min := max.Add(-(time.Hour * 24) * 7)

			dg, err := NewTimestampGenerator(TimestampGeneratorConfig{
				Source:  rand.NewSource(time.Now().UnixNano()),
				Range:   true,
				Minimum: min,
				Maximum: max,
			})

			So(err, ShouldBeNil)
			So(dg, ShouldNotBeNil)
			So(dg.src, ShouldNotBeNil)

			// dMin := civil.DateOf(min).AddDays(-1)
			// dMax := civil.DateOf(max).AddDays(1)

			for i := 0; i < 20; i++ {
				v, ok := dg.Next().(time.Time)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				// So(v.Before(dMax), ShouldBeTrue)
				// So(v.After(dMin), ShouldBeTrue)
			}
		})
	})
}

func BenchmarkTimestampGenerator(b *testing.B) {
	dg, err := NewTimestampGenerator(TimestampGeneratorConfig{
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
