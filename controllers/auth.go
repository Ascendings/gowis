package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"

	"gogs.ballantine.tech/gballan1/gowis/models"
	"gogs.ballantine.tech/gballan1/gowis/modules/auth"

	macaron "gopkg.in/macaron.v1"
)

// AuthController - authentication controller
type AuthController struct {
	*Controller
}

// Register - user registration view
func (a *AuthController) Register(ctx *macaron.Context, x csrf.CSRF) {
	// add the CSRF token to the view
	ctx.Data["csrf_token"] = a.CreateCsrfField(x)

	// set the page title
	ctx.Data["title"] = "Register | Gowis"
	// render the view
	a.Render(ctx, "auth/register")
}

// PostRegister - user registration post stuff
func (a *AuthController) PostRegister(ctx *macaron.Context, input auth.RegisterForm, f *session.Flash, x csrf.CSRF) {
	// validate form Data
	input.Validate()
	// check for validation errors
	if input.HasErrors() {
		// add errors back to view
		errors := input.GetErrors()
		ctx.Data["errors"] = errors

		// pass the user's input back to the view
		ctx.Data["input"] = input

		// pass a new CSRF token
		ctx.Data["csrf_token"] = a.CreateCsrfField(x)

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
func (a *AuthController) Login(ctx *macaron.Context, x csrf.CSRF) {
	// add the CSRF token to the view
	ctx.Data["csrf_token"] = a.CreateCsrfField(x)

	// set the page title
	ctx.Data["title"] = "Login | Gowis"
	// render the view
	a.Render(ctx, "auth/login")
}

// PostLogin - login backend
func (a *AuthController) PostLogin(ctx *macaron.Context, input auth.LoginForm, f *session.Flash, sess session.Store) {
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
		// set the user id in the session - the user object will be instated on next request
		sess.Set("user_id", user.ID)

		// flash a success message
		f.Success("You have logged in successfully!", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.list"))
	}
}

// Logout - log the user out
func (a *AuthController) Logout(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	// unset the session
	if sess.Get("user_id") != nil || sess.Get("user") != nil {
		// unset the stuff
		sess.Set("user_id", nil)
		sess.Set("user", nil)

		// flash message to user
		f.Info("You have been successfully logged out")
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.home"))
	}
}
