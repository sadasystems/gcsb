package schema

type (
	// ConstraintcolumnUsages is a collection of ConstraintColumnUsage
	ConstraintColumnUsages []*ConstraintColumnUsage

	// ConstraintColumnUsage is a row from information_schema.constraint_column_usage (see: https://cloud.google.com/spanner/docs/information-schema#information_schemaconstraint_column_usage)
	ConstraintColumnUsage struct {
		// The name of the column table's catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the column table's schema. This column is never null. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the column's table.
		TableName string `spanner:"TABLE_NAME"`
		// The name of the column that is used by the constraint.
		ColumnName string `spanner:"COLUMN_NAME"`
		// The name of the constraint's catalog. Always an empty string.
		ConstraintCatalog string `spanner:"CONSTRAINT_CATALOG"`
		// The name of the constraint's schema. An empty string if unnamed.
		ConstraintSchema string `spanner:"CONSTRAINT_SCHEMA"`
		// The name of the constraint.
		ConstraintName string `spanner:"CONSTRAINT_NAME"`
	}
)
