package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/alias"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	ImportTimestampProto = "google/protobuf/timestamp.proto"
)

type EntityProtoConfig struct {
	StorageConfig *config.StorageConfig
	TableName     string
	Fields        []*ProtoFieldConfig
	PKConfig      *PKConfig
	Imports       []string
	GoPackagePath string
}

type ProtoFieldConfig struct {
	IsRepeated bool
	FieldName  string
	FieldType  string
	FieldIndex int
}

type PKConfig struct {
	Name   string
	Fields []*ProtoFieldConfig
}

func GenerateEntityProto(storageConfig *config.StorageConfig) {
	// build config for template
	entityConfig := getEntityConfig(storageConfig)

	// get template file
	template := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ProtoTemplatePath, paths.ProtoSchemaTemplateFile))

	// get or create output file path
	outputFile, err := getOutputEntityProtoFile(entityConfig.TableName)
	if err != nil {
		panic(err)
	}

	// populate file with executed template
	err = template.Execute(outputFile, entityConfig)
	if err != nil {
		panic(err)
	}
}

func getEntityConfig(storageConfig *config.StorageConfig) *EntityProtoConfig {
	_, primaryKeys := generateUtils.GetFieldConfig(storageConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	entityConfig := &EntityProtoConfig{
		StorageConfig: storageConfig,
		Fields:        getProtoFields(storageConfig),
		TableName:     tableName,
		PKConfig:      getPrimaryKeyConfig(tableName, primaryKeys),
		GoPackagePath: paths.GoPackagePath,
	}
	for _, field := range storageConfig.Fields {
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			entityConfig.Imports = append(entityConfig.Imports, ImportTimestampProto)
		}
	}
	entityConfig.Imports = utils.GetUnique(entityConfig.Imports)
	return entityConfig
}

// generates primary key config. doesn't check if primary key object needs to be generated
func getPrimaryKeyConfig(tableName string, primaryKeys []*config.Field) *PKConfig {
	var fields []*ProtoFieldConfig
	name := fmt.Sprintf("%s_%s", tableName, PrimaryKeyFieldPlaceholder)
	for index, pk := range primaryKeys {
		fields = append(fields, &ProtoFieldConfig{
			FieldName:  pk.FieldName,
			FieldType:  alias.GetProtoFieldTypeFor[pk.FieldType],
			FieldIndex: ResponsePKStartIndex + index,
		})
	}
	return &PKConfig{Name: name, Fields: fields}
}

func getProtoFields(storageConfig *config.StorageConfig) []*ProtoFieldConfig {
	var fieldsConfig []*ProtoFieldConfig
	for index, field := range storageConfig.Fields {
		fieldsConfig = append(fieldsConfig, &ProtoFieldConfig{
			FieldType:  alias.GetProtoFieldTypeFor[field.FieldType],
			FieldName:  field.FieldName,
			FieldIndex: index + 1,
		})
	}
	return fieldsConfig
}

func getOutputEntityProtoFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.EntityProtoPath, paths.GeneratedFilePath, tableName)
	return os.Create(outputFilePath)
}
