package data

import (
	"errors"
	"math/rand"

	"cloud.google.com/go/spanner/spansql"
)

// Assert that RandomByteGenerator implements Generator
var _ Generator = (*RandomByteGenerator)(nil)

type (
	RandomByteGenerator struct {
		len int
		src *rand.Rand
	}
)

func NewRandomByteGenerator(cfg Config) (Generator, error) {
	if cfg.Length() <= 0 {
		return nil, errors.New("string length must be > 0")
	}

	ret := &RandomByteGenerator{
		len: cfg.Length(),
		src: rand.New(cfg.Source()),
	}

	return ret, nil
}

func (g *RandomByteGenerator) Next() interface{} {
	ret := make([]byte, g.len)
	_, _ = g.src.Read(ret)
	return ret
}

func (g *RandomByteGenerator) Type() spansql.TypeBase {
	return spansql.Bytes
}
