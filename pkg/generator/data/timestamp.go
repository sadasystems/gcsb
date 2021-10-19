package data

import (
	"fmt"
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
)

func NewTimestampGenerator(cfg Config) (Generator, error) {
	ret := &TimestampGenerator{
		src: cfg.Source(),
	}

	if cfg.Range() {
		ret.r = true

		switch min := cfg.Minimum().(type) {
		case time.Time:
			ret.min = min.Unix()
		case int64:
			ret.min = min
		default:
			return nil, fmt.Errorf("minimum '%s' of type '%T' invalid for timestamp generator", min, min)
		}

		switch max := cfg.Maximum().(type) {
		case time.Time:
			ret.max = max.Unix()
		case int64:
			ret.max = max
		default:
			return nil, fmt.Errorf("maximum '%s' of type '%T' invalid for timestamp generator", max, max)
		}
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
