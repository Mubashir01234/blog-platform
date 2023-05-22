package routes

import (
	"blog/controller"

	"github.com/gorilla/mux"
)

type Routes struct {
	controller.ControllerImpl
}

func NewRoutes() *Routes {
	cont := controller.NewController()
	return &Routes{
		cont,
	}
}

func NewRoutesImpl(router *mux.Router) *mux.Router {
	controller := NewRoutes()

	api := router.PathPrefix("/api").Subrouter()

	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", controller.Register).Methods("POST")

	return router
}
