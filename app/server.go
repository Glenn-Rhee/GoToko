package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct{
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Welcome to gotoko")
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening to port %s \n", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func Run(){
	var server = Server{};
	server.Initialize()
	server.Run(":8000")
}