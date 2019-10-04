package main

import (
	"eroauz/conf"
	"eroauz/models"
	"eroauz/server"
	"github.com/labstack/gommon/log"
)

// @title Ero BackEnd API
// @version 1.0
// @description Ero Server 's hmp BackEnd
// @termsOfService https://api.ero.ink

// @contact.name API Support
// @contact.url https://9bie.org
// @contact.email blackguwc@163.com

// @license.name What the Fuck PL
// @license.url https://ero.ink

// @host api.ero.ink
// @BasePath /api/v1
func main() {
	models.Database(conf.ParseDataBaseConfigure())
	e := server.NewRouter()
	log.Fatal(e.Start(conf.WebPort))
}
