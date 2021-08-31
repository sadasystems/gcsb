package data

import (
	"math/rand"
	"time"
)

// Assert that BooleanGenerator implements Generator
var _ Generator = (*BooleanGenerator)(nil)

type (
	BooleanGenerator struct {
		src       rand.Source
		cache     int64
		remaining int
		f         func() interface{}
		v         bool
	}

	BooleanGeneratorConfig struct {
		Source rand.Source
		Static bool
		Value  bool
	}
)

func NewBooleanGenerator(cfg BooleanGeneratorConfig) (*BooleanGenerator, error) {
	ret := &BooleanGenerator{
		src: cfg.Source,
	}

	if ret.src == nil {
		ret.src = rand.NewSource(time.Now().UnixNano())
	}

	// Generate  values by defualt

	// If the config is for a static value, return it instead
	ret.f = ret.next
	if cfg.Static {
		ret.v = cfg.Value
		ret.f = ret.nextStatic
	}

	return ret, nil
}

func (b *BooleanGenerator) Next() interface{} {
	return b.f()
}

// nextStatic returns a static bool value
func (b *BooleanGenerator) nextStatic() interface{} {
	return b.v
}

// next returns a ly generated bool value
func (b *BooleanGenerator) next() interface{} {

	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}
