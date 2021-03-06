package information

type (
	// TableConstraints is a collection of TableConstraint
	TableConstraints []*TableConstraint

	// TableConstraint is a row from information_schema.table_constraints (see: https://cloud.google.com/spanner/docs/information-schema#information_schematable_constraints)
	TableConstraint struct {
		// Always an emptry string.
		ConstraintCatalog string `spanner:"CONSTRAINT_CATALOG"`
		// The name of the constraint's schema. An empty string if unnamed.
		ConstraintSchema string `spanner:"CONSTRAINT_SCHEMA"`
		// The name of the constraint.
		ConstraintName string `spanner:"CONSTRAINT_NAME"`
		// The name of constrained table's catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of constrained table's schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the constrained table.
		TableName string `spanner:"TABLE_NAME"`
		// The type of the constraint. Possible values are:
		//   PRIMARY KEY
		//   FOREIGN KEY
		//   CHECK
		//   UNIQUE
		ConstraintType string `spanner:"CONSTRAINT_TYPE"`
		// Always NO.
		IsDeferrable string `spanner:"IS_DEFERRABLE"`
		// Always NO.
		InitiallyDeferred string `spanner:"INITIALLY_DEFERRED"`
		//Always YES.
		Enforced string `spanner:"ENFORCED"`
	}
)
