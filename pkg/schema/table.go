package schema

type (
	Table interface {
		Name() string
		Type() string
		HasParent() bool
		Parent() string
	}

	table struct {
		n string
		t string
		p string
	}
)

func NewTable() Table {
	return &table{}
}

func (t *table) Name() string {
	return t.n
}

func (t *table) Type() string {
	return t.t
}

func (t *table) Parent() string {
	return t.p
}

func (t *table) HasParent() bool {
	return t.p != ""
}
