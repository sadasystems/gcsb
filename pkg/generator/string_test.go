package generator

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringGenerator(t *testing.T) {
	Convey("StringGenerator", t, func() {
		Convey("Invalid length", func() {
			sg, err := NewStringGenerator(StringGeneratorConfig{
				Length: -10,
			})

			So(sg, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "string length must be > 0")
		})

		Convey("Missing Random Source", func() {
			sg, err := NewStringGenerator(StringGeneratorConfig{
				Length: 5,
			})

			So(err, ShouldBeNil)
			So(sg, ShouldNotBeNil)
			So(sg.src, ShouldNotBeNil)
		})

		Convey("Next", func() {
			sg, err := NewStringGenerator(StringGeneratorConfig{
				Length: 16,
				Source: rand.NewSource(time.Now().UnixNano()),
			})

			So(sg, ShouldNotBeNil)
			So(err, ShouldBeNil)

			for i := 0; i <= 20; i++ {
				So(sg.Next(), ShouldHaveLength, sg.len)
			}
		})
	})
}

func BenchmarkStringGenerator(b *testing.B) {
	sg, err := NewStringGenerator(StringGeneratorConfig{
		Length: 5,
		Source: rand.NewSource(time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sg.Next()
	}
}
