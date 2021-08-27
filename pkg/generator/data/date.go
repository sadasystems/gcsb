package data

import (
	"errors"
	"math/rand"
	"time"

	"cloud.google.com/go/civil"
)

// Assert that RandomDateGenerator implements Generator
var _ Generator = (*RandomDateGenerator)(nil)

type (
	RandomDateGenerator struct {
		src   rand.Source
		delta int64
		min   int64
		max   int64
	}

	RandomDateGeneratorConfig struct {
		Length int
		Source rand.Source
	}
)

func NewRandomDateGenerator(cfg RandomDateGeneratorConfig) (*RandomDateGenerator, error) {
	if cfg.Length <= 0 {
		return nil, errors.New("string length must be > 0")
	}

	ret := &RandomDateGenerator{
		src: cfg.Source,
		min: time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix(),
		max: time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix(),
	}

	if ret.src == nil {
		ret.src = rand.NewSource(time.Now().UnixNano())
	}

	ret.delta = ret.max - ret.min

	return ret, nil
}

func (g *RandomDateGenerator) Next() interface{} {
	sec := rand.Int63n(g.delta) + g.min
	return civil.DateOf(time.Unix(sec, 0))
}
