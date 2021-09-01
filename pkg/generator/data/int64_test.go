package data

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt64Generator(t *testing.T) {
	Convey("Int64Generator", t, func() {
		Convey("Nil Source", func() {
			fg, err := NewInt64Generator(Int64GeneratorConfig{})
			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)
		})

		Convey("Random", func() {
			fg, err := NewInt64Generator(Int64GeneratorConfig{
				Source: rand.NewSource(time.Now().UnixNano()),
			})

			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				v, ok := fg.Next().(int64)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			}
		})

		Convey("Ranged", func() {
			fg, err := NewInt64Generator(Int64GeneratorConfig{
				Source:  rand.NewSource(time.Now().UnixNano()),
				Range:   true,
				Minimum: 10,
				Maximum: 100000,
			})

			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				v, ok := fg.Next().(int64)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				So(v, ShouldBeLessThanOrEqualTo, 100000)
				So(v, ShouldBeGreaterThanOrEqualTo, 10)
			}
		})
	})
}

func BenchmarkRandomInt64Generator(b *testing.B) {
	bg, err := NewInt64Generator(Int64GeneratorConfig{
		Source: rand.NewSource(time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}

func BenchmarkRangedInt64Generator(b *testing.B) {
	bg, err := NewInt64Generator(Int64GeneratorConfig{
		Source:  rand.NewSource(time.Now().UnixNano()),
		Range:   true,
		Minimum: 1,
		Maximum: 1000000,
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}
