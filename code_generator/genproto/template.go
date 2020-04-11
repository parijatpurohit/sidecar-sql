package genproto

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils/template_methods"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
)

const (
	templatePath      = "templates/proto"
	basicTemplateFile = "basic.tgo"
)

func getTemplate(templateFile string) *template.Template {
	bytes, err := utils.GetFileBytes(fmt.Sprintf("%s/%s", templatePath, templateFile))
	if err != nil {
		panic(err)
	}
	t, err := template.New(templateFile).Funcs(template_methods.TemplateMethods).Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	return t
}

func GetTableName(tableName string, isPlural bool) string {
	if isPlural {
		return strings.Title(tableName[:len(tableName)-1])
	}
	return tableName
}
