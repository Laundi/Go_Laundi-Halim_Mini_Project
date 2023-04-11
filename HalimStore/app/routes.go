package app

func (server *Server) InitializeRoutes() {
server.Router.HundleFunc(path: "/", controllers.Home).Methods(methods: "GET")
}