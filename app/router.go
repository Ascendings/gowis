package app

import (
	"github.com/go-macaron/binding"

	"gogs.ballantine.tech/gballan1/gowis/app/forms"
	"gogs.ballantine.tech/gballan1/gowis/controllers"

	"gopkg.in/macaron.v1"
)

// InitRouter - initializes the router and sets routes
func InitRouter(m macaron.Macaron) {
	// create new Wiki controller
	w := new(controllers.WikiController)

	// define routes
	m.Get("/", w.Home).Name("wiki.home")
	m.Get("/list", w.List).Name("wiki.list")
	m.Get("/create", w.Create).Name("wiki.create")
	m.Post("/create", binding.Bind(forms.CreatePageForm{}), w.PostCreate)
	m.Get("/view/:urlSlug", w.View).Name("wiki.view")
}
