package wiki

import (
	"github.com/astaxie/beego/validation"
	"github.com/fatih/structs"

	"github.com/Ascendings/gowis/modules/web"
)

// PageForm - form used for creating a page
type PageForm struct {
	web.Form
	URLSlug       string `form:"url_slug"`
	PageContent   string `form:"page_content"`
	CommitMessage string `form:"commit_message"`
}

// Validate - validates the form data
func (cf *PageForm) Validate() {
	// set our fields
	cf.Fields = structs.Fields(&PageForm{})

	// create new validation object
	cf.Valid = validation.Validation{}

	// add rules
	cf.Valid.Required(cf.URLSlug, "urlslug").Message("URL slug is required")
	cf.Valid.MinSize(cf.URLSlug, 2, "urlslugmin").Message("URL slug must be at least 2 characters long")
	cf.Valid.Required(cf.CommitMessage, "commitmessage").Message("A commit message is required")
}
