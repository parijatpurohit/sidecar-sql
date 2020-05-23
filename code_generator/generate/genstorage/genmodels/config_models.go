package genmodels

import "fmt"

type FieldTag struct {
	TagName  string
	TagValue string
}

type FieldConfig struct {
	Name       string
	IsRepeated bool
	FieldType  string
	TagStr     string
}

type SchemaConfig struct {
	PackageName string
	Imports     []string
	TableName   string
	Fields      []*FieldConfig
}

type QueryConfig struct {
	PackageName string
	Imports     []string
	TableName   string
	ViewName    string
	Fields      []*FieldConfig
}

func GetTagsString(tags []*FieldTag) string {
	tagsStr := ""
	for _, tag := range tags {
		tagsStr = tagsStr + fmt.Sprintf("%s:\"%s\"", tag.TagName, tag.TagValue)
	}
	return tagsStr

}
