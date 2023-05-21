package server

import (
	"blog/api"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerImpl interface {
	HelloWorld(rw http.ResponseWriter, r *http.Request)
}

type Server struct {
	api api.BlogAPI
}

func NewServer() *Server {
	api := api.NewBlogAPIImpl()
	return &Server{
		api: api,
	}
}

func NewServerImpl(router *mux.Router) *mux.Router {
	server := NewServer()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/hello", server.HelloWorld).Methods("GET")
	return router
}

var _ ServerImpl = &Server{}
