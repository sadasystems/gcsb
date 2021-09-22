package sample

import "cloud.google.com/go/spanner/spansql"

var (
	singleTableSingleKey = spansql.CreateTable{
		Name: "Singers",
		Columns: []spansql.ColumnDef{
			{
				Name: "SingerId",
				Type: spansql.Type{
					Base: spansql.Int64,
				},
				NotNull: true,
			},
			{
				Name: "FirstName",
				Type: spansql.Type{
					Base: spansql.String,
					Len:  1024,
				},
			},
			{
				Name: "LastName",
				Type: spansql.Type{
					Base: spansql.String,
					Len:  1024,
				},
			},
			{
				Name: "BirthDate",
				Type: spansql.Type{
					Base: spansql.Date,
				},
			},
			{
				Name: "ByteField",
				Type: spansql.Type{
					Base: spansql.Bytes,
					Len:  1024,
				},
			},
			{
				Name: "FloatField",
				Type: spansql.Type{
					Base: spansql.Float64,
				},
			},
			{
				Name: "ArrayField",
				Type: spansql.Type{
					Array: true,
					Base:  spansql.Int64,
				},
			},
			{
				Name: "TSField",
				Type: spansql.Type{
					Base: spansql.Timestamp,
				},
			},
			{
				Name: "NumericField",
				Type: spansql.Type{
					Base: spansql.Numeric,
				},
			},
		},
		PrimaryKey: []spansql.KeyPart{
			{Column: "SingerId"},
		},
	}
)
