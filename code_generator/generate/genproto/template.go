package genproto

import (
	"fmt"
	"text/template"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils/template_methods"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
)

func getTemplate(templateFile string) *template.Template {
	bytes, err := utils.GetFileBytes(fmt.Sprintf("%s/%s", paths.TemplatePath, templateFile))
	if err != nil {
		panic(err)
	}
	t, err := template.New(templateFile).Funcs(template_methods.TemplateMethods).Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	return t
}
