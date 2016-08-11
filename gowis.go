package main

import (
	"log"
	"net/http"
	"strings"

	"gopkg.in/ini.v1"

	"gogs.ballantine.tech/gballan1/gowis/app"
)

func main() {
	// initialize router
	mac := app.InitRouter()

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
		cfg.Section("server").Key("port").String()}, ":"), mac))
}
