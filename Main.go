package main

import (
	"SE2023/controllers"
	"SE2023/repositories"
	"fmt"
	"net/http"
)

const ApplicationPort = ":8082"

func main() {
	err := repositories.InitRepository()
	if err != nil {
		fmt.Println("Error, when initializing repository ", err)
		return
	}
	router := routerInit()
	server := serverInit(router)
	serverRun(server)
}

func routerInit() Router {
	r := Router{}
	r.addNewRoute("GET", "/rate", controllers.RateHandler)
	r.addNewRoute("POST", "/subscribe", controllers.SubscriptionNewMailHandler)
	r.addNewRoute("POST", "/subscription", controllers.MailSendingHandler)
	return r
}

func serverInit(router Router) *http.Server {
	server := &http.Server{
		Addr:    ApplicationPort,
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
