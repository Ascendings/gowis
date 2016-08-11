package app

import (
	"gogs.ballantine.tech/gballan1/gowis/routers"

	"gopkg.in/macaron.v1"
)

// InitRouter - initializes the router and sets routes
func InitRouter(m macaron.Macaron) {
	// define routes
	m.Get("/", routers.WikiHome)
}
