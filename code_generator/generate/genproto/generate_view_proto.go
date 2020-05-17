package genproto

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	QueryFieldPlaceholder      = "Query"
	CountFieldName             = "count"
	RequestStartFieldIndex     = 1
	ResponseStartFieldIndex    = 1
	QueryStartFieldIndex       = 1
	ResponsePKStartIndex       = 1
	PrimaryKeyFieldPlaceholder = "PrimaryKey"
)

type QueryConfig struct {
	Name   string
	Fields []*ProtoFieldConfig
}

type RequestConfig struct {
	Field *ProtoFieldConfig
}

type DataReturnConfig struct {
	IsRepeated bool
	FieldName  string
	Index      int
}
type ResponseConfig struct {
	Field *ProtoFieldConfig
}
type ResponsePKConfig struct {
	Name   string
	Fields []*ProtoFieldConfig
}
type ViewProtoConfig struct {
	Imports               []string
	TableName             string
	ViewName              string
	IsQueryPopulated      bool
	IsResponsePKPopulated bool
	RequestConfig         *RequestConfig
	QueryConfig           *QueryConfig
	ResponseConfig        *ResponseConfig
	ResponsePKConfig      *ResponsePKConfig
}

// GenerateViewProto generates View protobuf config including request and response objects
func GenerateViewProto(storageConfig *config.StorageConfig) {
	fieldSchema, primaryKeys := generateUtils.GetFieldConfig(storageConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	log.Printf("generating views for table: %s", tableName)
	for _, view := range storageConfig.Views {
		log.Printf("generating view: %s_%s", tableName, view.Name)

		viewConfig := ViewProtoConfig{
			Imports:               getImports(tableName, primaryKeys, view, fieldSchema),
			TableName:             tableName,
			ViewName:              view.Name,
			IsQueryPopulated:      view.ViewType == config.VIEW_TYPE_READ,
			IsResponsePKPopulated: view.Config.ReturnType == config.RETURN_TYPE_PK,
			RequestConfig:         getRequestConfig(tableName, view),
			QueryConfig:           getQueryConfig(tableName, view, fieldSchema),
			ResponseConfig:        getResponseConfig(tableName, view),
			ResponsePKConfig:      getPrimaryKeyConfig(tableName, view.Name, primaryKeys),
		}
		template := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ProtoTemplatePath, paths.ProtoViewTemplateFile))
		outputFile, err := getOutputViewProtoFile(tableName, view.Name)
		if err != nil {
			panic(err)
		}
		err = template.Execute(outputFile, viewConfig)
		if err != nil {
			panic(err)
		}
	}
}

// generates all imports for the file and returns for template to use
func getImports(tableName string, primaryKeys []*config.Field, viewConfig *config.View, fieldSchema map[string]*config.Field) []string {
	schemaImport := fmt.Sprintf("%s_schema.proto", tableName)
	imports := []string{schemaImport}
	if viewConfig.Query != nil {
		for _, field := range viewConfig.Query.Fields {
			if fieldSchema[field.FieldName].FieldType == config.FIELD_TYPE_TIMESTAMP {
				imports = append(imports, ImportTimestampProto)
			}
		}
	}
	if viewConfig.Config.ReturnType == config.RETURN_TYPE_PK {
		for _, pk := range primaryKeys {
			if pk.FieldType == config.FIELD_TYPE_TIMESTAMP {
				imports = append(imports, ImportTimestampProto)
			}
		}
	}
	return utils.GetUnique(imports)

}

