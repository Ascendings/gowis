package template

import (
	"html/template"
	"runtime"
	"strings"

	"gopkg.in/macaron.v1"
)

func NewFuncMap(macaron *macaron.Macaron) []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		// Go version
		"GoVer": func() string {
			return strings.Title(runtime.Version())
		},
		// Application name
		"AppName": func() string {
			return "Macaron"
		},
		// application version
		"AppVer": func() string {
			return "1.0.0"
		},
		// construct URLs in views
		"URLFor": macaron.URLFor,
		// output raw HTML
		"raw": func(text string) template.HTML {
			return template.HTML(text)
		},
	}}
}
