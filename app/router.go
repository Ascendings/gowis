package app

import (
	"github.com/go-macaron/binding"

	"gogs.ballantine.tech/gballan1/gowis/controllers"
	"gogs.ballantine.tech/gballan1/gowis/modules/wiki"

	"gopkg.in/macaron.v1"
)

// InitRouter - initializes the router and sets routes
func InitRouter(m macaron.Macaron) {
	bindIgnErr := binding.BindIgnErr

	// create new Wiki controller
	w := new(controllers.WikiController)

	// define routes
	m.Get("/", w.Home).Name("wiki.home")
	m.Get("/list", w.List).Name("wiki.list")
	m.Combo("/create").Get(w.Create).Post(bindIgnErr(wiki.PageForm{}), w.PostCreate).Name("wiki.create")
	m.Get("/view/:urlSlug", w.View).Name("wiki.view")
	m.Combo("/edit/:urlSlug").Get(w.Edit).Post(bindIgnErr(wiki.PageForm{}), w.PostEdit).Name("wiki.edit")
}
