package generateUtils

import (
	"text/template"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
	"github.com/parijatpurohit/sidecar-sql/code_generator/utils/template_methods"
)

func GetTemplate(templateFile string) *template.Template {
	bytes, err := utils.GetFileBytes(templateFile)
	if err != nil {
		panic(err)
	}
	t, err := template.New(templateFile).Funcs(template_methods.TemplateMethods).Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	return t
}
