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

type DBConfig struct {
	DBHost		string
	DBUser		string
	DBPassword	string
	DBname		string
	DBPort		string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println(a...: "Welcome to" + appConfig.AppName)

	var err error
	dsn := fmt.Sprintf( format: "host:%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic( v: "Failed on connecting to the database server")
	}

	if err != nil {
		panic(v: "Failed on connection to the server")
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port #{addr}")
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
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(format: "Error on loading .env file")
	}

	appConfig.AppName = getEnv( key: "APP_NAME", fallback: "HalimStore")
	appConfig.AppEnv = getEnv( key: "APP_ENV", fallback: "development")
	appConfig.AppPort = getEnv( key: "APP_PORT", fallback: "9000")

	dbConfig.DBHost = getEnv( key: "DB_HOST", fallback: "localhost")
	dbConfig.DBUser = getEnv( key: "DB_USER", fallback: "user")
	dbConfig.DBPassword = getEnv( key: "DB_PASSWORD", fallback: "password")
	dbConfig.DBName = getEnv( key: "DB_NAME", fallback: "dbname")
	dbConfig.DBPort = getEnv( key: "DB_PORT", fallback: 5432)

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
