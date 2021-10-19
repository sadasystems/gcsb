package data

import (
	"fmt"
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

func NewDateGenerator(cfg Config) (Generator, error) {
	ret := &DateGenerator{
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
			return nil, fmt.Errorf("minimum '%s' of type '%T' invalid for date generator", min, min)
		}

		switch max := cfg.Maximum().(type) {
		case time.Time:
			ret.max = max.Unix()
		case int64:
			ret.max = max
		default:
			return nil, fmt.Errorf("maximum '%s' of type '%T' invalid for date generator", max, max)
		}
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
