package template_methods

import (
	"strings"
	"text/template"
)

var TemplateMethods template.FuncMap = template.FuncMap{
	"sum":     template_sum,
	"toLower": strings.ToLower,
}
