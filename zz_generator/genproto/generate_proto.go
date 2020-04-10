package genproto

import (
	"fmt"
	"text/template"

	"github.com/parijatpurohit/sidecar-sql/zz_generator/utils"
)

const (
	templatePath       = "templates/proto"
	basicTemplateFile  = "basic.tgo"
	schemaTemplateFile = "schema.tgo"
)

func getTemplate(templateFile string) *template.Template {
	bytes, err := utils.GetFileBytes(fmt.Sprintf("%s/%s", templatePath, templateFile))
	if err != nil {
		panic(err)
	}
	t, err := template.New(templateFile).Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	return t
}
