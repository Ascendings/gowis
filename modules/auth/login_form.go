package auth

import (
	"github.com/astaxie/beego/validation"
	"github.com/fatih/structs"

	"github.com/Ascendings/gowis/modules/web"
)

// LoginForm - form used for creating a page
type LoginForm struct {
	web.Form
	Identifier string `form:"identifier"`
	Password   string `form:"password"`
}

// Validate - validates the form data
func (lf *LoginForm) Validate() {
	// set our fields
	lf.Fields = structs.Fields(&LoginForm{})

	// create new validation object
	lf.Valid = validation.Validation{}

	// add rules
	lf.Valid.Required(lf.Identifier, "identifier").Message("Email or username is required")
	lf.Valid.Required(lf.Password, "password").Message("Your password is required")
}
