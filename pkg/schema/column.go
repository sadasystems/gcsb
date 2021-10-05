package schema

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	"github.com/sadasystems/gcsb/pkg/schema/information"
)

var lengthRegexp = regexp.MustCompile(`\(([0-9]+|MAX)\)$`)

type (
	Column interface {
		SetName(string)
		Name() string
		SetPosition(int64)
		Position() int64
		SetNullable(string)
		Nullable() string
		SetSpannerType(string)
		SpannerType() string
		SetIsGenerated(bool)
		IsGenerated() bool
		SetGenerationExpression(string)
		GenerationExpression() string
		SetIsStored(string)
		IsStored() string
		SetSpannerState(string)
		SpannerState() string
		SetPrimaryKey(bool)
		PrimaryKey() bool
		Type() spansql.Type
	}

	column struct {
		name                 string
		position             int64
		nullable             string
		spannerType          string
		isGenerated          bool
		generationExpression string
		isStored             string
		spannerState         string
		primaryKey           bool
	}
)

func NewColumn() Column {
	return &column{}
}

func NewColumnFromSchema(x information.Column) Column {
	c := NewColumn()

	// TODO: nil test these? This is unsafe
	c.SetName(*x.ColumnName)
	c.SetPosition(x.OrdinalPosition)
	c.SetNullable(*x.IsNullable)
	c.SetSpannerType(*x.SpannerType)
	c.SetIsGenerated(x.IsGenerated)
	if x.GenerationExpression != nil {
		c.SetGenerationExpression(*x.GenerationExpression)
	}

	if x.IsStored != nil {
		c.SetIsStored(*x.IsStored)
	}
	c.SetSpannerState(*x.SpannerState)
	c.SetPrimaryKey(x.IsPrimaryKey)

	return c
}

func LoadColumns(ctx context.Context, client *spanner.Client, t Table) error {
	iter := client.Single().Query(ctx, information.GetColumnsQuery(t.Name()))
	defer iter.Stop()

	err := iter.Do(func(row *spanner.Row) error {
		var ci information.Column
		if err := row.ToStruct(&ci); err != nil {
			return err
		}

		cp := NewColumnFromSchema(ci)
		t.AddColumn(cp)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *column) SetName(x string) {
	c.name = x
}

func (c *column) Name() string {
	return c.name
}

func (c *column) SetPosition(x int64) {
	c.position = x
}

func (c *column) Position() int64 {
	return c.position
}

func (c *column) SetNullable(x string) {
	c.nullable = x
}

func (c *column) Nullable() string {
	return c.nullable
}

func (c *column) SetSpannerType(x string) {
	c.spannerType = x
}

func (c *column) SpannerType() string {
	return c.spannerType
}

func (c *column) SetIsGenerated(x bool) {
	c.isGenerated = x
}

func (c *column) IsGenerated() bool {
	return c.isGenerated
}

func (c *column) SetGenerationExpression(x string) {
	c.generationExpression = x
}

func (c *column) GenerationExpression() string {
	return c.generationExpression
}

func (c *column) SetIsStored(x string) {
	c.isStored = x
}

func (c *column) IsStored() string {
	return c.isStored
}

func (c *column) SetSpannerState(x string) {
	c.spannerState = x
}

func (c *column) SpannerState() string {
	return c.spannerState
}

func (c *column) SetPrimaryKey(x bool) {
	c.primaryKey = x
}

func (c *column) PrimaryKey() bool {
	return c.primaryKey
}

// Parse the sql type and wrap it with spansql.Type. Parts of this function borrowed from Yo
func (c *column) Type() spansql.Type {
	ret := spansql.Type{}

	dt := c.spannerType

	// separate type and length from dt with length such as STRING(32) or BYTES(256)
	m := lengthRegexp.FindStringSubmatchIndex(dt)
	if m != nil {
		lengthStr := dt[m[2]:m[3]]
		if lengthStr == "MAX" {
			ret.Len = spansql.MaxLen
		} else {
			l, err := strconv.Atoi(lengthStr)
			if err != nil {
				panic("could not convert precision")
			}
			ret.Len = int64(l)
		}

		// trim length from dt
		dt = dt[:m[0]] + dt[m[1]:]
	}

	if strings.HasPrefix(dt, "ARRAY<") {
		ret.Array = true
		dt = strings.TrimSuffix(strings.TrimPrefix(dt, "ARRAY<"), ">")
	}

	ret.Base = parseType(dt)

	return ret
}

func parseType(dt string) spansql.TypeBase {
	var ret spansql.TypeBase
	switch dt {
	case "BOOL":
		ret = spansql.Bool
	case "STRING":
		ret = spansql.String
	case "INT64":
		ret = spansql.Int64
	case "FLOAT64":
		ret = spansql.Float64
	case "BYTES":
		ret = spansql.Bytes
	case "TIMESTAMP":
		ret = spansql.Timestamp
	case "DATE":
		ret = spansql.Date
	default:
		panic(fmt.Sprintf("unknown spanner type '%s'", dt))
	}

	return ret
}
