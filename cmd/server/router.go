package server

import (
	"net/http"
	"bytes"
	"encoding/json"
)

type Router struct {
	defaultHandler http.Handler
}

func GetRouter() *Router {
	return &Router{
		defaultHandler: getHandler(),
	}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.defaultHandler.ServeHTTP(w, r)
}

func toJson(w http.ResponseWriter, data interface{}) {
	bs := bytes.Buffer{}
	e := json.NewEncoder(&bs)
	e.Encode(data)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	bs.WriteTo(w)
}
