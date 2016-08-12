package controllers

import "gopkg.in/macaron.v1"

// WikiController - wiki controller
type WikiController struct {
	*Controller
}

// Home - home page
func (w WikiController) Home(ctx *macaron.Context) {
	// set the page title
	ctx.Data["title"] = "Gowis"
	// render the view
	w.Render(ctx, "wiki/home")
}

// List - wiki pages
func (w WikiController) List(ctx *macaron.Context) {
	// set the title
	ctx.Data["title"] = "List of Pages | Gowis"
	// render view
	w.Render(ctx, "wiki/list")
}

// Create - create new wiki page
func (w WikiController) Create(ctx *macaron.Context) {
	// set the title
	ctx.Data["title"] = "Create New Page | Gowis"
	// render view
	w.Render(ctx, "wiki/create")
}
