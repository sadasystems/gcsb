package data

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRandomByteGenerator(b *testing.B) {
	bg, err := NewRandomByteGenerator(RandomByteGeneratorConfig{
		Length: 4,
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
