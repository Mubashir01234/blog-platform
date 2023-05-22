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
	profile.HandleFunc("", controller.UpdateProfile).Methods("PUT")
	profile.HandleFunc("", controller.GetProfile).Methods("Get")
	profile.HandleFunc("", controller.DeleteProfile).Methods("Delete")

	blog := api.PathPrefix("/blog").Subrouter()
	blog.Use(middleware.IsAuthorized)
	blog.HandleFunc("", controller.CreateBlog).Methods("POST")
	blog.HandleFunc("/{blog_id}", controller.UpdateBlog).Methods("PUT")

	api.HandleFunc("/blog/{blog_id}", controller.GetBlogById).Methods("GET")

	return router
}
