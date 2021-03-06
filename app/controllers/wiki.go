package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"

	"github.com/Ascendings/gowis/app/models"
	"github.com/Ascendings/gowis/modules/wiki"
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
func (w WikiController) PostCreate(ctx *macaron.Context, input wiki.PageForm, sess session.Store, f *session.Flash, x csrf.CSRF) {
	// validate form Data
	input.Validate()
	// check for validation errors
	if input.HasErrors() {
		// add errors to the view
		errors := input.GetErrors()
		ctx.Data["errors"] = errors

		// pass the user's input back to the view
		ctx.Data["input"] = input

		// let the user know that were some problems with their submission
		f.Error("There were some problems with your submission. Please review your information", true)
		// render the create page view
		w.Render(ctx, "wiki/create")
	} else {
		// Page model
		page := new(models.Page)

		user := sess.Get("user").(models.User)

		// set the page attributes
		page.URLSlug = input.URLSlug
		page.PageContent = input.PageContent
		page.CreatedBy = &user

		// save the page
		pageID, pageErr := models.DB.Insert(page)
		// set page ID value to the returned value from the insert query
		page.ID = int(pageID)
		// check for errors
		if pageErr != nil {
			// flash the error message to the user
			f.Error(pageErr.Error(), false)
			// redirect the user to the home page
			ctx.Redirect(ctx.URLFor("wiki.home"))
		}

		// create new commit model
		commit := models.Commit{}.NewCreateCommit(page.PageContent, input.CommitMessage, page, &user)

		// save the commit
		_, commitErr := models.DB.Insert(commit)
		// check for errors
		if commitErr != nil {
			// flash the error message to the user
			f.Error(commitErr.Error(), false)
			// redirect the user to the home page
			ctx.Redirect(ctx.URLFor("wiki.home"))
		}

		// let the user know we're all good
		f.Success("Your page was created successfully!", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.view", ":urlSlug", page.URLSlug))
	}
}

// View - view a wiki page
func (w WikiController) View(ctx *macaron.Context, f *session.Flash) {
	// Page model
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "url_slug")
	// check for errors
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		// let the user know the page doesn't exist
		f.Info("That page doesn't exist", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.list"))
	} else {
		// load relationship(s)
		models.DB.LoadRelated(&page, "CreatedBy")

		// add the page result to the view
		ctx.Data["page"] = page
		// add converted HTML to view
		ctx.Data["convertedPageContent"] = page.ConvertPageContent()

		// set the title
		ctx.Data["title"] = "View Page | Gowis"
		// render the view
		w.Render(ctx, "wiki/view")
	}
}

// Edit - edit a page
func (w WikiController) Edit(ctx *macaron.Context, f *session.Flash, x csrf.CSRF) {
	// Page model
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "url_slug")
	// check for errors
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		// let the user know the page doesn't exist
		f.Info("That page doesn't exist", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.list"))
	} else {
		// add the page result to the view
		ctx.Data["page"] = page
		ctx.Data["oldslug"] = ctx.Params("urlSlug")

		// set the title
		ctx.Data["title"] = "View Page | Gowis"
		// render the view
		w.Render(ctx, "wiki/edit")
	}
}

// PostEdit - post backend for editing a page
func (w WikiController) PostEdit(ctx *macaron.Context, input wiki.PageForm, f *session.Flash, sess session.Store, x csrf.CSRF) {
	// Page model
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "url_slug")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		// let the user know the page doesn't exist
		f.Info("That page doesn't exist", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.list"))
	} else {
		// save old content for commit record
		oldContent := page.PageContent

		// change the page attributes
		page.URLSlug = input.URLSlug
		page.PageContent = input.PageContent

		// validate form data
		input.Validate()
		// check for validation errors
		if input.HasErrors() {
			errors := input.GetErrors()

			// add errors to the view
			ctx.Data["errors"] = errors

			// add the page and old URL slug to the view
			ctx.Data["page"] = page
			ctx.Data["oldslug"] = ctx.Params("urlSlug")

			w.Render(ctx, "wiki/edit")
		} else {
			// update the record
			_, err = models.DB.Update(&page)
			// check for errors
			if err != nil {
				// flash the error message to the user
				f.Error(err.Error(), false)
				// redirect the user to the home page
				ctx.Redirect(ctx.URLFor("wiki.home"))
			}

			user := sess.Get("user").(models.User)

			// create new commit model
			commit := models.Commit{}.NewEditCommit(oldContent, page.PageContent, input.CommitMessage, &page, &user)

			// save the commit
			_, commitErr := models.DB.Insert(commit)
			// check for errors
			if commitErr != nil {
				// flash the error message to the user
				f.Error(commitErr.Error(), false)
				// redirect the user to the home page
				ctx.Redirect(ctx.URLFor("wiki.home"))
			}

			// redirect the user
			ctx.Redirect(ctx.URLFor("wiki.view", ":urlSlug", page.URLSlug))
		}
	}
}

// Delete - delete a wiki page
func (w WikiController) Delete(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	// find the page
	page := models.Page{URLSlug: ctx.Params("urlSlug")}

	// find the page
	err := models.DB.Read(&page, "url_slug")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		// let the user know the page doesn't exist
		f.Info("That page doesn't exist", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.list"))
	}

	// attempt to delete the page record
	if _, delErr := models.DB.Delete(&page); err != nil {
		// there's an error, let's panic now
		panic(delErr)
	}

	// success! Flash the user and redirect him/her
	f.Success("The page was successfully deleted.", false)
	ctx.Redirect(ctx.URLFor("wiki.list"))
}
