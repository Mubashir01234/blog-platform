package controller

import (
	"blog/db"
	"net/http"

	"github.com/gorilla/mux"
)

type ControllerImpl interface {
	Register(rw http.ResponseWriter, r *http.Request)
}

type Controller struct {
	db db.BlogDB
}

func NewController() *Controller {
	db := db.ConnectDB()
	return &Controller{
		db: db,
	}
}

func NewControllerImpl(router *mux.Router) *mux.Router {
	controller := NewController()

	api := router.PathPrefix("/api").Subrouter()

	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", controller.Register).Methods("POST")

	return router
}

var _ ControllerImpl = &Controller{}
