package configreader

type (
	FieldType  string
	ReturnType string
	ViewType   string
	SearchType string
)

const (
	FIELD_TYPE_STRING    FieldType = "string"
	FIELD_TYPE_INT64     FieldType = "int64"
	FIELD_TYPE_TIMESTAMP FieldType = "timestamp"

	RETURN_TYPE_PK        ReturnType = "primaryKey"
	RETURN_TYPE_DATA      ReturnType = "data"
	RETURN_TYPE_ROW_COUNT ReturnType = "rowCount"

	VIEW_TYPE_CREATE ViewType = "CREATE"
	VIEW_TYPE_READ   ViewType = "READ"
	VIEW_TYPE_UPDATE ViewType = "UPDATE"
	VIEW_TYPE_DELETE ViewType = "DELETE"

	SEARCH_TYPE_EQUALS       SearchType = "EQUALS"
	SEARCH_TYPE_GREATER_THAN SearchType = "GREATER_THAN"
	SEARCH_TYPE_CONTAINS     SearchType = "CONTAINS"
)
