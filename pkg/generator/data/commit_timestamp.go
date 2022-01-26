package data

import (
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
)

// Assert that CommitTimestampGenerator implements Generator
var _ Generator = (*CommitTimestampGenerator)(nil)

type (
	CommitTimestampGenerator struct{}
)

func NewCommitTimestampGenerator(cfg Config) (Generator, error) {
	return &CommitTimestampGenerator{}, nil
}

func (g *CommitTimestampGenerator) Next() interface{} {
	return spanner.CommitTimestamp
}

func (g *CommitTimestampGenerator) Type() spansql.TypeBase {
	return spansql.Timestamp
}
