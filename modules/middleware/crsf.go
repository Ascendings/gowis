package middleware

import (
	"github.com/go-macaron/csrf"
	"gopkg.in/macaron.v1"
)

func CsrfView(ctx *macaron.Context, x csrf.CSRF) {
	// append csrf value field to the view
	ctx.Data["csrf_token"] = "<input type=\"hidden\" name=\"_csrf\" value=\"" + x.GetToken() + "\">"
}
