package App

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}
type AppConfig struct {
	Appname string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(AppConfig AppConfig) {
	fmt.Println("Selamat Datang di GoPulsa" + AppConfig.Appname)

	server.Router = mux.NewRouter()
	server.InitializeRouter()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var AppConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error on loading .env file")
	}

	AppConfig.Appname = Getenv("APP_NAME", "GOpulsa")
	AppConfig.AppEnv = Getenv("APP_ENV", "development")
	AppConfig.AppPort = Getenv("APP_PORT", "9000")

	server.Initialize(AppConfig)
	server.Run(":" + AppConfig.AppPort)
}
