package genmodels

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	TagOmitEmpty = "omitempty"
	TagSeparator = ","
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
	fieldSchema, _ := generateUtils.GetFieldConfig(storageConfig)
	tplConfig := QueryConfig{
		TableName: generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural),
		ViewName:  view.Name,
	}
	for _, queryField := range view.Query.Fields {
		field := fieldSchema[queryField.FieldName]
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			tplConfig.HasTimestampField = true
		}
		queryFieldCfg := &QueryFieldConfig{
			Name:       field.FieldName,
			FieldType:  config.GetGoFieldTypeFor[field.FieldType],
			IsRepeated: queryField.SearchType == config.SEARCH_TYPE_CONTAINS,
			TagStr:     GetTagsString(GetQueryFieldTags(queryField)),
		}
		tplConfig.Fields = append(tplConfig.Fields, queryFieldCfg)
	}
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, paths.ModelQueryTemplateFile))
	outFile, err := getOutputViewQueryFile(storageConfig.Table, view.Name)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, tplConfig)
	if err != nil {
		panic(err)
	}
}

func GetQueryFieldTags(queryField *config.QueryField) []*FieldTag {
	var tags []*FieldTag
	tags = append(tags, GetQueryJSONTag(queryField))
	return tags
}

func GetQueryJSONTag(queryField *config.QueryField) *FieldTag {
	tag := &FieldTag{TagName: "json", TagValue: ","}
	tag.TagValue = strings.Title(queryField.FieldName) + TagSeparator
	if queryField.IsOptional {
		tag.TagValue = tag.TagValue + TagOmitEmpty + TagSeparator
	}
	tag.TagValue = strings.Trim(tag.TagValue, TagSeparator)
	if tag.TagValue == "" {
		return nil
	}
	return tag
}

func getOutputViewQueryFile(tableName string, queryName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.QueryFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), queryName)
	return os.Create(outputFilePath)
}
