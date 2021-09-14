package schema

type (
	// Tables is a collection of tables
	Tables []*Table

	// Table is a row in information_schema.tables (see: https://cloud.google.com/spanner/docs/information-schema#information_schematables)
	Table struct {
		// The name of the catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the table or view.
		TableName string `spanner:"TABLE_NAME"`
		// The type of the table. For tables it has the value BASE TABLE; for views it has the value VIEW.
		TableType string `spanner:"TABLE_TYPE"`
		// The name of the parent table if this table is interleaved, or NULL.
		ParentTableName string `spanner:"PARENT_TABLE_NAME"`
		// This is set to CASCADE or NO ACTION for interleaved tables, and NULL otherwise. See TABLE statements for more information.
		OnDeleteAction string `spanner:"ON_DELETE_ACTION"`
		// A table can go through multiple states during creation, if bulk operations are involved. For example, when the table is created with a foreign key that requires backfilling of its indexes. Possible states are:
		//   ADDING_FOREIGN_KEY: Adding the table's foreign keys.
		//   WAITING_FOR_COMMIT: Finalizing the schema change.
		//   COMMITTED: The schema change to create the table has been committed. You cannot write to the table until the change is committed.
		SpannerState string `spanner:"SPANNER_STATE"`
	}
)
