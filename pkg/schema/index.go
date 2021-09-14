package schema

type (
	// Indexes is a collection of Index
	Indexes []*Index

	// Index is a row in information_schema.indexes (see: https://cloud.google.com/spanner/docs/information-schema#indexes)
	Index struct {
		// The name of the catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the table.
		TableName string `spanner:"TABLE_NAME"`
		// The name of the index. Tables with a PRIMARY KEY specification have a pseudo-index entry generated with the name PRIMARY_KEY, which allows the fields of the primary key to be determined.
		IndexName string `spanner:"INDEX_NAME"`
		// The type of the index. The type is INDEX or PRIMARY_KEY.
		IndexType string `spanner:"INDEX_TYPE"`
		// Secondary indexes can be interleaved in a parent table, as discussed in Creating a secondary index. This column holds the name of that parent table, or NULL if the index is not interleaved.
		ParentTableName string `spanner:"PARENT_TABLE_NAME"`
		// Whether the index keys must be unique.
		IsUnique bool `spanner:"IS_UNIQUE"`
		// Whether the index includes entries with NULL values.
		IsNullFiltered bool `spanner:"IS_NULL_FILTERED"`
		// The current state of the index. Possible values and the states they represent are:
		//   PREPARE: creating empty tables for a new index.
		//   WRITE_ONLY: backfilling data for a new index.
		//   WRITE_ONLY_CLEANUP: cleaning up a new index.
		//   WRITE_ONLY_VALIDATE_UNIQUE: checking uniqueness of data in a new index.
		//   READ_WRITE: normal index operation.
		IndexState string `spanner:"INDEX_STATE"`
		// True if the index is managed by Cloud Spanner; Otherwise, False. Secondary backing indexes for foreign keys are managed by Cloud Spanner.
		SpannerIsManaged bool `spanner:"SPANNER_IS_MANAGED"`
	}
)
