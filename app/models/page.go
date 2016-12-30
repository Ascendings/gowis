package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang-commonmark/markdown"
)

// Page - wiki page model
type Page struct {
	ID          int       `orm:"pk;auto;column(id)"`
	URLSlug     string    `orm:"unique;column(url_slug)"`
	PageContent string    `orm:"type(text)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`

	// page has many commits
	Commits []*Commit `orm:"reverse(many)"`

	// page has one creator
	CreatedBy *User `orm:"rel(fk);on_delete(do_nothing);null"`
}

// String - string representation of page
func (p *Page) String() string {
	return fmt.Sprintf("Page ID: %d  Page Slug: %s", p.ID, p.URLSlug)
}

// ConvertPageContent - converts the markdown content into HTML
func (p *Page) ConvertPageContent() string {
	// create new markdown object
	md := markdown.New(
		markdown.Tables(true),
		markdown.Breaks(true))
	// return converted converted content
	return md.RenderToString([]byte(p.PageContent))
}

// LastEditedBy - returns the user model that last edited this page
func (p Page) LastEditedBy() *User {
	// load commit relationship
	DB.LoadRelated(&p, "Commits")

	// retrieve latest commit
	lastCommit := p.Commits[0]

	// load the user relationship for the commit
	DB.LoadRelated(lastCommit, "User")

	return lastCommit.User
}

// registers model with DB
func init() {
	orm.RegisterModel(new(Page))
}
