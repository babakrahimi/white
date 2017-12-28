package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/server/handlers"
	"github.com/rs/cors"
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
	hr := httprouter.New()

	hr.GET("/api/users", handlers.GetUsers)
	hr.GET("/api/user/:username", handlers.GetUser)
	hr.POST("/api/user", handlers.PostUser)

	h := getCORSHandler(hr)

	return h
}

func getCORSHandler(hr *httprouter.Router) http.Handler {
	h := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(hr)
	return h
}
