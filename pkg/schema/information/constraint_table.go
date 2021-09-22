package information

type (
	// ConstraintTableUsages is a collection of ConstraintTableUsage
	ConstraintTableUsages []*ConstraintTableUsage

	// ConstraintTable is a row rom information_schema.constraint_table_usage (see: https://cloud.google.com/spanner/docs/information-schema#information_schemaconstraint_table_usage)
	ConstraintTableUsage struct {
		// The name of the constrained table's catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the constrained table's schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the constrained table.
		TableName string `spanner:"TABLE_NAME"`
		// The name of the constraint's catalog. Always an empty string.
		ConstraintCatalog string `spanner:"CONSTRAINT_CATALOG"`
		// The name of the constraint's schema. An empty string if unnamed.
		ConstraintSchema string `spanner:"CONSTRAINT_SCHEMA"`
		// The name of the constraint.
		ConstraintName string `spanner:"CONSTRAINT_NAME"`
	}
)
