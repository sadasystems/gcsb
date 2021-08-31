package data

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBooleanGenerator(t *testing.T) {
	Convey("BooleanGenerator", t, func() {

		Convey("Missing  Source", func() {
			bg, err := NewBooleanGenerator(BooleanGeneratorConfig{})

			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)
			So(bg.src, ShouldNotBeNil)
		})

		Convey("Next", func() {
			bg, err := NewBooleanGenerator(BooleanGeneratorConfig{
				Source: rand.NewSource(time.Now().UnixNano()),
			})

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
			bg, err := NewBooleanGenerator(BooleanGeneratorConfig{
				Static: true,
				Value:  false,
			})

			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)

			for i := 0; i < 20; i++ {
				So(bg.Next(), ShouldBeFalse)
			}
		})
	})
}

func BenchmarkBooleanGenerator(b *testing.B) {
	bg, err := NewBooleanGenerator(BooleanGeneratorConfig{
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

func BenchmarkStaticBooleanGenerator(b *testing.B) {
	bg, err := NewBooleanGenerator(BooleanGeneratorConfig{
		Static: true,
		Value:  false,
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}
