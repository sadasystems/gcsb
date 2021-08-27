package data

import (
	"errors"
	"math/rand"
	"time"
)

// Assert that RandomByteGenerator implements Generator
var _ Generator = (*RandomByteGenerator)(nil)

type (
	RandomByteGenerator struct {
		len int
		src *rand.Rand
	}

	RandomByteGeneratorConfig struct {
		Length int
		Source rand.Source
	}
)

func NewRandomByteGenerator(cfg RandomByteGeneratorConfig) (*RandomByteGenerator, error) {
	if cfg.Length <= 0 {
		return nil, errors.New("string length must be > 0")
	}

	src := cfg.Source
	if src == nil {
		src = rand.NewSource(time.Now().UnixNano())
	}

	ret := &RandomByteGenerator{
		len: cfg.Length,
		src: rand.New(src),
	}

	return ret, nil
}

func (g *RandomByteGenerator) Next() interface{} {
	ret := make([]byte, g.len)
	_, _ = g.src.Read(ret)
	return ret
}
