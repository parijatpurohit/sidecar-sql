package generateUtils

import "strings"

func GetTableName(tableName string, isPlural bool) string {
	name := strings.Title(tableName)
	if isPlural {
		return name[:len(tableName)-1]
	}
	return name
}
