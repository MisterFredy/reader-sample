package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reader/controller"
)

type Network struct {
	Port   string
	Router *mux.Router
}

func (n *Network) Route() {
	//s := *mux.Router

	// Login Route
	n.Router.HandleFunc("/login", SetMiddlewareJSON(controller.Login)).Methods("POST")

	//Users routes
	n.Router.HandleFunc("/users", SetMiddlewareJSON(controller.CreateUser)).Methods("POST")
	n.Router.HandleFunc("/users", SetMiddlewareJSON(controller.GetUsers)).Methods("GET")
	n.Router.HandleFunc("/users/{id}", SetMiddlewareJSON(controller.GetUser)).Methods("GET")
	n.Router.HandleFunc("/users/{id}", SetMiddlewareJSON(SetMiddlewareAuthentication(controller.UpdateUser))).Methods("PUT")
	n.Router.HandleFunc("/users/{id}", SetMiddlewareAuthentication(controller.DeleteUser)).Methods("DELETE")

	//Posts routes
	n.Router.HandleFunc("/posts", SetMiddlewareJSON(controller.CreatePost)).Methods("POST")
	n.Router.HandleFunc("/posts", SetMiddlewareJSON(controller.GetPosts)).Methods("GET")
	n.Router.HandleFunc("/posts/{id}", SetMiddlewareJSON(controller.GetPost)).Methods("GET")
	n.Router.HandleFunc("/posts/{id}", SetMiddlewareJSON(SetMiddlewareAuthentication(controller.UpdatePost))).Methods("PUT")
	n.Router.HandleFunc("/posts/{id}", SetMiddlewareAuthentication(controller.DeletePost)).Methods("DELETE")

	/*if err := http.ListenAndServe(n.Port, n.Router); err != nil {
		log.Fatal(err)
	}
	log.Fatal("running ")*/
	log.Fatal(http.ListenAndServe(n.Port, n.Router))
}
