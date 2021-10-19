package data

import (
	"math/rand"
	"time"
)

type (
	Config interface {
		SetBegin(interface{})
		Begin() interface{}
		SetEnd(interface{})
		End() interface{}
		SetLength(int)
		Length() int
		SetStatic(bool)
		Static() bool
		SetValue(interface{})
		Value() interface{}
		SetMinimum(interface{})
		Minimum() interface{}
		SetMaximum(interface{})
		Maximum() interface{}
		SetSource(rand.Source)
		Source() rand.Source
		SetRange(bool)
		Range() bool
	}

	generatorConfig struct {
		source  rand.Source
		begin   interface{}
		end     interface{}
		length  int
		static  bool
		value   interface{}
		minimum interface{}
		maximum interface{}
		ranged  bool
	}
)

func NewConfig() Config {
	return &generatorConfig{}
}

func (c *generatorConfig) SetSource(x rand.Source) {
	c.source = x
}

func (c *generatorConfig) Source() rand.Source {
	if c.source == nil {
		c.source = rand.NewSource(time.Now().UnixNano())
	}

	return c.source
}

func (c *generatorConfig) SetBegin(x interface{}) {
	c.begin = x
}

func (c *generatorConfig) Begin() interface{} {
	return c.begin
}

func (c *generatorConfig) SetEnd(x interface{}) {
	c.end = x
}

func (c *generatorConfig) End() interface{} {
	return c.end
}

func (c *generatorConfig) SetLength(x int) {
	c.length = x
}

func (c *generatorConfig) Length() int {
	return c.length
}

func (c *generatorConfig) SetStatic(x bool) {
	c.static = x
}

func (c *generatorConfig) Static() bool {
	return c.static
}

func (c *generatorConfig) SetValue(x interface{}) {
	c.value = x
}

func (c *generatorConfig) Value() interface{} {
	return c.value
}

func (c *generatorConfig) SetMinimum(x interface{}) {
	c.minimum = x
}

func (c *generatorConfig) Minimum() interface{} {
	return c.minimum
}

func (c *generatorConfig) SetMaximum(x interface{}) {
	c.maximum = x
}

func (c *generatorConfig) Maximum() interface{} {
	return c.maximum
}

func (c *generatorConfig) SetRange(x bool) {
	c.ranged = x
}

func (c *generatorConfig) Range() bool {
	return c.ranged
}
