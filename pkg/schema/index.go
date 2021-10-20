package schema

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/sadasystems/gcsb/pkg/schema/information"
)

type (
	Index interface {
		SetIndexName(string)
		IndexName() string
		SetIsUnique(bool)
		IsUnique() bool
		SetIsNullFiltered(bool)
		IsNullFiltered() bool
		SetIndexState(string)
		IndexState() string
	}

	index struct {
		indexName      string
		isUnique       bool
		isNullFiltered bool
		indexState     string
	}
)

func NewIndex() Index {
	return &index{}
}

func NewIndexFromSchema(x information.Index) Index {
	i := NewIndex()

	i.SetIndexName(x.IndexName)
	i.SetIsUnique(x.IsUnique)
	i.SetIsNullFiltered(x.IsNullFiltered)
	i.SetIndexState(x.IndexState)

	return i
}

func LoadIndexes(ctx context.Context, client *spanner.Client, t Table) error {
	iter := client.Single().Query(ctx, information.GetIndexesQuery(t.Name()))
	defer iter.Stop()
	err := iter.Do(func(row *spanner.Row) error {
		var ti information.Index
		if err := row.ToStruct(&ti); err != nil {
			return err
		}

		tp := NewIndexFromSchema(ti)
		t.AddIndex(tp)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (i *index) SetIndexName(x string) {
	i.indexName = x
}

func (i *index) IndexName() string {
	return i.indexName
}

func (i *index) SetIsUnique(x bool) {
	i.isUnique = x
}

func (i *index) IsUnique() bool {
	return i.isUnique
}

func (i *index) SetIsNullFiltered(x bool) {
	i.isNullFiltered = x
}

func (i *index) IsNullFiltered() bool {
	return i.isNullFiltered
}

func (i *index) SetIndexState(x string) {
	i.indexState = x
}

func (i *index) IndexState() string {
	return i.indexState
}
