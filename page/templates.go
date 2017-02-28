package page

import (
	"fmt"
	"html/template"
)

type templateMap map[string]template.Template

var templates templateMap

func init() {
	templates = make(templateMap)
}

func TemplateSet(name string) *template.Template {
	if ret, ok := templates[name]; ok {
		return &ret
	} else {
		return nil
	}
}

func TemplateSubdirMust(name string) {
	tpl := template.New(name)

	path := fmt.Sprintf("%s/*.html", name)
	template.Must(tpl.ParseGlob(path))

	templates[name] = *tpl
}
