package genmodels

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/alias"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

type GormTags string

const (
	SchemaPackageFormat = "%s_models"
	ImportTimestamp     = "time"
	GormTagKey          = "gorm"
	GormTagPrimaryKey   = "PRIMARY_KEY"
	GormTagNotNull      = "NOT NULL"
	TagSeparator        = ";"
)

func GenerateSchema(storageConfig *config.StorageConfig) {
	conf := getSchemaConfig(storageConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, paths.ModelSchemaTemplateFile))
	outFile, err := getOutputViewSchemaFile(tableName)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getSchemaConfig(storageConfig *config.StorageConfig) *SchemaConfig {
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	return &SchemaConfig{
		PackageName: fmt.Sprintf(SchemaPackageFormat, strings.ToLower(tableName)),
		Imports:     getImports(storageConfig.Fields),
		TableName:   tableName,
		Fields:      getSchemaFieldsConfig(storageConfig.Fields),
	}
}

func getSchemaFieldsConfig(fields []*config.Field) []*FieldConfig {
	var fieldConfig []*FieldConfig
	for _, field := range fields {
		fieldCfg := &FieldConfig{
			Name:      field.FieldName,
			FieldType: alias.GetGoFieldTypeFor[field.FieldType],
			TagStr:    GetTagsString(getSchemaFieldTags(field)),
		}
		fieldConfig = append(fieldConfig, fieldCfg)
	}
	return fieldConfig
}

func getImports(fields []*config.Field) []string {
	var imports []string
	for _, field := range fields {
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			imports = append(imports, ImportTimestamp)
		}
	}
	return utils.GetUnique(imports)
}

func getSchemaFieldTags(fieldConfig *config.Field) []*FieldTag {
	var tags []*FieldTag
	gormTag := getGormTag(fieldConfig)
	if gormTag != nil {
		tags = append(tags, gormTag)
	}
	return tags
}

func getGormTag(fieldConfig *config.Field) *FieldTag {
	gormTag := &FieldTag{TagName: GormTagKey, TagValue: TagSeparator}
	if fieldConfig.PrimaryKey {
		gormTag.TagValue = gormTag.TagValue + GormTagPrimaryKey + TagSeparator
	}
	if fieldConfig.NotNull {
		gormTag.TagValue = gormTag.TagValue + GormTagNotNull + TagSeparator
	}
	if gormTag.TagValue != TagSeparator {
		gormTag.TagValue = strings.Trim(gormTag.TagValue, TagSeparator)
		return gormTag
	}
	return nil
}

func getOutputViewSchemaFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.SchemaFilePath, *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, strings.ToLower(tableName))
	return os.Create(outputFilePath)
}
