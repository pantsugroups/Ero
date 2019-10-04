package main

import (
	"eroauz/conf"
	"eroauz/models"
	"eroauz/server"
	"github.com/labstack/gommon/log"
)

func main() {
	models.Database(conf.ParseDataBaseConfigure())
	e := server.NewRouter()
	log.Fatal(e.Start(conf.WebPort))
}
