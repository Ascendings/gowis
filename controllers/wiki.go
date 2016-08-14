package controllers

import (
	macaron "gopkg.in/macaron.v1"

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
	// pages array
	var pages []models.Page

	// get pages from the DB
	qs := models.DB.QueryTable("page")
	// order the results and put them into the array
	qs.OrderBy("-created_at").All(&pages)

	// add the pages to the view context
	ctx.Data["pages"] = pages

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
	// Page model
	page := new(models.Page)

	// set the page attributes
	page.URLSlug = input.URLSlug
	page.PageContent = input.PageContent
	page.CreatedBy = 1

	// save the page
	_, err := models.DB.Insert(page)
	// check for errors
	if err != nil {
		panic(err)
	}

	// redirect the user
	ctx.Redirect(ctx.URLFor("wiki.list"))
}
