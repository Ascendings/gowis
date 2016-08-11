package app

import (
	"gogs.ballantine.tech/gballan1/gowis/routers"

	"gopkg.in/macaron.v1"
)

// InitRouter - initializes the router and sets routes
func InitRouter() *macaron.Macaron {
	// create new router
	m := macaron.Classic()

	// define routes
	m.Get("/", routers.WikiHome)

	// return the compeleted router
	return m
}
