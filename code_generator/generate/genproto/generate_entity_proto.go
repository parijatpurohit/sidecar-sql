package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/alias"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	ImportTimestampProto = "google/protobuf/timestamp.proto"
)

type EntityProtoConfig struct {
	StorageConfig *config.StorageConfig
	TableName     string
	Fields        []*ProtoFieldConfig
	Imports       []string
}

type ProtoFieldConfig struct {
	IsRepeated bool
	FieldName  string
	FieldType  string
	FieldIndex int
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
	entityConfig := &EntityProtoConfig{
		StorageConfig: storageConfig,
		Fields:        getProtoFields(storageConfig),
		TableName:     generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural),
	}
	for _, field := range storageConfig.Fields {
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			entityConfig.Imports = append(entityConfig.Imports, ImportTimestampProto)
		}
	}
	entityConfig.Imports = utils.GetUnique(entityConfig.Imports)
	return entityConfig
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
