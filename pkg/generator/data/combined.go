package data

import (
	"fmt"

	"cloud.google.com/go/spanner/spansql"
	"github.com/sadasystems/gcsb/pkg/config"
)

// Generator is the interface that all data generators must satisfy
type (
	CombinedGenerator struct {
		prefixLength int
		stringLength int
		min          int
		max          int
		sg           Generator
		hg           Generator
	}

	CombinedGeneratorConfig struct {
		PrefixLength int
		StringLength int
		Min          int
		Max          int
		KeyRange     *config.TableConfigGeneratorRange
	}
)

var _ Generator = (*CombinedGenerator)(nil)

func NewCombinedGenerator(cfg CombinedGeneratorConfig) (*CombinedGenerator, error) {
	gen := &CombinedGenerator{
		min:          cfg.Min,
		max:          cfg.Max,
		prefixLength: cfg.PrefixLength,
		stringLength: cfg.StringLength,
	}

	// TODO: Should stringgenerator receive the same rand.Source as the combined generator?
	scfg := NewConfig()
	scfg.SetLength(gen.stringLength - gen.prefixLength)
	sg, err := NewStringGenerator(scfg)

	if err != nil {
		return nil, err
	}

	// TODO: Should HexavigesimaGenerator receive the same rand.Source as combined generator?
	hg, err := NewHexavigesimalGenerator(HexavigesimalGeneratorConfig{
		Length:   cfg.PrefixLength,
		Minimum:  gen.min,
		Maximum:  gen.max,
		KeyRange: cfg.KeyRange,
	})

	if err != nil {
		return nil, err
	}

	gen.sg = sg
	gen.hg = hg

	return gen, nil
}

func (s *CombinedGenerator) Next() interface{} {
	prefix := s.hg.Next()
	rest := s.sg.Next()
	ret := fmt.Sprintf("%s%s", prefix, rest)

	return ret
}

func (s *CombinedGenerator) Type() spansql.TypeBase {
	return spansql.Bytes
}
