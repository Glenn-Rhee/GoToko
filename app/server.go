package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct{
	DB *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv 	string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to" + appConfig.AppName)
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening to port %s \n", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func getEnv(key string, fallBack string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallBack
}

func Run(){
	var server = Server{};
	var appConfig = AppConfig{}
	
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error on loading .env file: %v", err)
	}

	appConfig.AppName = getEnv("APP_NAME", "GoToko")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort =getEnv("APP_PORT", ":8081")

	server.Initialize(appConfig)
	server.Run(appConfig.AppPort)
}