package controllers

import (
	"github.com/go-macaron/csrf"

	macaron "gopkg.in/macaron.v1"
)

// Controller - base class for controllers
type Controller struct{}

// CreateCsrfField - generates the HTML for the csrf hidden input field
func (c *Controller) CreateCsrfField(x csrf.CSRF) string {
	return "<input type=\"hidden\" name=\"_csrf\" value=\"" + x.GetToken() + "\">"
}

// Render - renders a view
func (c *Controller) Render(ctx *macaron.Context, view string) {
	// do the view
	ctx.HTML(200, view)
}