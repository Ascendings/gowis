package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-macaron/pongo2"

	"gopkg.in/macaron.v1"

	"gogs.ballantine.tech/gballan1/gowis/app"
)

func main() {
	// initialize macaron router
	m := macaron.Classic()

	// load our configuration
	cfg := app.InitConfig()

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

	// load our DB stuff
	app.InitDB()

	// let the user know we're running!
	log.Println("Server is running...")
	log.Println(http.ListenAndServe(strings.Join([]string{
		cfg.Section("server").Key("address").String(),
		cfg.Section("server").Key("port").String()}, ":"), m))
}
