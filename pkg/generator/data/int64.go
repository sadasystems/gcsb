package data

import (
	"fmt"
	"math/rand"
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
)

func NewInt64Generator(cfg Config) (Generator, error) {
	ret := &Int64Generator{
		src: rand.New(cfg.Source()),
	}

	ret.f = ret.nextRandom
	if cfg.Range() {
		ret.r = true

		switch min := cfg.Minimum().(type) {
		case int:
			ret.min = int64(min)
		case int64:
			ret.min = min
		default:
			return nil, fmt.Errorf("minimum '%s' of type '%T' invalid for int64 generator", min, min)
		}

		switch max := cfg.Maximum().(type) {
		case int:
			ret.max = int64(max)
		case int64:
			ret.max = max
		default:
			return nil, fmt.Errorf("maximum '%s' of type '%T' invalid for int64 generator", max, max)
		}

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
