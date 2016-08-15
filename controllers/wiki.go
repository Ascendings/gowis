package controllers

import (
	"github.com/astaxie/beego/orm"
	macaron "gopkg.in/macaron.v1"

	"gogs.ballantine.tech/gballan1/gowis/models"
	"gogs.ballantine.tech/gballan1/gowis/modules/wiki"
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
func (w WikiController) PostCreate(ctx *macaron.Context, input wiki.PageForm) {
	// Page model
	page := new(models.Page)

	// validate form Data
	input.Validate()
	// check for validation errors
	if input.HasErrors() {
		errors := input.GetErrors()

		ctx.Data["errors"] = errors

		w.Render(ctx, "wiki/create")
	} else {
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
}

// View - view a wiki page
func (w WikiController) View(ctx *macaron.Context) {
	// Page model
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "u_r_l_slug")
	// check for errors
	if err == orm.ErrNoRows {
		panic("No result found.")
	} else if err == orm.ErrMissPK {
		panic("No primary key found.")
	}

	// add the page result to the view
	ctx.Data["page"] = page
	// add converted HTML to view
	ctx.Data["convertedPageContent"] = page.ConvertPageContent()

	// set the title
	ctx.Data["title"] = "View Page | Gowis"
	// render the view
	w.Render(ctx, "wiki/view")
}

// Edit - edit a page
func (w WikiController) Edit(ctx *macaron.Context) {
	// Page model
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "u_r_l_slug")
	// check for errors
	if err == orm.ErrNoRows {
		panic("No result found.")
	} else if err == orm.ErrMissPK {
		panic("No primary key found.")
	}

	// add the page result to the view
	ctx.Data["page"] = page

	// set the title
	ctx.Data["title"] = "View Page | Gowis"
	// render the view
	w.Render(ctx, "wiki/edit")
}

// PostEdit - post backend for editing a page
func (w WikiController) PostEdit(ctx *macaron.Context, input wiki.PageForm) {
	// Page model
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "u_r_l_slug")
	// check for errors
	if err == orm.ErrNoRows {
		panic("No result found.")
	} else if err == orm.ErrMissPK {
		panic("No primary key found.")
	}

	// change the page attributes
	page.URLSlug = input.URLSlug
	page.PageContent = input.PageContent

	// update the record
	_, err = models.DB.Update(&page)
	// check for errors
	if err != nil {
		panic(err)
	}

	// redirect the user
	ctx.Redirect(ctx.URLFor("wiki.view", ":urlSlug", page.URLSlug))
}
