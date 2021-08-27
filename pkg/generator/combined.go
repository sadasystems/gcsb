package generator

import (
	"math/rand"
	"time"
)

// Generator is the interface that all data generators must satisfy
type (
	CombinedGenerator struct {
		prefixLength int
		stringLength int
		min          int
		max          int
		sg           StringGenerator
		hg           HexavigesimalGenerator
	}

	CombinedGeneratorConfig struct {
		PrefixLength int
		StringLength int
		Min          int
		Max          int
	}
)

var _ Generator = (*CombinedGenerator)(nil)

func NewCombinedGenerator(cfg CombinedGeneratorConfig) CombinedGenerator {
	gen := CombinedGenerator{}
	gen.prefixLength = cfg.PrefixLength
	gen.stringLength = cfg.StringLength
	gen.min = cfg.Min
	gen.max = cfg.Max
	sg, _ := NewStringGenerator(StringGeneratorConfig{
		Length: gen.stringLength - gen.prefixLength,
		Source: rand.NewSource(time.Now().UnixNano() * int64(gen.min)),
	})
	hg, _ := NewHexavigesimalGenerator(HexavigesimalGeneratorConfig{
		Minimum: gen.min,
		Maximum: gen.max,
	})
	gen.sg = sg
	gen.hg = hg
	return gen
}

func (s *CombinedGenerator) Next() string {

	prefix := s.hg.Next()
	rest := s.sg.Next()
	ret := prefix + rest
	return ret
}
