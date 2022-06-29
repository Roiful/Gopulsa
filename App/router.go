package App

import "github.com/Roiful/Gopulsa/App/controllers"

func (server *Server) InitializeRouter() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
