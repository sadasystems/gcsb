package schema

import (
	"context"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spansql"
	"github.com/sadasystems/gcsb/pkg/schema/information"
)

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
		SetAllowCommitTimestamp(bool)
		AllowCommitTimestamp() bool
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
		allowCommitTimestamp bool
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

	c.SetAllowCommitTimestamp(x.AllowCommitTimestamp)

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

func (c *column) SetAllowCommitTimestamp(x bool) {
	c.allowCommitTimestamp = x
}

func (c *column) AllowCommitTimestamp() bool {
	return c.allowCommitTimestamp
}

// Parse the sql type and wrap it with spansql.Type. Parts of this function borrowed from Yo
func (c *column) Type() spansql.Type {
	return ParseSpannerType(c.spannerType)
}
