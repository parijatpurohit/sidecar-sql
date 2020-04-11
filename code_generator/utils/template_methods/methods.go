package template_methods

import "text/template"

var TemplateMethods template.FuncMap = template.FuncMap{
	"sum": template_sum,
}
