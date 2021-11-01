package data

import (
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	"github.com/sadasystems/gcsb/pkg/generator/sample"
)

var (
	_ Generator = (*CompositeKey)(nil)
	_ Generator = (*SingleKey)(nil)
)

type (
	// Composite key returns a spanner key containing multiple static values
	CompositeKey struct {
		g *sample.SampleGenerator
	}

	// SingleKey is used to return a spanner key containing only 1 static value
	SingleKey struct {
		g Generator
	}
)

func (g *CompositeKey) Next() interface{} {
	return g.g.Next()
}

func (g *CompositeKey) Type() spansql.TypeBase {
	return 0
}

func (g *SingleKey) Next() interface{} {
	return spanner.Key{g.g.Next()}
}

func (g *SingleKey) Type() spansql.TypeBase {
	return 0
}
