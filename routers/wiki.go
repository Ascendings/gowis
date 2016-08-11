package routers

import (
  "gopkg.in/macaron.v1"
)

func Home(ctx *macaron.Context) string {
  return "the request path is: " + ctx.Req.RequestURI
}
