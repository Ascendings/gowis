package app

import (
	"github.com/go-macaron/binding"
	"github.com/go-macaron/csrf"

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
	// new Users controller
	u := new(controllers.UserController)

	// group stuff together so that global middleware can run
	m.Group("", func() {

		// regular routes
		m.Get("/", w.Home).Name("wiki.home")

		m.Get("/wiki/list", w.List).Name("wiki.list")
		m.Get("/wiki/view/:urlSlug", w.View).Name("wiki.view")

		m.Get("/users", u.List).Name("users.list")
		m.Get("/users/:userID", u.View).Name("users.view")

		// authenticated users only routes
		m.Group("", func() {
			m.Combo("/wiki/create").Get(w.Create).Post(bindIgnErr(wiki.PageForm{}), csrf.Validate, w.PostCreate).Name("wiki.create")
			m.Combo("/wiki/edit/:urlSlug").Get(w.Edit).Post(bindIgnErr(wiki.PageForm{}), csrf.Validate, w.PostEdit).Name("wiki.edit")
			m.Get("/auth/logout", a.Logout).Name("auth.logout")
		}, middleware.Auth)

		// guest only routes
		m.Group("", func() {
			m.Combo("/auth/register").Get(a.Register).Post(bindIgnErr(auth.RegisterForm{}), csrf.Validate, a.PostRegister).Name("auth.register")
			m.Combo("/auth/login").Get(a.Login).Post(bindIgnErr(auth.LoginForm{}), csrf.Validate, a.PostLogin).Name("auth.login")
		}, middleware.Guest)

	}, middleware.CheckUser)

}
