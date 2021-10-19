package data

import (
	"testing"
)

func BenchmarkRandomByteGenerator(b *testing.B) {
	cfg := NewConfig()
	cfg.SetLength(4)
	bg, err := NewRandomByteGenerator(cfg)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Next()
	}
}
