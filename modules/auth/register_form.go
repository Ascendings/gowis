package auth

import (
	"regexp"

	"github.com/astaxie/beego/validation"
	"github.com/fatih/structs"

	"github.com/Ascendings/gowis/modules/web"
)

// RegisterForm - form used for creating a page
type RegisterForm struct {
	web.Form
	Email         string `form:"email"`
	EmailAgain    string `form:"email_again"`
	Username      string `form:"username"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password_again"`
	FirstName     string `form:"first_name"`
	LastName      string `form:"last_name"`
}

// Validate - validates the form data
func (rf *RegisterForm) Validate() {
	// set our fields
	rf.Fields = structs.Fields(&RegisterForm{})

	// create new validation object
	rf.Valid = validation.Validation{}

	// create regexp objects for matches
	emailReg := regexp.MustCompile(rf.Email)
	passwordReg := regexp.MustCompile(rf.Password)

	// add rules
	rf.Valid.Required(rf.Email, "email").Message("An email address is required")
	rf.Valid.Email(rf.Email, "email").Message("Your email address isn't valid")
	rf.Valid.Match(rf.EmailAgain, emailReg, "email_again").Message("Your emails must match")
	rf.Valid.Required(rf.Username, "username").Message("A username is required")
	rf.Valid.AlphaNumeric(rf.Username, "username").Message("Your username may only contain alpha-numeric characters")
	rf.Valid.Required(rf.Password, "password").Message("A password is required")
	rf.Valid.MinSize(rf.Password, 6, "password").Message("Your password must be at least 6 characters long")
	rf.Valid.Match(rf.PasswordAgain, passwordReg, "password_again").Message("Your passwords must match")
}
