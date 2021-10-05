package schema

type (
	TableIterator interface {
		ResetIterator()
		HasNext() bool
		GetNext() Table
	}

	Tables interface {
		TableIterator
		Tables() []Table
		AddTable(Table)
	}

	tables struct {
		iteratorIndex int
		tables        []Table
	}
)

func NewTables() Tables {
	return &tables{
		tables: make([]Table, 0),
	}
}

func (t *tables) Len() int {
	return len(t.tables)
}

func (t *tables) ResetIterator() {
	t.iteratorIndex = 0
}

func (t *tables) HasNext() bool {
	return t.iteratorIndex < len(t.tables)
}

func (t *tables) GetNext() Table {
	if t.HasNext() {
		to := t.tables[t.iteratorIndex]
		t.iteratorIndex++
		return to
	}

	return nil
}

func (t *tables) Tables() []Table {
	return t.tables
}

func (t *tables) AddTable(x Table) {
	t.tables = append(t.tables, x)
}
