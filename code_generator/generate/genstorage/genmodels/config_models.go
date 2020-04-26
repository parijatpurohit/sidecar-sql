package genmodels

import "fmt"

type FieldTag struct {
	TagName  string
	TagValue string
}

type SchemaFieldConfig struct {
	Name      string
	FieldType string
	TagStr    string
}

type QueryFieldConfig struct {
	Name       string
	IsRepeated bool
	FieldType  string
	TagStr     string
}

type SchemaConfig struct {
	TableName         string
	Fields            []*SchemaFieldConfig
	HasTimestampField bool
}

type QueryConfig struct {
	TableName         string
	ViewName          string
	Fields            []*QueryFieldConfig
	HasTimestampField bool
}

func GetTagsString(tags []*FieldTag) string {
	tagsStr := ""
	for _, tag := range tags {
		tagsStr = tagsStr + fmt.Sprintf("%s:\"%s\"", tag.TagName, tag.TagValue)
	}
	return tagsStr

}
