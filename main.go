package main

import (
	"log"
	"net/http"

	"blog/config"
	"blog/controller"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()
	color.Cyan("🌏 Server running on localhost:" + config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, controller.NewControllerImpl(router)))
}
