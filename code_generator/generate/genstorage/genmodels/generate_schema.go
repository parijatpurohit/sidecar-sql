package genmodels

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

type GormTags string

const (
	GormTagKey        = "gorm"
	GormTagPrimaryKey = "PRIMARY_KEY"
	GormTagNotNull    = "NOT NULL"
)

func GenerateSchema(storageConfig *config.StorageConfig) {
	tplConfig := SchemaConfig{TableName: generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)}
	for _, field := range storageConfig.Fields {
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			tplConfig.HasTimestampField = true
		}
		fieldCfg := &SchemaFieldConfig{
			Name:      field.FieldName,
			FieldType: config.GetGoFieldTypeFor[field.FieldType],
			TagStr:    GetTagsString(GetSchemaFieldTags(field)),
		}
		tplConfig.Fields = append(tplConfig.Fields, fieldCfg)
	}
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, paths.ModelSchemaTemplateFile))
	outFile, err := getOutputViewSchemaFile(storageConfig.Table)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, tplConfig)
	if err != nil {
		panic(err)
	}
}

func GetSchemaFieldTags(fieldConfig *config.Field) []*FieldTag {
	var tags []*FieldTag
	gormTag := GetGormTag(fieldConfig)
	if gormTag != nil {
		tags = append(tags, gormTag)
	}
	return tags
}

func GetGormTag(fieldConfig *config.Field) *FieldTag {
	gormTag := &FieldTag{TagName: GormTagKey, TagValue: ";"}
	if fieldConfig.PrimaryKey {
		gormTag.TagValue = gormTag.TagValue + GormTagPrimaryKey + ";"
	}
	if fieldConfig.NotNull {
		gormTag.TagValue = gormTag.TagValue + GormTagNotNull + ";"
	}
	if gormTag.TagValue != ";" {
		gormTag.TagValue = strings.Trim(gormTag.TagValue, ";")
		return gormTag
	}
	return nil
}

func getOutputViewSchemaFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.SchemaFilePath, paths.GeneratedFilePath, strings.ToLower(tableName))
	return os.Create(outputFilePath)
}
