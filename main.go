package main

import (
	"eroauz/models"
	"eroauz/server"
	"github.com/labstack/gommon/log"
)
func main() {
	models.Database("root:bakabie@/ero?charset=utf8&parseTime=True&loc=Local")
	e := server.NewRouter()
	log.Fatal(e.Start(":8000"))
}
