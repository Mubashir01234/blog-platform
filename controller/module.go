package controller

import (
	"blog/db"
	"net/http"
)

type ControllerImpl interface {
	Register(rw http.ResponseWriter, r *http.Request)
	Login(rw http.ResponseWriter, r *http.Request)
	UpdateProfile(rw http.ResponseWriter, r *http.Request)
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

var _ ControllerImpl = &Controller{}
