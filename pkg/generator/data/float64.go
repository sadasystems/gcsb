package data

import (
	"fmt"
	"math/rand"
)

// Assert that Float64Generator implements Generator
var _ Generator = (*Float64Generator)(nil)

type (
	Float64Generator struct {
		src *rand.Rand
		f   func() interface{}
		r   bool
		min float64
		max float64
	}
)

func NewFloat64Generator(cfg Config) (Generator, error) {
	ret := &Float64Generator{
		src: rand.New(cfg.Source()),
	}

	ret.f = ret.nextRandom
	if cfg.Range() {
		ret.f = ret.nextRanged
		ret.r = true

		switch min := cfg.Minimum().(type) {
		case float64:
			ret.min = min
		default:
			return nil, fmt.Errorf("minimum '%s' of type '%T' invalid for float64 generator", min, min)
		}

		switch max := cfg.Maximum().(type) {
		case float64:
			ret.max = max
		default:
			return nil, fmt.Errorf("maximum '%s' of type '%T' invalid for float64 generator", max, max)
		}
	}

	return ret, nil
}

func (g *Float64Generator) Next() interface{} {
	return g.f()
}

func (g *Float64Generator) nextRandom() interface{} {
	return g.src.Float64()
}

func (g *Float64Generator) nextRanged() interface{} {
	return g.min + g.src.Float64()*(g.max-g.min)
}
