package schema

type (
	ColumnIterator interface {
		ResetIterator()
		HasNext() bool
		GetNext() Column
	}

	Columns interface {
		ColumnIterator
		Columns() []Column
		AddColumn(Column)
		ColumnNames() []string
	}

	columns struct {
		iteratorIndex int
		columns       []Column
	}
)

func NewColumns() Columns {
	return &columns{
		columns: make([]Column, 0),
	}
}

func (c *columns) Len() int {
	return len(c.columns)
}

func (c *columns) ResetIterator() {
	c.iteratorIndex = 0
}

func (c *columns) HasNext() bool {
	return c.iteratorIndex < len(c.columns)
}

func (c *columns) GetNext() Column {
	if c.HasNext() {
		to := c.columns[c.iteratorIndex]
		c.iteratorIndex++
		return to
	}

	return nil
}

func (c *columns) Columns() []Column {
	return c.columns
}

func (c *columns) AddColumn(x Column) {
	c.columns = append(c.columns, x)
}

func (c *columns) ColumnNames() []string {
	ret := make([]string, 0)
	for _, col := range c.columns {
		ret = append(ret, col.Name())
	}

	return ret
}
