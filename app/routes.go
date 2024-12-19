package app

import (
	"github.com/Glenn-Rhee/gotoko/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/profile", controllers.Profile).Methods("GET")
}