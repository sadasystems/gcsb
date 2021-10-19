package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBooleanGenerator(t *testing.T) {
	Convey("BooleanGenerator", t, func() {

		Convey("Missing  Source", func() {
			bg, err := NewBooleanGenerator(NewConfig())

			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)
		})

		Convey("Next", func() {
			bg, err := NewBooleanGenerator(NewConfig())

			So(bg, ShouldNotBeNil)
			So(err, ShouldBeNil)

			var falseSeen, trueSeen bool
			for i := 0; i <= 200; i++ {
				if bg.Next().(bool) {
					trueSeen = true
					continue
				}
				falseSeen = true
			}

			So(falseSeen, ShouldBeTrue)
			So(trueSeen, ShouldBeTrue)
		})

		Convey("Static value", func() {
			cfg := NewConfig()
			cfg.SetStatic(true)
			cfg.SetValue(false)
			bg, err := NewBooleanGenerator(cfg)

			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				So(bg.Next(), ShouldBeFalse)
			}
		})
	})
}

func BenchmarkBooleanGenerator(b *testing.B) {
	bg, err := NewBooleanGenerator(NewConfig())
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}

func BenchmarkStaticBooleanGenerator(b *testing.B) {
	cfg := NewConfig()
	cfg.SetStatic(true)
	cfg.SetValue(false)
	bg, err := NewBooleanGenerator(cfg)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}
