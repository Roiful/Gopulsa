package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Roiful/Gopulsa/App/models"
	"github.com/Roiful/Gopulsa/database/seeders"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
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
	fmt.Println("Welcome to " + AppConfig.AppName)

	server.initializeDB(DBConfig)
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) initializeDB(DBConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBConfig.DBUser, DBConfig.DBPassword, DBConfig.DBHost, DBConfig.DBPort, DBConfig.DBName)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Yah gagal konek")
	}
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully.")
}

func (server *Server) InitCommands(config AppConfig, DBConfig DBConfig) {
	server.initializeDB(DBConfig)
	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
