package data

import (
	"math/rand"
	"time"
)

// Assert that Int64Generator implements Generator
var _ Generator = (*Int64Generator)(nil)

type (
	Int64Generator struct {
		src *rand.Rand
		f   func() interface{}
		r   bool
		min int64
		max int64
	}

	Int64GeneratorConfig struct {
		Source  rand.Source
		Range   bool
		Minimum int64
		Maximum int64
	}
)

func NewInt64Generator(cfg Int64GeneratorConfig) (*Int64Generator, error) {
	ret := &Int64Generator{
		r:   cfg.Range,
		min: cfg.Minimum,
		max: cfg.Maximum,
	}

	if cfg.Source == nil {
		ret.src = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else {
		ret.src = rand.New(cfg.Source)
	}

	ret.f = ret.nextRandom
	if ret.r {
		ret.f = ret.nextRanged
	}

	return ret, nil
}

func (g *Int64Generator) Next() interface{} {
	return g.f()
}

func (g *Int64Generator) nextRandom() interface{} {
	return g.src.Int63()
}

func (g *Int64Generator) nextRanged() interface{} {
	return g.src.Int63n(g.max-g.min) + g.min
}
