package alias

import (
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

var GetFileNameForView = map[config.ViewType]string{
	config.VIEW_TYPE_CREATE: strings.ToTitle(string(config.VIEW_TYPE_CREATE)),
	config.VIEW_TYPE_UPDATE: strings.ToTitle(string(config.VIEW_TYPE_UPDATE)),
	config.VIEW_TYPE_DELETE: strings.ToTitle(string(config.VIEW_TYPE_DELETE)),
}

var GetTemplateForView = map[config.ViewType]string{
	config.VIEW_TYPE_CREATE: paths.CreateViewTemplateFile,
	config.VIEW_TYPE_UPDATE: paths.UpdateViewTemplateFile,
	config.VIEW_TYPE_DELETE: paths.DeleteViewTemplateFile,
	config.VIEW_TYPE_READ:   paths.ReadViewTemplateFile,
}

var GetGoFieldTypeFor = map[config.FieldType]string{
	config.FIELD_TYPE_TIMESTAMP: config.GO_TYPE_TIMESTAMP,
	config.FIELD_TYPE_INT64:     config.GO_TYPE_INT64,
	config.FIELD_TYPE_STRING:    config.GO_TYPE_STRING,
}

var GetProtoFieldTypeFor = map[config.FieldType]string{
	config.FIELD_TYPE_TIMESTAMP: config.PROTO_TYPE_TIMESTAMP,
	config.FIELD_TYPE_INT64:     config.PROTO_TYPE_INT64,
	config.FIELD_TYPE_STRING:    config.PROTO_TYPE_STRING,
}

var OperatorMap = map[config.SearchType]string{
	config.SEARCH_TYPE_EQUALS:       "=",
	config.SEARCH_TYPE_GREATER_THAN: ">",
	config.SEARCH_TYPE_SMALLER_THAN: "<",
	config.SEARCH_TYPE_CONTAINS:     "IN",
}
