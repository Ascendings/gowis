package controllers

import (
	"fmt"

	"gopkg.in/macaron.v1"

	"gogs.ballantine.tech/gballan1/gowis/app/forms"
	"gogs.ballantine.tech/gballan1/gowis/models"
)

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

// PostCreate - post route for creating page
func (w WikiController) PostCreate(ctx *macaron.Context, input forms.CreatePageForm) {
	fmt.Println("Here we are!")

	// create the new page model
	models.DB.Create(&models.Page{
		URLSlug:     input.URLSlug,
		PageContent: input.PageContent,
		CreatedBy:   1,
	})

	fmt.Println("And another!")

	// redirect the user
	ctx.Redirect(ctx.URLFor("wiki.list"))
}
