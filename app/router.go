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

	// group stuff together so that global middleware can run
	m.Group("", func() {

		// regular routes
		m.Get("/", w.Home).Name("wiki.home")
		m.Get("/list", w.List).Name("wiki.list")
		m.Get("/view/:urlSlug", w.View).Name("wiki.view")

		// authenticated users only routes
		m.Group("", func() {
			m.Combo("/create").Get(w.Create).Post(bindIgnErr(wiki.PageForm{}), w.PostCreate).Name("wiki.create")
			m.Combo("/edit/:urlSlug").Get(w.Edit).Post(bindIgnErr(wiki.PageForm{}), w.PostEdit).Name("wiki.edit")
			m.Get("/logout", a.Logout).Name("auth.logout")
		}, middleware.Auth)

		// guest only routes
		m.Group("", func() {
			m.Combo("/register").Get(a.Register).Post(bindIgnErr(auth.RegisterForm{}), a.PostRegister).Name("auth.register")
			m.Combo("/login").Get(a.Login).Post(bindIgnErr(auth.LoginForm{}), a.PostLogin).Name("auth.login")
		}, middleware.Guest)

	}, middleware.CheckUser)

}
