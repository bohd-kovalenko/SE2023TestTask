package main

import "net/http"

const BasePath = "/api"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
type Router struct {
	Routes []Route
}

func (r *Router) addNewRoute(requestType, path string, handler http.HandlerFunc) {
	route := Route{
		requestType,
		BasePath + path,
		handler,
	}
	r.Routes = append(r.Routes, route)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, route := range r.Routes {
		if route.Method == request.Method && route.Path == request.URL.Path {
			route.Handler(writer, request)
			return
		}
	}
	http.NotFound(writer, request)
}
