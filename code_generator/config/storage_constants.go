package config

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

	GO_TYPE_TIMESTAMP string = "*time.Time"
	GO_TYPE_STRING    string = "string"
	GO_TYPE_INT64     string = "int64"

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

var GetGoFieldTypeFor = map[FieldType]string{
	FIELD_TYPE_TIMESTAMP: GO_TYPE_TIMESTAMP,
	FIELD_TYPE_INT64:     GO_TYPE_INT64,
	FIELD_TYPE_STRING:    GO_TYPE_STRING,
}
