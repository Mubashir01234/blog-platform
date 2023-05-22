package routes

import (
	"blog/controller" // Importing the controller package
	"blog/middleware" // Importing the middleware package

	"github.com/gorilla/mux" // Importing the gorilla/mux package for router functionality
)

type Routes struct {
	controller.ControllerImpl // Embedding the ControllerImpl struct from the controller package
}

func NewRoutes() *Routes {
	cont := controller.NewController() // Creating a new instance of the controller
	return &Routes{
		cont,
	}
}

func NewRoutesImpl() *mux.Router {
	router := mux.NewRouter() // Creating a new instance of the router
	controller := NewRoutes() // Creating a new instance of the routes

	api := router.PathPrefix("/api").Subrouter() // Creating a subrouter for the "/api" path prefix

	user := api.PathPrefix("/user").Subrouter()                       // Creating a subrouter for the "/api/user" path prefix
	user.HandleFunc("/register", controller.Register).Methods("POST") // Registering the Register function of the controller for the "/api/user/register" path, with the HTTP method "POST"
	user.HandleFunc("/login", controller.Login).Methods("POST")       // Registering the Login function of the controller for the "/api/user/login" path, with the HTTP method "POST"

	profile := api.PathPrefix("/profile").Subrouter()                  // Creating a subrouter for the "/api/profile" path prefix
	profile.Use(middleware.IsAuthorized)                               // Applying the IsAuthorized middleware to all the routes under the "/api/profile" path
	profile.HandleFunc("", controller.UpdateProfile).Methods("PUT")    // Registering the UpdateProfile function of the controller for the "/api/profile" path, with the HTTP method "PUT"
	profile.HandleFunc("", controller.GetProfile).Methods("Get")       // Registering the GetProfile function of the controller for the "/api/profile" path, with the HTTP method "GET"
	profile.HandleFunc("", controller.DeleteProfile).Methods("DELETE") // Registering the DeleteProfile function of the controller for the "/api/profile" path, with the HTTP method "DELETE"

	blog := api.PathPrefix("/blog").Subrouter()                            // Creating a subrouter for the "/api/blog" path prefix
	blog.Use(middleware.IsAuthorized)                                      // Applying the IsAuthorized middleware to all the routes under the "/api/blog" path
	blog.HandleFunc("", controller.CreateBlog).Methods("POST")             // Registering the CreateBlog function of the controller for the "/api/blog" path, with the HTTP method "POST"
	blog.HandleFunc("/{blog_id}", controller.UpdateBlog).Methods("PUT")    // Registering the UpdateBlog function of the controller for the "/api/blog/{blog_id}" path, with the HTTP method "PUT"
	blog.HandleFunc("/{blog_id}", controller.DeleteBlog).Methods("DELETE") // Registering the DeleteBlog function of the controller for the "/api/blog/{blog_id}" path, with the HTTP method "DELETE"

	api.HandleFunc("/blog/{blog_id}", controller.GetBlogById).Methods("GET")                 // Registering the GetBlogById function of the controller for the "/api/blog/{blog_id}" path, with the HTTP method "GET"
	api.HandleFunc("/{username}/blogs", controller.GetUserAllBlogsByUsername).Methods("GET") // Registering the GetUserAllBlogsByUsername function of the controller for the "/api/{username}/blogs" path, with the HTTP method "GET"

	return router // Returning the router
}
