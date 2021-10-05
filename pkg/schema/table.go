package schema

import "github.com/sadasystems/gcsb/pkg/schema/information"

type (
	Table interface {
		SetName(string)
		Name() string
		SetType(string)
		Type() string
		HasParent() bool
		SetParentName(string)
		ParentName() string
		SetParent(Table)
		Parent() Table
		SetSpanenrState(string)
		SpannerState() string
	}

	table struct {
		n            string
		t            string
		p            string
		parent       Table
		spannerState string
	}
)

func NewTable() Table {
	return &table{}
}

func NewTableFromSchema(x information.Table) Table {
	t := NewTable()

	// TODO: I guess check for nil? This isn't safe
	t.SetName(*x.TableName)
	t.SetType(*x.TableType)
	t.SetSpanenrState(*x.SpannerState)

	if x.ParentTableName != nil {
		t.SetParentName(*x.ParentTableName)
	}

	return t
}

func (t *table) SetName(x string) {
	t.n = x
}

func (t *table) Name() string {
	return t.n
}

func (t *table) SetType(x string) {
	t.t = x
}

func (t *table) Type() string {
	return t.t
}

func (t *table) SetParentName(x string) {
	t.p = x
}

func (t *table) ParentName() string {
	return t.p
}

func (t *table) HasParent() bool {
	return t.p != ""
}

func (t *table) SetParent(x Table) {
	t.parent = x
}

func (t *table) Parent() Table {
	return t.parent
}

func (t *table) SetSpanenrState(x string) {
	t.spannerState = x
}

func (t *table) SpannerState() string {
	return t.spannerState
}
