package generator

import (
	"errors"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/data"
)

var errUnimplemented = errors.New("data generator is not implemented")

// GetGenerator returns a configured generator for the table config
func GetGenerator(config config.TableConfigGenerator) (data.Generator, error) {
	var gen data.Generator
	var err error

	switch config.Type {
	case "hexavigesmal":
		gen, err = data.NewHexavigesimalGenerator(data.HexavigesimalGeneratorConfig{
			Length:   config.Length,
			KeyRange: &config.KeyRange,
		})
	case "combined":
		gen, err = data.NewCombinedGenerator(data.CombinedGeneratorConfig{
			StringLength: config.Length,
			PrefixLength: config.PrefixLength,
			KeyRange:     &config.KeyRange,
		})
	case "string":
		gen, err = data.NewStringGenerator(data.StringGeneratorConfig{
			Length: config.Length,
		})

	case "int64":
		gen, err = data.NewInt64Generator(data.Int64GeneratorConfig{
			Range:   false,
			Minimum: int64(config.Min),
			Maximum: int64(config.Max),
		})
	default:
		err = errUnimplemented
	}

	return gen, err
}
