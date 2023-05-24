package main

import (
	"SE2023/controllers"
	"net/http"
)

func main() {
	router := routerInit()
	server := &http.Server{
		Addr:    ":8082",
		Handler: &router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func routerInit() Router {
	r := Router{}
	r.addNewRoute("GET", "/api/rate", controllers.RateHandler)
	return r
}