// generates all request fields excluding query and returns for template to use
func getRequestConfig(tableName string, viewConfig *config.View) *RequestConfig {
	requestConfig := &RequestConfig{}
	switch viewConfig.ViewType {
	case config.VIEW_TYPE_READ:
		requestConfig.Field = &ProtoFieldConfig{
			FieldType:  fmt.Sprintf("%s_%s_%s", tableName, viewConfig.Name, QueryFieldPlaceholder),
			FieldName:  strings.ToLower(QueryFieldPlaceholder),
			FieldIndex: RequestStartFieldIndex,
		}
	case config.VIEW_TYPE_CREATE, config.VIEW_TYPE_UPDATE, config.VIEW_TYPE_DELETE:
		requestConfig.Field = &ProtoFieldConfig{
			IsRepeated: viewConfig.Config.MultiInput,
			FieldType:  tableName,
			FieldName:  getPlural(viewConfig.Config.MultiInput, tableName),
			FieldIndex: RequestStartFieldIndex,
		}
	}
	return requestConfig
}

// generates query struct. doesn't check if query needs to be generated
func getQueryConfig(tableName string, viewConfig *config.View, fieldSchema map[string]*config.Field) *QueryConfig {
	queryName := fmt.Sprintf("%s_%s_%s", tableName, viewConfig.Name, QueryFieldPlaceholder)
	var queryFields []*ProtoFieldConfig
	if viewConfig.Query != nil {
		for index, field := range viewConfig.Query.Fields {
			queryFields = append(queryFields, &ProtoFieldConfig{
				IsRepeated: field.SearchType == config.SEARCH_TYPE_CONTAINS,
				FieldName:  field.FieldName,
				FieldType:  config.GetProtoFieldTypeFor[fieldSchema[field.FieldName].FieldType],
				FieldIndex: QueryStartFieldIndex + index,
			})
		}
	}
	return &QueryConfig{Name: queryName, Fields: queryFields}
}

// generates response object config without primary key config
func getResponseConfig(tableName string, viewConfig *config.View) *ResponseConfig {
	var responseField *ProtoFieldConfig
	switch viewConfig.Config.ReturnType {
	case config.RETURN_TYPE_DATA:
		responseField = &ProtoFieldConfig{
			IsRepeated: viewConfig.Config.MultiReturn,
			FieldName:  getPlural(viewConfig.Config.MultiInput, tableName),
			FieldType:  tableName,
			FieldIndex: ResponseStartFieldIndex,
		}
	case config.RETURN_TYPE_ROW_COUNT:
		responseField = &ProtoFieldConfig{
			FieldName:  CountFieldName,
			FieldType:  config.PROTO_TYPE_INT64,
			FieldIndex: ResponseStartFieldIndex,
		}
	case config.RETURN_TYPE_PK:
		responseField = &ProtoFieldConfig{
			IsRepeated: viewConfig.Config.MultiReturn,
			FieldType:  fmt.Sprintf("%s_%s_%s", tableName, viewConfig.Name, PrimaryKeyFieldPlaceholder),
			FieldName:  getPlural(viewConfig.Config.MultiReturn, PrimaryKeyFieldPlaceholder),
			FieldIndex: ResponseStartFieldIndex,
		}
	}
	return &ResponseConfig{Field: responseField}
}

// generates primary key config. doesn't check if primary key object needs to be generated
func getPrimaryKeyConfig(tableName string, viewName string, primaryKeys []*config.Field) *ResponsePKConfig {
	var fields []*ProtoFieldConfig
	name := fmt.Sprintf("%s_%s_%s", tableName, viewName, PrimaryKeyFieldPlaceholder)
	for index, pk := range primaryKeys {
		fields = append(fields, &ProtoFieldConfig{
			FieldName:  pk.FieldName,
			FieldType:  config.GetProtoFieldTypeFor[pk.FieldType],
			FieldIndex: ResponsePKStartIndex + index,
		})
	}
	return &ResponsePKConfig{Name: name, Fields: fields}
}

func getOutputViewProtoFile(tableName string, viewName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewProtoPath, paths.GeneratedFilePath, tableName, viewName)
	return os.Create(outputFilePath)
}

func getPlural(shouldReturnPlural bool, s string) string {
	if shouldReturnPlural {
		return s + "s"
	}
	return s
}
