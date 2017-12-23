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
	allowCORS(w)
	router.handler.ServeHTTP(w, r)
}

func allowCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func getHandler() http.Handler {
	r := httprouter.New()

	r.GET("/api/users", handlers.GetUsers)
	r.GET("/api/user/:username", handlers.GetUser)
	r.POST("/api/user", handlers.PostUser)

	return r
}
