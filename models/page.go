package models

import (
	"fmt"

	"github.com/golang-commonmark/markdown"
	"github.com/jinzhu/gorm"
)

// Page - wiki page model
type Page struct {
	gorm.Model
	URLSlug     string
	PageContent string
	CreatedBy   int
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
