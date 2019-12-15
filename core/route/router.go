package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reader/config"
)

func Run(port string) {
	var routelist = config.Routerlist{}

	routelist.Router = mux.NewRouter()
	routelist.Routinglist()

	if err := http.ListenAndServe(port, routelist.Router); err != nil {
		log.Fatal(err)
	}
	log.Fatal("running ")
}
