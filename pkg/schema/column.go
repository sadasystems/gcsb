package schema

type (
	// Columns is a collection of Collumns
	Columns []*Column

	// Column is a row from information_schema.columns (see: https://cloud.google.com/spanner/docs/information-schema#information_schemacolumns)
	Column struct {
		// The name of the catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the table.
		TableName string `spanner:"TABLE_NAME"`
		// The name of the column.
		ColumnName string `spanner:"COLUMN_NAME"`
		// The ordinal position of the column in the table, starting with a value of 1.
		OrdinalPosition int64 `spanner:"ORDINAL_POSITION"`
		// Included to satisfy the SQL standard. Always NULL.
		ColumnDefault []byte `spanner:"COLUMN_DEFAULT"`
		// Included to satisfy the SQL standard. Always NULL.
		DataType string `spanner:"DATA_TYPE"`
		// A string that indicates whether the column is nullable. In accordance with the SQL standard, the string is either YES or NO, rather than a Boolean value.
		IsNullable string `spanner:"IS_NULLABLE"`
		// The data type of the column.
		SpannerType string `spanner:"SPANNER_TYPE"`
		// A string that indicates whether the column is generated. The string is either ALWAYS for a generated column or NEVER for a non-generated column.
		IsGenerated string `spanner:"IS_GENERATED"`
		// A string representing the SQL expression of a generated column. NULL if the column is not a generated column.
		GenerationExpression string `spanner:"GENERATION_EXPRESSION"`
		// A string that indicates whether the generated column is stored. The string is always YES for generated columns, and NULL for non-generated columns.
		IsStored string `spanner:"IS_STORED"`
		// The current state of the column. A new stored generated column added to an existing table may go through multiple user-observable states before it is fully usable. Possible values are:
		//   WRITE_ONLY: The column is being backfilled. No read is allowed.
		// 	 COMMITTED: The column is fully usable.
		SpannerState string `spanner:"SPANNER_STATE"`
	}
)
