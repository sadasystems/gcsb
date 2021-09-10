package data

import (
	"math/rand"
	"time"
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

	Float64GeneratorConfig struct {
		Source  rand.Source
		Range   bool
		Minimum float64
		Maximum float64
	}
)

func NewFloat64Generator(cfg Float64GeneratorConfig) (*Float64Generator, error) {
	ret := &Float64Generator{
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

func (g *Float64Generator) Next() interface{} {
	return g.f()
}

func (g *Float64Generator) nextRandom() interface{} {
	return g.src.Float64()
}

func (g *Float64Generator) nextRanged() interface{} {
	return g.min + g.src.Float64()*(g.max-g.min)
}
