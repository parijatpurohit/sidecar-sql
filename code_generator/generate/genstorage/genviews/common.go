package genviews

import (
	"fmt"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genmodels"
)

var (
	viewNamePrefix = map[config.ViewType]string{
		config.VIEW_TYPE_CREATE: CreatePrefix,
		config.VIEW_TYPE_UPDATE: UpdatePrefix,
		config.VIEW_TYPE_DELETE: DeletePrefix,
	}
	ContextField = &FuncFieldConfig{
		FieldName: "ctx",
		FieldType: "context.Context",
	}
	ErrorField = &FuncFieldConfig{
		FieldName: "err",
		FieldType: "error",
	}
)

func getFuncDef(tableName string, view *config.View) *ViewFuncConfig {
	return &ViewFuncConfig{
		Name:       getViewName(view),
		InputVars:  getViewInputParams(tableName, view),
		ReturnVars: getViewResponseParams(tableName, view),
	}
}

func getViewName(view *config.View) string {
	if view.ViewType == config.VIEW_TYPE_READ {
		return view.Name
	}
	return viewNamePrefix[view.ViewType]
}

func getViewInputParams(tableName string, view *config.View) string {
	fields := []*FuncFieldConfig{
		ContextField,
	}
	switch view.ViewType {
	case config.VIEW_TYPE_READ:
		fields = append(fields, &FuncFieldConfig{
			FieldName: QueryFieldName,
			FieldType: getQueryFieldType(tableName, view.Name),
		})
	default:
		fields = append(fields, &FuncFieldConfig{
			FieldName: EntityFieldName,
			FieldType: getEntityType(tableName),
		})
	}
	return getFieldStr(fields)
}

func getViewResponseParams(tableName string, view *config.View) string {
	var fields []*FuncFieldConfig
	if view.ViewType == config.VIEW_TYPE_READ {
		fields = append(fields, getReadReturnParams(tableName))
	}
	fields = append(fields, ErrorField)
	return getFieldStr(fields)
}

func getReadReturnParams(tableName string) *FuncFieldConfig {
	return &FuncFieldConfig{
		FieldName: EntityFieldName,
		FieldType: getEntityType(tableName),
	}
}

func getEntityType(tableName string) string {
	return fmt.Sprintf("[]*%s.%s", fmt.Sprintf(genmodels.SchemaPackageFormat, strings.ToLower(tableName)), tableName)
}

func getQueryFieldType(tableName string, viewName string) string {
	return fmt.Sprintf("*%s.%s", fmt.Sprintf(genmodels.SchemaPackageFormat, strings.ToLower(tableName)), fmt.Sprintf(QueryFieldType, tableName, viewName))
}

func getFieldStr(fields []*FuncFieldConfig) string {
	returnStr := FieldDelimiter
	for _, field := range fields {
		returnStr = fmt.Sprintf("%s%s %s %s", returnStr, FieldDelimiter, field.FieldName, field.FieldType)
	}
	return strings.Trim(strings.Trim(returnStr, FieldDelimiter), " ")
}
