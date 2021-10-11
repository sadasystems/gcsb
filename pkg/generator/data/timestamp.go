package data

import (
	"math/rand"
	"time"
)

// Assert that TimestampGenerator implements Generator
var _ Generator = (*TimestampGenerator)(nil)

type (
	TimestampGenerator struct {
		src   rand.Source
		delta int64
		min   int64
		max   int64
		r     bool
	}

	TimestampGeneratorConfig struct {
		Source  rand.Source
		Range   bool // If true, only generate dates within min and max range
		Minimum time.Time
		Maximum time.Time
	}
)

func NewTimestampGenerator(cfg TimestampGeneratorConfig) (*TimestampGenerator, error) {
	ret := &TimestampGenerator{}

	if cfg.Source == nil {
		ret.src = rand.NewSource(time.Now().UnixNano())
	} else {
		ret.src = cfg.Source
	}

	if cfg.Range {
		ret.r = true
		ret.min = cfg.Minimum.Unix()
		ret.max = cfg.Maximum.Unix()
	} else {
		ret.min = time.Date(defaultDateMinYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		ret.max = time.Date(defaultDateMaxYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	}

	ret.delta = ret.max - ret.min

	return ret, nil
}

func (g *TimestampGenerator) Next() interface{} {
	sec := rand.Int63n(g.delta) + g.min
	return time.Unix(sec, 0)
}
