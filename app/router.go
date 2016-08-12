package app

import (
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
}
