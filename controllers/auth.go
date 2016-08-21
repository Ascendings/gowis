package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-macaron/session"

	"gogs.ballantine.tech/gballan1/gowis/models"
	"gogs.ballantine.tech/gballan1/gowis/modules/auth"

	"gopkg.in/macaron.v1"
)

// AuthController - authentication controller
type AuthController struct {
	*Controller
}

// Register - user registration view
func (a *AuthController) Register(ctx *macaron.Context) {
	// set the page title
	ctx.Data["title"] = "Register | Gowis"
	// render the view
	a.Render(ctx, "auth/register")
}

// PostRegister - user registration post stuff
func (a *AuthController) PostRegister(ctx *macaron.Context, input auth.RegisterForm, f *session.Flash) {
	// validate form Data
	input.Validate()
	// check for validation errors
	if input.HasErrors() {
		errors := input.GetErrors()
		ctx.Data["errors"] = errors
		ctx.Data["input"] = input

		// let the user know that were some problems with their submission
		f.Error("There were some problems with your submission. Please review your information", true)
		// render the create page view
		a.Render(ctx, "auth/register")
	} else {
		// User model
		user := new(models.User)

		// create password hash
		passwordHash := user.HashPassword(input.Password)

		// set the user attributes
		user.Email = input.Email
		user.Username = input.Username
		user.Password = passwordHash
		user.FirstName = input.FirstName
		user.LastName = input.LastName

		// save the user
		_, err := models.DB.Insert(user)
		// check for errors
		if err != nil {
			// flash the error message to the user
			f.Error(err.Error(), false)
			// redirect the user to the registration page
			ctx.Redirect(ctx.URLFor("auth.register"))
		}

		// let the user know we're all good
		f.Success("Your account was created successfully!", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("auth.login"))
	}
}

// Login - user login view
func (a *AuthController) Login(ctx *macaron.Context) {
	// set the page title
	ctx.Data["title"] = "Login | Gowis"
	// render the view
	a.Render(ctx, "auth/login")
}

// PostLogin - login backend
func (a *AuthController) PostLogin(ctx *macaron.Context, input auth.LoginForm, f *session.Flash) {
	// boolean to determine whether or not login was successful
	fail := false

	// User model with the email set
	user := models.User{Email: input.Identifier}

	// find the user by email
	emailErr := models.DB.Read(&user, "email")
	// check for errors
	if emailErr == orm.ErrNoRows || emailErr == orm.ErrMissPK {
		// user model with the username set
		user = models.User{Username: input.Identifier}

		// find the user by username
		usernameErr := models.DB.Read(&user, "username")
		// check for errors
		if usernameErr == orm.ErrNoRows || usernameErr == orm.ErrMissPK {
			// set fail to true
			fail = true
		}
	}

	if !fail {
		fail = !user.CheckPassword(input.Password)
	}

	if fail {
		// flash failure message to the user
		f.Error("Invalid credentials", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("auth.login"))
	} else {
		// flash a success message
		f.Success("You have logged in successfully!", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.list"))
	}
}
