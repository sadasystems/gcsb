package generator

import (
	"errors"
	"fmt"
	"math/rand"

	"cloud.google.com/go/spanner/spansql"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/schema"
)

var errUnimplemented = errors.New("data generator is not implemented")

// GetGenerator returns a configured generator for the table config
func GetGenerator(config config.TableConfigGenerator) (data.Generator, error) {
	cfg := data.NewConfig()
	var gen data.Generator
	var err error

	switch config.Type {
	case "hexadecimal":
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
		cfg.SetLength(config.Length)
		gen, err = data.NewStringGenerator(cfg)

	case "int64":
		cfg.SetRange(false) // ??? Why set min/max without a range?
		gen, err = data.NewInt64Generator(cfg)
	default:
		err = errUnimplemented
	}

	return gen, err
}

func GetDataGeneratorMap(cfg config.Config, s schema.Schema) (data.GeneratorMap, error) {
	// col
	// gm := make(data.GeneratorMap, s.Tables().GetNext().Columns().Len())

	// Iterate over the schema

	return nil, nil
}

// TODO: Check that schema column and config column are compatible types
// TODO: Check that generator config and column type are compatible types
func GetDataGeneratorMapForTable(cfg config.Config, t schema.Table) (data.GeneratorMap, error) {
	cols := t.Columns()
	gm := make(data.GeneratorMap, cols.Len())

	// Check if table is referenced in config
	ct := cfg.Table(t.Name())

	// Iterate over columns
	for cols.HasNext() {
		col := cols.GetNext()
		colType := col.Type()

		var g data.Generator

		var gErr error
		// There is no table/col configs. Use default generators
		if ct == nil {
			g, gErr = GetDefaultGeneratorForType(colType, nil)
		} else {
			// Check if column is in config
			cc := ct.Column(col.Name())

			// The table is in the config, but it has no column config for this column, use default generators
			if cc == nil {
				g, gErr = GetDefaultGeneratorForType(colType, nil)
			} else {
				// The column is referenced in the configuration... Use it to create a generator
				g, gErr = GetConfiguredGenerator(colType, cc)
			}
		}

		if gErr != nil {
			return nil, fmt.Errorf("building generator map: %s", gErr.Error())
		}

		if g == nil {
			return nil, fmt.Errorf("error getting generator for column '%s', %+v", col.Name(), colType)
		}

		// Assign the generator to the map
		gm[col.Name()] = g
	}

	// Reset the schema columns iterator for future use
	cols.ResetIterator()

	return gm, nil
}

func GetConfiguredGenerator(t spansql.Type, col *config.Column) (data.Generator, error) {
	// The column is referenced in the config file but has no generator config. Use a default
	if col.Generator == nil {
		return GetDefaultGeneratorForType(t, nil)
	}

	var g data.Generator
	var err error

	cfg := data.NewConfig()

	// If the generator config block has a seed, use it as our source
	if col.Generator.Seed != nil {
		cfg.SetSource(rand.NewSource(*col.Generator.Seed))
	}

	// if col.Generator.Length != nil {

	// }

	// // The generator is config exists but has no ranges
	// if len(col.Generator.Range) <= 0 {
	// 	// cfg.SetLength()

	// }

	// If there are multiple ranges, assemble a sub range generator that
	// is the sum of all range configs
	if len(col.Generator.Range) > 1 {
		// g, err = data.NewSubRangeGenerator(cfg)
		// for _, r := range col.Generator.Range {
		// 	sg, sErr := GetDefaultGeneratorForType(t, cfg)
		// }
	} else {
		g, err = GetDefaultGeneratorForType(t, cfg)
	}

	return g, err
}

func GetDefaultGeneratorForType(t spansql.Type, cfg data.Config) (data.Generator, error) {
	if cfg == nil {
		cfg = data.NewConfig()
	}

	var g data.Generator
	var err error

	switch t.Base {
	case spansql.Bool:
		g, err = data.NewBooleanGenerator(cfg)
	case spansql.String:
		// Config for generator has no length specified. Take the columns length
		if cfg.Length() == 0 {
			cfg.SetLength(int(t.Len))
		}
		g, err = data.NewStringGenerator(cfg)
	case spansql.Int64:
		g, err = data.NewInt64Generator(cfg)
	case spansql.Float64:
		g, err = data.NewFloat64Generator(cfg)
	case spansql.Bytes:
		// Config for generator has no length specified. Take the columns length
		if cfg.Length() == 0 {
			cfg.SetLength(int(t.Len))
		}
		g, err = data.NewRandomByteGenerator(cfg)
	case spansql.Timestamp:
		g, err = data.NewTimestampGenerator(cfg)
	case spansql.Date:
		g, err = data.NewDateGenerator(cfg)
	}

	// The column is an array, re-use our generator
	if t.Array {
		if cfg.Length() <= 0 {
			cfg.SetLength(10)
		}
		g, err = data.NewArrayGenerator(cfg)
	}

	return g, err
}
