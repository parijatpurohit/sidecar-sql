package generateUtils

import (
	"strings"

	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func GetTableName(tableName string, isPlural bool) string {
	name := strings.Title(tableName)
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
