package main

import (
	"log"
	"net/http"
	"strings"

	"gopkg.in/ini.v1"
	macaron "gopkg.in/macaron.v1"

	"gogs.ballantine.tech/gballan1/gowis/app"
)

func main() {
	// initialize macaron router
	m := macaron.Classic()

	// setup the template engine
	app.InitTemplates(*m)

	// initialize the router with routes
	app.InitRouter(*m)

	// load the config file
	cfg, err := ini.InsensitiveLoad("./app/app.ini")

	// check for errors while loading the configuration
	if err != nil {
		panic(err)
	}

	// let the user know we're running!
	log.Println("Server is running...")
	log.Println(http.ListenAndServe(strings.Join([]string{
		cfg.Section("server").Key("address").String(),
		cfg.Section("server").Key("port").String()}, ":"), m))
}
