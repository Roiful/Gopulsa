package App

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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
type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(AppConfig AppConfig) {
	fmt.Println("Selamat Datang di GoPulsa" + AppConfig.Appname)

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", "localhost", "Roiful", "password", "Roiful_Gopulsadb", "5432")
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal koneksi ke database")
	}

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
	var DBConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error on loading .env file")
	}

	AppConfig.Appname = Getenv("APP_NAME", "GOpulsa")
	AppConfig.AppEnv = Getenv("APP_ENV", "development")
	AppConfig.AppPort = Getenv("APP_PORT", "9000")

	DBConfig.DBHost = Getenv("DB_HOST", "localhost")
	DBConfig.DBUser = Getenv("DB_USER", "user")
	DBConfig.DBPassword = Getenv("DB_PASSWORD", "password")
	DBConfig.DBName = Getenv("DB_NAME", "dbname")
	DBConfig.DBPort = Getenv("DB_PORT", "5432")

	server.Initialize(AppConfig)
	server.Run(":" + AppConfig.AppPort)
}
