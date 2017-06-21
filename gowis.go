package main

import (
  "fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"

	macaron "gopkg.in/macaron.v1"

	"github.com/Ascendings/gowis/app"
	"github.com/Ascendings/gowis/app/models"
	"github.com/Ascendings/gowis/modules/middleware"
	"github.com/Ascendings/gowis/modules/settings"
	"github.com/Ascendings/gowis/modules/template"
)

func main() {
	// initialize macaron router
	m := macaron.Classic()

	// create DB connection
	models.InitDB()

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

	// integrate macaron's session module
  m.Use(session.Sessioner(session.Options{
    Provider:       "file",
    ProviderConfig: "data/sessions",
  }))

	// integrate macaron's caching module
	m.Use(cache.Cacher())

	// integrate CSRF protection stuff
	m.Use(csrf.Csrfer())

	// integrate our middleware into the application
	m.Use(middleware.CheckUser)
	m.Use(middleware.CsrfView)

	// initialize the router with routes
	app.InitRouter(*m)

  // get the values for the address and port to listen on
  hostname := settings.Cfg.Section("server").Key("address").String()
  port := settings.Cfg.Section("server").Key("port").String()
  socketAddress := strings.Join([]string{hostname, port}, ":")

	// let the user know we're running!
	log.Println(fmt.Sprintf("Starting the Gowis Wiki server on %s", socketAddress))
	if err := http.ListenAndServe(socketAddress, m);err != nil {
    log.Fatalln("Gowis Wiki has gone down with issue: ", err)
  }
}
