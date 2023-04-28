package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB 		*gorm.DB
	Router 	*mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv 	string
	AppPort	string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println(a...: "Welcome to" + appConfig.AppName)

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println(format...: "Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(format: "Error on loading .env file")
	}

	appConfig.AppName = getEnv( key: "APP_NAME", fallback: "HalimStore")
	appConfig.AppEnv = getEnv( key: "APP_ENV", fallback: "development")
	appConfig.AppPort = getEnv( key: "APP_PORT", fallback: "9000")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
