package app

import (
	"github.com/go-macaron/binding"

	"gogs.ballantine.tech/gballan1/gowis/controllers"
	"gogs.ballantine.tech/gballan1/gowis/modules/auth"
	"gogs.ballantine.tech/gballan1/gowis/modules/middleware"
	"gogs.ballantine.tech/gballan1/gowis/modules/wiki"

	"gopkg.in/macaron.v1"
)

// InitRouter - initializes the router and sets routes
func InitRouter(m macaron.Macaron) {
	bindIgnErr := binding.BindIgnErr

	// create new Wiki controller
	w := new(controllers.WikiController)
	// new Auth controller
	a := new(controllers.AuthController)

	// define routes
	m.Get("/", w.Home).Name("wiki.home")
	m.Get("/list", middleware.CheckUser, w.List).Name("wiki.list")
	m.Combo("/create").Get(w.Create).Post(bindIgnErr(wiki.PageForm{}), w.PostCreate).Name("wiki.create")
	m.Get("/view/:urlSlug", w.View).Name("wiki.view")
	m.Combo("/edit/:urlSlug").Get(w.Edit).Post(bindIgnErr(wiki.PageForm{}), w.PostEdit).Name("wiki.edit")

	m.Combo("/register").Get(a.Register).Post(bindIgnErr(auth.RegisterForm{}), a.PostRegister).Name("auth.register")
	m.Combo("/login").Get(a.Login).Post(bindIgnErr(auth.LoginForm{}), a.PostLogin).Name("auth.login")
}
