package App

import (
	"flag"
	"log"
	"os"

	"github.com/Roiful/Gopulsa/App/controllers"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
func Run() {
	var server = controllers.Server{}
	var AppConfig = controllers.AppConfig{}
	var DBConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	AppConfig.AppName = GetEnv("APP_NAME", "GOpulsa")
	AppConfig.AppEnv = GetEnv("APP_ENV", "development")
	AppConfig.AppPort = GetEnv("APP_PORT", "9000")

	DBConfig.DBHost = GetEnv("DB_HOST", "127.0.0.1")
	DBConfig.DBUser = GetEnv("DB_USER", "")
	DBConfig.DBPassword = GetEnv("DB_PASSWORD", "")
	DBConfig.DBName = GetEnv("DB_NAME", "DBgopulsa")
	DBConfig.DBPort = GetEnv("DB_PORT", "3306")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommands(AppConfig, DBConfig)
	} else {
		server.Initialize(AppConfig, DBConfig)
		server.Run(":" + AppConfig.AppPort)
	}
}
