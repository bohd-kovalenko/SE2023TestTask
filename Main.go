package main

import (
	"SE2023/controllers"
	"fmt"
	"net/http"
)

func main() {
	router := routerInit()
	server := serverInit(router)
	serverRun(server)
}

func routerInit() Router {
	r := Router{}
	r.addNewRoute("GET", "/rate", controllers.RateHandler)
	return r
}

func serverInit(router Router) *http.Server {
	server := &http.Server{
		Addr:    ":8082",
		Handler: &router,
	}
	return server
}

func serverRun(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error when tried to start server")
		return
	}
}
