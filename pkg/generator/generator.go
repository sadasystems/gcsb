package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"cloud.google.com/go/spanner/spansql"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/sadasystems/gcsb/pkg/schema"
)

var errUnimplemented = errors.New("data generator is not implemented")

// GetGenerator returns a configured generator for the table config
func GetGenerator(config config.TableConfigGenerator) (data.Generator, error) {
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

func GetDataGeneratorMap(cfg config.Config, s schema.Schema) (data.GeneratorMap, error) {
	// col
	// gm := make(data.GeneratorMap, s.Tables().GetNext().Columns().Len())

	// Iterate over the schema

	return nil, nil
}

func GetDataGeneratorMapForTable(cfg config.Config, t schema.Table) (data.GeneratorMap, error) {
	cols := t.Columns()
	gm := make(data.GeneratorMap, cols.Len())

	// Check if table is referenced in config
	ct := cfg.Table(t.Name())

	for cols.HasNext() {
		col := cols.GetNext()
		colType := col.Type()

		var g data.Generator

		var gErr error
		// There is no table/col configs. Use default generators
		if ct == nil {
			g, gErr = GetDefaultGeneratorForType(colType)
		} else {
			// Check if column is in config
			cc := ct.Column(col.Name())

			// The table is in the config, but it has no column config for this column, use default generators
			if cc == nil {
				g, gErr = GetDefaultGeneratorForType(colType)
			} else {
				// The column is referenced in the configuration... Use it to create a generator
				// ..........
			}
		}

		if gErr != nil {
			return nil, fmt.Errorf("building generator map: %s", gErr.Error())
		}

		if g == nil {
			return nil, fmt.Errorf("error getting generator for column '%s', %+v", col.Name(), colType)
		}

		gm[col.Name()] = g
	}

	return gm, nil
}

func GetDefaultGeneratorForType(t spansql.Type) (data.Generator, error) {
	var g data.Generator
	var err error

	switch t.Base {
	case spansql.Bool:
		g, err = data.NewBooleanGenerator(data.BooleanGeneratorConfig{
			Source: rand.NewSource(time.Now().UnixNano()),
		})
	case spansql.String:
		g, err = data.NewStringGenerator(data.StringGeneratorConfig{
			Source: rand.NewSource(time.Now().UnixNano()),
			Length: int(t.Len),
		})
	case spansql.Int64:
		g, err = data.NewInt64Generator(data.Int64GeneratorConfig{
			Source: rand.NewSource(time.Now().UnixNano()),
		})
	case spansql.Float64:
		g, err = data.NewFloat64Generator(data.Float64GeneratorConfig{
			Source: rand.NewSource(time.Now().UnixNano()),
		})
	case spansql.Bytes:
		g, err = data.NewRandomByteGenerator(data.RandomByteGeneratorConfig{
			Length: int(t.Len),
			Source: rand.NewSource(time.Now().UnixNano()),
		})
	case spansql.Timestamp:
		g, err = data.NewTimestampGenerator(data.TimestampGeneratorConfig{
			Source: rand.NewSource(time.Now().UnixNano()),
		})
	case spansql.Date:
		g, err = data.NewDateGenerator(data.DateGeneratorConfig{
			Source: rand.NewSource(time.Now().UnixNano()),
		})
	}

	// The column is an array, re-use our generator
	if t.Array {
		g, err = data.NewArrayGenerator(data.ArrayGeneratorConfig{
			Generator: g,
			Length:    10,
		})
	}

	return g, err
}
