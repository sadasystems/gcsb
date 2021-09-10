package data

import "errors"

// Assert that ArrayGenerator implements Generator
var _ Generator = (*ArrayGenerator)(nil)

type (
	ArrayGenerator struct {
		g Generator
		l int
	}

	ArrayGeneratorConfig struct {
		Generator Generator
		Length    int
	}
)

func NewArrayGenerator(cfg ArrayGeneratorConfig) (*ArrayGenerator, error) {
	if cfg.Generator == nil {
		return nil, errors.New("array generator requires a generator")
	}

	if cfg.Length <= 0 {
		return nil, errors.New("array generator length must be <= 0")
	}

	ret := &ArrayGenerator{
		g: cfg.Generator,
		l: cfg.Length,
	}

	return ret, nil
}

func (g *ArrayGenerator) Next() interface{} {
	ret := make([]interface{}, 0, g.l)
	for i := 0; i < g.l; i++ {
		ret = append(ret, g.g.Next())
	}

	return ret
}
