package generator

import (
	. "github.com/sadasystems/gcsb/pkg/config"
	. "github.com/sadasystems/gcsb/pkg/generator/data"
)

func GetGenerator(config TableConfigGenerator) Generator {
	switch config.Type {
	case "hexawhatever":
		{
			gen, _ := NewHexavigesimalGenerator(HexavigesimalGeneratorConfig{
				Length:   config.Length,
				KeyRange: &config.KeyRange,
			})
			return gen
		}
	case "combined":
		{
			gen, _ := NewCombinedGenerator(CombinedGeneratorConfig{
				StringLength: config.Length,
				PrefixLength: config.PrefixLength,
				KeyRange:     config.KeyRange,
			})
			return gen
		}
	case "string":
		{
			gen, _ := NewStringGenerator(StringGeneratorConfig{
				Length: config.Length,
			})
			return gen
		}
	case "int64":
		{
			{
				gen, _ := NewInt64Generator(Int64GeneratorConfig{
					Range:   false,
					Minimum: 0,
					Maximum: 1000000,
				})
				return gen
			}
		}
	default:
		return nil
	}
}
