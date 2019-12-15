package config

import (
	"github.com/gorilla/mux"
	"github.com/reader/controller"
)

type Routerlist struct {
	Router *mux.Router
}

func (s *Routerlist) Routinglist() {

	// Login Route
	s.Router.HandleFunc("/login", SetMiddlewareJSON(controller.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", SetMiddlewareJSON(controller.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", SetMiddlewareJSON(controller.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", SetMiddlewareJSON(controller.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", SetMiddlewareJSON(SetMiddlewareAuthentication(controller.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", SetMiddlewareAuthentication(controller.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", SetMiddlewareJSON(controller.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", SetMiddlewareJSON(controller.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", SetMiddlewareJSON(controller.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", SetMiddlewareJSON(SetMiddlewareAuthentication(controller.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", SetMiddlewareAuthentication(controller.DeletePost)).Methods("DELETE")

}
