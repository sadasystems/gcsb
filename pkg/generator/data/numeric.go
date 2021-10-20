package data

import (
	"math/big"
	"math/rand"

	"cloud.google.com/go/spanner/spansql"
)

// Assert that NumericGenerator implements Generator
var _ Generator = (*NumericGenerator)(nil)

type (
	NumericGenerator struct {
		src *rand.Rand
		r   bool
		min int64
		max int64
	}
)

func NewNumericGenerator(cfg Config) (Generator, error) {
	return &NumericGenerator{
		src: rand.New(cfg.Source()),
	}, nil
}

func (g *NumericGenerator) Next() interface{} {
	return big.NewRat(g.src.Int63(), g.src.Int63())
}

func (g *NumericGenerator) Type() spansql.TypeBase {
	return spansql.Numeric
}
