package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArrayGenerator(t *testing.T) {
	Convey("ArrayGenerator", t, func() {

		Convey("Missing  Generator", func() {
			bg, err := NewBooleanGenerator(NewConfig())
			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)

			cfg := NewConfig()
			cfg.SetLength(10)
			ag, err := NewArrayGenerator(cfg)

			So(err, ShouldNotBeNil)
			So(ag, ShouldBeNil)
		})

		Convey("Invalid Length", func() {
			bg, err := NewBooleanGenerator(NewConfig())
			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)

			cfg := NewConfig()
			cfg.SetGenerator(bg)
			cfg.SetLength(-5)
			ag, err := NewArrayGenerator(cfg)

			So(err, ShouldNotBeNil)
			So(ag, ShouldBeNil)
		})

		Convey("Next", func() {
			bg, err := NewBooleanGenerator(NewConfig())
			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)

			cfg := NewConfig()
			cfg.SetLength(10)
			cfg.SetGenerator(bg)

			ag, err := NewArrayGenerator(cfg)

			So(err, ShouldBeNil)
			So(ag, ShouldNotBeNil)

			v, ok := ag.Next().([]interface{})
			So(ok, ShouldBeTrue)
			So(v, ShouldNotBeNil)

			tagl, ok := ag.(*ArrayGenerator)
			So(ok, ShouldBeTrue)
			So(tagl, ShouldNotBeNil)

			So(v, ShouldHaveLength, tagl.l)

			// So now v is not an interface type, so we can't further type assert it?
			// So I guess we assert each element?
			for _, e := range v {
				_, ok := e.(bool)
				So(ok, ShouldBeTrue)
			}
		})
	})
}

func BenchmarkBooleanArrayGenerator(b *testing.B) {
	bg, err := NewBooleanGenerator(NewConfig())
	if err != nil {
		panic(err)
	}

	cfg := NewConfig()
	cfg.SetLength(10)
	cfg.SetGenerator(bg)
	ag, err := NewArrayGenerator(cfg)

	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ag.Next()
	}
}

func BenchmarkInt64ArrayGenerator(b *testing.B) {
	ig, err := NewInt64Generator(NewConfig())
	if err != nil {
		panic(err)
	}

	cfg := NewConfig()
	cfg.SetLength(10)
	cfg.SetGenerator(ig)
	ag, err := NewArrayGenerator(cfg)

	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ag.Next()
	}
}
