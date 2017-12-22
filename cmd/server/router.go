package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/server/handlers"
)

type Router struct {
	handler http.Handler
}

func NewRouter() *Router {
	r := &Router{
		handler: getHandler(),
	}
	return r
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.handler.ServeHTTP(w, r)
}

func getHandler() http.Handler {
	r := httprouter.New()

	r.GET("/api/users", handlers.GetUsers)
	r.GET("/api/user/:username", handlers.GetUser)

	return r
}
