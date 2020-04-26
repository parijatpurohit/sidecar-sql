package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

type EntityProtoConfig struct {
	StorageConfig     *config.StorageConfig
	TableName         string
	HasTimestampField bool
}

func GenerateEntityProto(storageConfig *config.StorageConfig) {
	entityConfig := EntityProtoConfig{StorageConfig: storageConfig}
	for _, field := range storageConfig.Fields {
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			entityConfig.HasTimestampField = true
			break
		}
	}
	entityConfig.TableName = generateUtils.GetTableName(entityConfig.StorageConfig.Table, entityConfig.StorageConfig.Common.IsPlural)

	template := getTemplate(paths.ProtoSchemaTemplateFile)
	outputFile, err := getOutputEntityProtoFile(&entityConfig)
	if err != nil {
		panic(err)
	}
	err = template.Execute(outputFile, entityConfig)
	if err != nil {
		panic(err)
	}
}

func getOutputEntityProtoFile(entityConfig *EntityProtoConfig) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.EntityProtoPath, paths.GeneratedFilePath, entityConfig.TableName)
	return os.Create(outputFilePath)
}
