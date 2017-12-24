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
	toOk(w, user)
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
	toOk(w, users)
}

func PostUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := r.PostFormValue("username")
	p := r.PostFormValue("password")
	e := r.PostFormValue("email")

	a, err := app.New()
	if err != nil {
		toServerError(w, err)
		return
	}

	if err = a.CreateUser(u, p, e); err != nil {
		toServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
