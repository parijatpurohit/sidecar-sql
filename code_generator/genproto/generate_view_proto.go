package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
)

const (
	viewProtoPath    = "%s/pb/storage/%s_%s_View.proto"
	viewTemplateFile = "view.tgo"
)

type TemplateQueryField struct {
	QueryField  *config.QueryField
	FieldSchema *config.Field
	IsArray     bool
}

type QueryParams struct {
	Fields []*TemplateQueryField
}

type ViewProtoConfig struct {
	QueryParams
	*config.View
	TableName         string
	PrimaryKeys       []*config.Field
	HasTimestampField bool
}

func GenerateViewProto(storageConfig *config.StorageConfig) {
	fieldSchema := map[string]*config.Field{}
	var primaryKeys []*config.Field
	for _, field := range storageConfig.Fields {
		fieldSchema[field.FieldName] = field
		if field.PrimaryKey {
			primaryKeys = append(primaryKeys, field)
		}
	}
	for _, view := range storageConfig.Views {
		viewConfig := ViewProtoConfig{View: view, PrimaryKeys: primaryKeys}
		viewConfig.TableName = GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
		if view.Query != nil {
			for _, field := range view.Query.Fields {
				queryField := TemplateQueryField{
					QueryField:  field,
					FieldSchema: fieldSchema[field.FieldName],
					IsArray:     field.SearchType == config.SEARCH_TYPE_CONTAINS,
				}
				viewConfig.QueryParams.Fields = append(viewConfig.QueryParams.Fields, &queryField)
				if queryField.FieldSchema.FieldType == config.FIELD_TYPE_TIMESTAMP {
					viewConfig.HasTimestampField = true
				}
			}
		}
		template := getTemplate(viewTemplateFile)
		outputFile, err := getOutputViewProtoFile(&viewConfig, view.Name)
		if err != nil {
			panic(err)
		}
		err = template.Execute(outputFile, viewConfig)
		if err != nil {
			panic(err)
		}
	}

}

func getOutputViewProtoFile(viewConfig *ViewProtoConfig, viewName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(viewProtoPath, config.GeneratedFilePath, viewConfig.TableName, viewName)
	return os.Create(outputFilePath)
}
