package app

import "github.com/Glenn-Rhee/gotoko/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/profile", controllers.Profile).Methods("GET")
}