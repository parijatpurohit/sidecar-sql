package generateUtils

import (
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
)

func GetTableName(tableName string, isPlural bool) string {
	name := strings.Title(tableName)
	if isPlural {
		return name[:len(tableName)-1]
	}
	return name
}

func GetFieldConfig(storageConfig *config.StorageConfig) (map[string]*config.Field, []*config.Field) {
	fieldSchema := map[string]*config.Field{}
	var primaryKeys []*config.Field
	for _, field := range storageConfig.Fields {
		fieldSchema[field.FieldName] = field
		if field.PrimaryKey {
			primaryKeys = append(primaryKeys, field)
		}
	}
	return fieldSchema, primaryKeys
}
