package wiki

import (
	"github.com/astaxie/beego/validation"
)

// PageForm - form used for creating a page
type PageForm struct {
	URLSlug       string `form:"url_slug"`
	PageContent   string `form:"page_content"`
	CommitMessage string `form:"commit_message"`
}

// Validate - validates the form data
func (cf *PageForm) Validate() validation.Validation {
	// create new validation object
	valid := validation.Validation{}

	// add rules
	valid.Required(cf.URLSlug, "urlslug").Message("URL slug is required")
	valid.MinSize(cf.URLSlug, 2, "urlslugmin").Message("URL slug must be at least 2 characters long")
	valid.Required(cf.CommitMessage, "commitmessage").Message("A commit message is required")

	// return the validation result
	return valid
}
