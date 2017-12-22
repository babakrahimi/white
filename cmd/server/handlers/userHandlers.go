package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/app"
)

func GetUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	a, err := app.New()
	if err != nil {
		toServerError(w, err)
		return
	}
	user, err := a.GetUser(ps.ByName("username"))
	if err != nil {
		toServerError(w, err)
		return
	}
	toJson(w, user)
}

func GetUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	a, err := app.New()
	if err != nil {
		toServerError(w, err)
		return
	}

	users, err := a.GetUsers()
	if err != nil {
		toServerError(w, err)
		return
	}
	toJson(w, users)
}
