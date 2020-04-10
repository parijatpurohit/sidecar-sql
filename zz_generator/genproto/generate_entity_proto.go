package genproto

import (
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/zz_generator/config"
)

type EntityProtoConfig struct {
	StorageConfig     *config.StorageConfig
	TableName         string
	HasTimestampField bool
}

func GenerateEntityProto(storageConfig *config.StorageConfig, fileName string) {
	entityConfig := EntityProtoConfig{StorageConfig: storageConfig}
	for _, field := range storageConfig.Fields {
		if field.FieldType == config.FIELD_TYPE_TIMESTAMP {
			entityConfig.HasTimestampField = true
			break
		}
	}
	entityConfig.TableName = entityConfig.StorageConfig.Table
	if entityConfig.StorageConfig.Common.IsPlural {
		entityConfig.TableName = strings.Title(entityConfig.TableName[:len(entityConfig.TableName)-1])

	}
	template := getTemplate(schemaTemplateFile)
	err := template.Execute(os.Stdout, entityConfig)
	if err != nil {
		panic(err)
	}
}
