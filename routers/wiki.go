package routers

import (
	"gopkg.in/macaron.v1"
)

// WikiHome - home page
func WikiHome(ctx *macaron.Context) {
	ctx.Data["title"] = "Gowis"
	ctx.HTML(200, "wiki/home")
}
