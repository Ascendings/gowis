package app

import (
  "gopkg.in/macaron.v1"

  . "gogs.ballantine.tech/gballan1/gowis/routers"
)

func InitRouter() *macaron.Macaron {
  // craete new router
  m := macaron.Classic()

  // define routes
  m.Get("/", Home)

  // return the compeleted router
  return m
}
