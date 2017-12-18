package server

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/app"
)

func getHandler() http.Handler {
	r := httprouter.New()
	r.GET("/api/restaurant", getRestaurant)
	return r
}

func getRestaurant(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	toJson(w, app.GetRestaurants())
}

