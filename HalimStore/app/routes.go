package app

import "github.com/gieart87/halimstore/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc(path: "/", controllers.Home).Methods(methods: "GET")
}