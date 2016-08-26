package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/session"

	"gogs.ballantine.tech/gballan1/gowis/app"
	"gogs.ballantine.tech/gballan1/gowis/models"

	"gopkg.in/macaron.v1"
)

func main() {
	// initialize macaron router
	m := macaron.Classic()

	// load our configuration
	cfg := app.InitConfig()

	// integrate macaron's caching module
	m.Use(cache.Cacher())

	// integrate macaron's session module
	m.Use(session.Sessioner())

	// integrate CSRF protection stuff
	m.Use(csrf.Csrfer())

	// setup the Pongo2 template engine
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "views",
		Extensions: []string{".jinja", ".tmpl"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}))

	// initialize the router with routes
	app.InitRouter(*m)

	// create DB connection
	models.InitDB()

	// let the user know we're running!
	log.Println("Server is running...")
	log.Println(http.ListenAndServe(strings.Join([]string{
		cfg.Section("server").Key("address").String(),
		cfg.Section("server").Key("port").String()}, ":"), m))
}
