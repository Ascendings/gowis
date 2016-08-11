package main

import (
	"log"
	"net/http"

	"gogs.ballantine.tech/gballan1/gowis/app"
)

func main() {
	// initialize router
	m := app.InitRouter()

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}
