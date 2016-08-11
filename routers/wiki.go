package routers

import (
	"gopkg.in/macaron.v1"
)

// Wiki router
type Wiki struct{}

// Home - home page
func (w *Wiki) Home(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}

// Test - testing stuff
func (w *Wiki) Test() string {
	return "Testing"
}
