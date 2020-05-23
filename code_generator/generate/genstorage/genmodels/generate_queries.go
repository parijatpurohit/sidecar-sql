package genmodels

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
	TagOmitEmpty      = "omitempty"
	TagValueSeparator = ","
)

func GenerateQueryModels(storageConfig *config.StorageConfig) {
	for _, view := range storageConfig.Views {
		if view.Query != nil {
			log.Printf("generating query models for: %s", storageConfig.Table)
			generateQueryModel(storageConfig, view)
		}
	}
}

func generateQueryModel(storageConfig *config.StorageConfig, view *config.View) {
	conf := getQueryConfig(storageConfig, view)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, paths.ModelQueryTemplateFile))
	outFile, err := getOutputViewQueryFile(storageConfig.Table, view.Name)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getQueryConfig(storageConfig *config.StorageConfig, view *config.View) *QueryConfig {
	fieldSchema, _ := generateUtils.GetFieldConfig(storageConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	return &QueryConfig{
		PackageName: fmt.Sprintf(SchemaPackageFormat, strings.ToLower(tableName)),
		Imports:     getQueryImports(view.Query.Fields, fieldSchema),
		TableName:   tableName,
		ViewName:    view.Name,
		Fields:      getQueryFieldConfig(view.Query.Fields, fieldSchema),
	}
}

func getQueryFieldConfig(fields []*config.QueryField, fieldSchema map[string]*config.Field) []*FieldConfig {
	var queryFieldConfig []*FieldConfig
	for _, queryField := range fields {
		field := fieldSchema[queryField.FieldName]
		queryFieldCfg := &FieldConfig{
			Name:       field.FieldName,
			FieldType:  config.GetGoFieldTypeFor[field.FieldType],
			IsRepeated: queryField.SearchType == config.SEARCH_TYPE_CONTAINS,
			TagStr:     GetTagsString(GetQueryFieldTags(queryField)),
		}
		queryFieldConfig = append(queryFieldConfig, queryFieldCfg)
	}
	return queryFieldConfig
}

func getQueryImports(fields []*config.QueryField, fieldSchema map[string]*config.Field) []string {
	var imports []string
	for _, queryField := range fields {
		field := fieldSchema[queryField.FieldName]
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			imports = append(imports, ImportTimestamp)
		}
	}
	return utils.GetUnique(imports)
}

func GetQueryFieldTags(queryField *config.QueryField) []*FieldTag {
	var tags []*FieldTag
	tags = append(tags, GetQueryJSONTag(queryField))
	return tags
}

func GetQueryJSONTag(queryField *config.QueryField) *FieldTag {
	tag := &FieldTag{TagName: "json", TagValue: TagValueSeparator}
	tag.TagValue = strings.Title(queryField.FieldName) + TagValueSeparator
	if queryField.IsOptional {
		tag.TagValue = tag.TagValue + TagOmitEmpty + TagValueSeparator
	}
	tag.TagValue = strings.Trim(tag.TagValue, TagValueSeparator)
	if tag.TagValue == "" {
		return nil
	}
	return tag
}

func getOutputViewQueryFile(tableName string, queryName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.QueryFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), queryName)
	return os.Create(outputFilePath)
}
