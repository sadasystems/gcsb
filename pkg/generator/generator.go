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

// TODO: Handle static value generator (table samples)
// TODO: Handle random string generator vs ranged string generation

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

	if col.Generator.Length != nil {
		cfg.SetLength(*col.Generator.Length)
	}

	// If there are multiple ranges, assemble a sub range generator that
	// contains a generator per range config
	if len(col.Generator.Range) > 1 {
		// Initialize a sub generator
		g, err = data.NewSubRangeGenerator(cfg)

		// Type assert Generator as SubGenerator so we can call methods not defined in the Generator interface
		tg, ok := g.(*data.SubRangeGenerator)

		if !ok {
			// If this happens, something weird is going on. The constructor above returned the wrong type
			return nil, errors.New("subrangegenerator failed to implement generator interface (This is a bug)")
		}

		// Iterate over each range in the generator config
		for _, r := range col.Generator.Range {
			// Copy the config
			cpCfg := cfg.Copy()

			// Set the copies settings based off the current range
			SetDataConfigFromRange(cpCfg, r)

			// Initialize a generator for this range
			sg, sErr := GetDefaultGeneratorForType(t, cpCfg)
			if sErr != nil {
				return nil, fmt.Errorf("failed to initialize subrange generator: %s", sErr.Error())
			}

			// Add generator to SubGenerator
			tg.AddGenerator(sg) // Add generator to subrange
		}
	} else {
		// If there is no range use the default generator for the column
		if len(col.Generator.Range) <= 0 {
			return GetDefaultGeneratorForType(t, cfg)
		}

		// If we are here, it means there is only 1 range config for the generator
		// This means that we only expect one generator for this column

		// Copy the config
		cpCfg := cfg.Copy()

		// Set it's values from the 1 and only range configured
		SetDataConfigFromRange(cpCfg, col.Generator.Range[0])

		// Initialize the generator based on that config
		g, err = GetDefaultGeneratorForType(t, cpCfg)
	}

	return g, err
}

// GetDefaultGeneratorForType will assemble a generator for a spanner column type. If a config is passed,
// it will us that config when initializing the generator
func GetDefaultGeneratorForType(t spansql.Type, cfg data.Config) (data.Generator, error) {
	if cfg == nil {
		cfg = data.NewConfig()
	}

	cfg.SetSpannerType(t)

	var g data.Generator
	var err error

	switch t.Base {
	case spansql.Bool:
		g, err = data.NewBooleanGenerator(cfg)
	case spansql.String:
		// TODO: Check if config indicates we should do random string generator or hexvigesimal

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
	case spansql.Numeric:
		g, err = data.NewNumericGenerator(cfg)
	}

	// The column is an array, re-use our generator
	if t.Array {
		cfg.SetGenerator(g)
		if cfg.Length() <= 0 {
			// TODO: Make default array length configurable
			cfg.SetLength(10) // If no length is specified, default to 10
		}
		g, err = data.NewArrayGenerator(cfg)
	}

	return g, err
}

// SetDataConfigFromRange will set values from a range in the data.Config if they're defined
func SetDataConfigFromRange(cpCfg data.Config, r *config.Range) {
	if r.Begin != nil {
		cpCfg.SetBegin(*r.Begin)
	}

	if r.End != nil {
		cpCfg.SetEnd(*r.End)
	}

	if r.Length != nil {
		cpCfg.SetLength(*r.Length)
	}

	if r.Maximum != nil {
		cpCfg.SetMaximum(*r.Maximum)
	}

	if r.Minimum != nil {
		cpCfg.SetMinimum(*r.Minimum)
	}

	if r.Static != nil {
		cpCfg.SetStatic(*r.Static)
	}

	if r.Value != nil {
		cpCfg.SetValue(*r.Value)
	}

	// If minimum and maximum are not zero, constrain the generator to the range
	if cpCfg.Minimum() != 0 || cpCfg.Maximum() != 0 {
		cpCfg.SetRange(true)
	}

}
