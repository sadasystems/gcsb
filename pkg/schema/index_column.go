package schema

type (
	// IndexColumns is a collection of IndexColumn
	IndexColumns []*IndexColumn

	// IndexColumn is a row in information_schema.index_columns (see: https://cloud.google.com/spanner/docs/information-schema#information_schemaindex_columns)
	IndexColumn struct {
		// The name of the catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the table.
		TableName string `spanner:"TABLE_NAME"`
		// The name of the index.
		IndexName string `spanner:"INDEX_NAME"`
		// The name of the column.
		ColumnName string `spanner:"COLUMN_NAME"`
		// The ordinal position of the column in the index (or primary key), starting with a value of 1. This value is NULL for non-key columns (for example, columns specified in the STORING clause of an index).
		OrdinalPosition int64 `spanner:"ORDINAL_POSITION"`
		// The ordering of the column. The value is ASC or DESC for key columns, and NULL for non-key columns (for example, columns specified in the STORING clause of an index).
		ColumnOrdering string `spanner:"COLUMN_ORDERING"`
		// A string that indicates whether the column is nullable. In accordance with the SQL standard, the string is either YES or NO, rather than a Boolean value.
		IsNullable string `spanner:"IS_NULLABLE"`
		// The data type of the column.
		SpannerType string `spanner:"SPANNER_TYPE"`
	}
)
