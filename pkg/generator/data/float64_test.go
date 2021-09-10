package data

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFloat64Generator(t *testing.T) {
	Convey("Float64Generator", t, func() {
		Convey("Nil Source", func() {
			fg, err := NewFloat64Generator(Float64GeneratorConfig{})
			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)
		})

		Convey("Random", func() {
			fg, err := NewFloat64Generator(Float64GeneratorConfig{
				Source: rand.NewSource(time.Now().UnixNano()),
			})

			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				v, ok := fg.Next().(float64)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			}
		})

		Convey("Ranged", func() {
			fg, err := NewFloat64Generator(Float64GeneratorConfig{
				Source:  rand.NewSource(time.Now().UnixNano()),
				Range:   true,
				Minimum: 5.0,
				Maximum: -5.0,
			})

			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				v, ok := fg.Next().(float64)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
				So(v, ShouldBeLessThanOrEqualTo, 5.0)
				So(v, ShouldBeGreaterThanOrEqualTo, -5.0)
			}
		})
	})
}

func BenchmarkRandomFloatGenerator(b *testing.B) {
	bg, err := NewFloat64Generator(Float64GeneratorConfig{
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

func BenchmarkRangedFloatGenerator(b *testing.B) {
	bg, err := NewFloat64Generator(Float64GeneratorConfig{
		Source:  rand.NewSource(time.Now().UnixNano()),
		Range:   true,
		Minimum: 0,
		Maximum: 1000,
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}
