package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"

	macaron "gopkg.in/macaron.v1"

	"gogs.ballantine.tech/gballan1/gowis/app"
	"gogs.ballantine.tech/gballan1/gowis/app/models"
	"gogs.ballantine.tech/gballan1/gowis/app/modules/template"
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

	// get the template function map
	funcMap := template.NewFuncMap(m)

	m.Use(macaron.Renderer(macaron.RenderOptions{
		// Directory to load templates. Default is "templates".
		Directory: "resources/views",
		// Funcs is a slice of FuncMaps to apply to the template upon compilation. Default is [].
		Funcs: funcMap,
		// Outputs human readable JSON. Default is false.
		IndentJSON: true,
		// Outputs human readable XML. Default is false.
		IndentXML: true,
		// Prefixes the JSON output with the given bytes. Default is no prefix.
		PrefixJSON: []byte("macaron"),
		// Prefixes the XML output with the given bytes. Default is no prefix.
		PrefixXML: []byte("macaron"),
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
