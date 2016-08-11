package app

import (
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"
)

// InitTemplates - sets up the Macaron router with the Pongo2 template engine
func InitTemplates(m macaron.Macaron) {
	// setup the Pongo2 template engine
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "views",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}))
}
