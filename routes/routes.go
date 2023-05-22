package routes

import (
	"blog/controller"
	"blog/middleware"

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

func NewRoutesImpl() *mux.Router {
	router := mux.NewRouter()
	controller := NewRoutes()

	api := router.PathPrefix("/api").Subrouter()

	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", controller.Register).Methods("POST")
	user.HandleFunc("/login", controller.Login).Methods("POST")

	profile := api.PathPrefix("/profile").Subrouter()
	profile.Use(middleware.IsAuthorized)
	profile.HandleFunc("/update", controller.UpdateProfile).Methods("PUT")
	profile.HandleFunc("", controller.GetProfile).Methods("Get")
	profile.HandleFunc("", controller.DeleteProfile).Methods("Delete")

	return router
}
