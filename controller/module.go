package controller

import (
	"blog/db" // Importing the db package
	"net/http"
)

// This is interface of the controller that shows all the functions of the controller
type ControllerImpl interface {
	Register(rw http.ResponseWriter, r *http.Request)
	Login(rw http.ResponseWriter, r *http.Request)
	UpdateProfile(rw http.ResponseWriter, r *http.Request)
	GetProfile(rw http.ResponseWriter, r *http.Request)
	DeleteProfile(rw http.ResponseWriter, r *http.Request)
	CreateBlog(rw http.ResponseWriter, r *http.Request)
	UpdateBlog(rw http.ResponseWriter, r *http.Request)
	GetBlogById(rw http.ResponseWriter, r *http.Request)
	DeleteBlog(rw http.ResponseWriter, r *http.Request)
	GetUserAllBlogsByUsername(rw http.ResponseWriter, r *http.Request)
}

type Controller struct {
	db db.BlogDB // Struct field to hold a database connection
}

func NewController() *Controller {
	db := db.ConnectDB() // Establishing a new database connection
	return &Controller{
		db: db,
	}
}

var _ ControllerImpl = &Controller{} // Ensuring that the Controller struct implements the ControllerImpl interface
