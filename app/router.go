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
	bs := bytes.Buffer{}
	e := json.NewEncoder(&bs)

	switch r.RequestURI {
	case "/api/restaurant":
		e.Encode(GetRestaurants())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		bs.WriteTo(w)
	}

}
