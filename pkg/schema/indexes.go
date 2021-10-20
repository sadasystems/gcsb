package schema

type (
	IndexIterator interface {
		ResetIterator()
		HasNext() bool
		GetNext() Index
	}

	Indexes interface {
		IndexIterator
		Indexes() []Index
		AddIndex(Index)
	}

	indexes struct {
		iteratorIndex int
		indexes       []Index
	}
)

func NewIndexes() Indexes {
	return &indexes{
		indexes: make([]Index, 0),
	}
}

func (t *indexes) Len() int {
	return len(t.indexes)
}

func (t *indexes) ResetIterator() {
	t.iteratorIndex = 0
}

func (t *indexes) HasNext() bool {
	return t.iteratorIndex < len(t.indexes)
}

func (t *indexes) GetNext() Index {
	if t.HasNext() {
		to := t.indexes[t.iteratorIndex]
		t.iteratorIndex++
		return to
	}

	return nil
}

func (t *indexes) Indexes() []Index {
	return t.indexes
}

func (t *indexes) AddIndex(x Index) {
	t.indexes = append(t.indexes, x)
}
