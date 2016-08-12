package controllers

import (
	"gogs.ballantine.tech/gballan1/gowis/lib"

	"gopkg.in/macaron.v1"
)

// Wiki - wiki controller
type Wiki struct {
	*lib.Controller
}

// Home - home page
func (w Wiki) Home(ctx *macaron.Context) {
	// set the page title
	ctx.Data["title"] = "Gowis"
	// render the view
	w.Render(ctx, "wiki/home")
}
