package controllers

import macaron "gopkg.in/macaron.v1"

// Controller - base class for controllers
type Controller struct{}

// Render - renders a view
func (c *Controller) Render(ctx *macaron.Context, view string) {
	// do the view
	ctx.HTML(200, view)
}
