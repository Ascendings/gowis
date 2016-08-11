package routers

import (
	"gopkg.in/macaron.v1"
)

// WikiHome - home page
func WikiHome(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}
