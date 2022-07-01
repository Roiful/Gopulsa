package App

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
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
	DBDriver   string
}

func (server *Server) Initialize(AppConfig AppConfig, DBConfig DBConfig) {
	fmt.Println("Selamat Datang di GoPulsa" + AppConfig.Appname)
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBConfig.DBUser, DBConfig.DBPassword, DBConfig.DBHost, DBConfig.DBPort, DBConfig.DBName)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Yah gagal konek")
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

	DBConfig.DBHost = Getenv("DB_HOST", "127.0.0.1")
	DBConfig.DBUser = Getenv("DB_USER", "")
	DBConfig.DBPassword = Getenv("DB_PASSWORD", "")
	DBConfig.DBName = Getenv("DB_NAME", "DBgopulsa")
	DBConfig.DBPort = Getenv("DB_PORT", "3306")

	server.Initialize(AppConfig, DBConfig)
	server.Run(":" + AppConfig.AppPort)
}
