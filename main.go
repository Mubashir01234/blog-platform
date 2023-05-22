package main

import (
	"log"
	"net/http"

	"blog/config"
	"blog/routes"

	"github.com/fatih/color"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	color.Cyan("🌏 Server running on localhost:" + config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, routes.NewRoutesImpl()))
}
