package data

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomBooleanGenerator(t *testing.T) {
	Convey("RandomBooleanGenerator", t, func() {

		Convey("Missing Random Source", func() {
			bg, err := NewRandomBooleanGenerator(RandomBooleanGeneratorConfig{})

			So(err, ShouldBeNil)
			So(bg, ShouldNotBeNil)
			So(bg.src, ShouldNotBeNil)
		})

		Convey("Next", func() {
			bg, err := NewRandomBooleanGenerator(RandomBooleanGeneratorConfig{
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
	})
}

func BenchmarkRandomBooleanGenerator(b *testing.B) {
	bg, err := NewRandomBooleanGenerator(RandomBooleanGeneratorConfig{
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
