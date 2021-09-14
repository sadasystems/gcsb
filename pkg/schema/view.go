package schema

type (
	// Views is a collection of View
	Views []*View

	// View is a row in information_schema.views (see: https://cloud.google.com/spanner/docs/information-schema#information_schemaviews)
	View struct {
		// The name of the catalog. Always an empty string.
		TableCatalog string `spanner:"TABLE_CATALOG"`
		// The name of the schema. An empty string if unnamed.
		TableSchema string `spanner:"TABLE_SCHEMA"`
		// The name of the view.
		TableName string `spanner:"TABLE_NAME"`
		// The SQL text of the query that defines the view.
		ViewDefinition string `spanner:"VIEW_DEFINITION"`
	}
)
