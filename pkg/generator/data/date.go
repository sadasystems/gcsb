package data

import (
	"math/rand"
	"time"

	"cloud.google.com/go/civil"
)

// Assert that DateGenerator implements Generator
var _ Generator = (*DateGenerator)(nil)

const (
	defaultDateMinYear = 1970
	defaultDateMaxYear = 2070
)

type (
	DateGenerator struct {
		src   rand.Source
		delta int64
		min   int64
		max   int64
		r     bool
	}

	DateGeneratorConfig struct {
		Source  rand.Source
		Range   bool // If true, only generate dates within min and max range
		Minimum time.Time
		Maximum time.Time
	}
)

func NewDateGenerator(cfg DateGeneratorConfig) (*DateGenerator, error) {
	ret := &DateGenerator{}

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

func (g *DateGenerator) Next() interface{} {
	sec := rand.Int63n(g.delta) + g.min
	return civil.DateOf(time.Unix(sec, 0))
}
