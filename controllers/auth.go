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

// Login - user login view
func (a *AuthController) Login(ctx *macaron.Context) {
  // set the page title
  ctx.Data["title"] = "Login | Gowis"
  // render the view
  a.Render(ctx, "auth/login")
}

// PostLogin - login backend
func (a *AuthController) PostLogin(ctx *macaron.Context, input auth.LoginForm, f *session.Flash) {
  // User model with the email set
  user := models.User{Email: ctx.Params("identifier")}

  // find the user by email
  emailErr := models.DB.Read(&user, "email")
  // check for errors
  if emailErr == orm.ErrNoRows || emailErr == orm.ErrMissPK {
    // user model with the username set
    user = models.User{Username: ctx.Params("identifier")}

    // find the user by username
    usernameErr := models.DB.Read(&user, "username")
    // check for errors
    if usernameErr == orm.ErrNoRows || usernameErr == orm.ErrMissPK {
      f.Error("Invalid credentials", false)
      // redirect the user
      ctx.Redirect(ctx.URLFor("auth.login"))
    }
  } else {
    // set the title
    f.Success("Successful login", false)
    // render the view
    ctx.Redirect(ctx.URLFor("wiki.list"))
  }
}
