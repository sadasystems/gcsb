package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt64Generator(t *testing.T) {
	Convey("Int64Generator", t, func() {
		Convey("Nil Source", func() {
			fg, err := NewInt64Generator(NewConfig())
			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)
		})

		Convey("Random", func() {
			fg, err := NewInt64Generator(NewConfig())

			So(err, ShouldBeNil)
			So(fg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				v, ok := fg.Next().(int64)
				So(v, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			}
		})

		Convey("Ranged", func() {
			cfg := NewConfig()
			cfg.SetRange(true)
			cfg.SetMinimum(10)
			cfg.SetMaximum(100000)
			fg, err := NewInt64Generator(cfg)

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
	bg, err := NewInt64Generator(NewConfig())
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}

func BenchmarkRangedInt64Generator(b *testing.B) {
	cfg := NewConfig()
	cfg.SetRange(true)
	cfg.SetMinimum(10)
	cfg.SetMaximum(100000)

	bg, err := NewInt64Generator(cfg)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}
