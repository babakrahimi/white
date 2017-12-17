package app

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/api/restaurant":
		writeToJson(w, GetRestaurants())
	}
}

func writeToJson(w http.ResponseWriter, data interface{}) {
	bs := bytes.Buffer{}
	e := json.NewEncoder(&bs)
	e.Encode(data)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	bs.WriteTo(w)
}
