package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang-commonmark/markdown"
)

// Page - wiki page model
type Page struct {
	ID          int    `orm:"pk;auto;column(id)"`
	URLSlug     string `orm:"unique;column(url_slug)"`
	PageContent string `orm:"type(text)"`
	CreatedBy   int
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`

	// page has many commits
	Commits []*Commit `orm:"reverse(many)"`
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

// registers model with DB
func init() {
	orm.RegisterModel(new(Page))
}
