package controllers

import (
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
func (a *AuthController) PostLogin(ctx *macaron.Context, input wiki.PageForm, f *session.Flash) {

}
