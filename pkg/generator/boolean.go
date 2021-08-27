package generator

import (
	"math/rand"
	"time"
)

// Assert that RandomBooleanGenerator implements Generator
var _ Generator = (*RandomBooleanGenerator)(nil)

type (
	RandomBooleanGenerator struct {
		src       rand.Source
		cache     int64
		remaining int
	}

	RandomBooleanGeneratorConfig struct {
		Source rand.Source
	}
)

func NewRandomBooleanGenerator(cfg RandomBooleanGeneratorConfig) (*RandomBooleanGenerator, error) {
	ret := &RandomBooleanGenerator{
		src: cfg.Source,
	}

	if ret.src == nil {
		ret.src = rand.NewSource(time.Now().UnixNano())
	}

	return ret, nil
}

func (b *RandomBooleanGenerator) Next() interface{} {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}
