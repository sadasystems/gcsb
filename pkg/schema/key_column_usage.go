package schema

type (
	// KeyColumnUsages is a collection of KeyColumnUsage
	KeyColumnUsages []*KeyColumnUsage

	// KeyColumnUsage is a row in information_schema.key_column_usage (see: https://cloud.google.com/spanner/docs/information-schema#information_schemakey_column_usage)
	KeyColumnUsage struct {
		// The name of the constraint's catalog. Always an empty string.
		ConstraintCatalog string `spanner:"CONSTRAINT_CATALOG"`
		// The name of the constraint's schema. This column is never null. An empty string if unnamed.
		ConstraintSchema string `spanner:"CONSTRAINT_SCHEMA"`
		// The name of the constraint.
		ConstraintName string `spanner:"CONSTRAINT_NAME"`
		// The name of the constrained column's catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the constrained column's schema. This column is never null. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the constrained column's table.
		TableName string `spanner:"TABLE_NAME"`
		// The name of the column.
		ColumnName string `spanner:"COLUMN_NAME"`
		// The ordinal position of the column within the constraint's key, starting with a value of 1.
		OrdinalPosition int64 `spanner:"ORDINAL_POSITION"`
		// For FOREIGN KEYs, the ordinal position of the column within the unique constraint, starting with a value of 1. This column is null for other constraint types.
		PositionInUniqueConstraint int64 `spanner:"POSITION_IN_UNIQUE_CONSTRAINT"`
	}
)
