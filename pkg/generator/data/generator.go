package data

import "cloud.google.com/go/spanner/spansql"

// Generator is the interface that all data generators must satisfy
type (
	Generator interface {
		Type() spansql.TypeBase
		Next() interface{}
	}
)
