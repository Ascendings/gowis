package auth

import (
  "github.com/astaxie/beego/validation"
  "github.com/fatih/structs"

  "gogs.ballantine.tech/gballan1/gowis/modules/web"
)

// LoginForm - form used for creating a page
type LoginForm struct {
  web.Form
  indetifier string `form:"identifier"`
  password   string `form:"password"`
}

// Validate - validates the form data
func (lf *LoginForm) Validate() {
  // set our fields
  lf.Fields = structs.Fields(&LoginForm{})

  // create new validation object
  lf.Valid = validation.Validation{}

  // add rules
  lf.Valid.Required(lf.URLSlug, "urlslug").Message("URL slug is required")
  lf.Valid.Required(lf.CommitMessage, "commitmessage").Message("A commit message is required")
}
