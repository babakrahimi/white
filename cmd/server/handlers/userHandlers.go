package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/app"
	"encoding/json"
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

func PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	d := json.NewDecoder(r.Body)
	u := app.User{}
	err := d.Decode(&u)
	if err != nil {
		toServerError(w, err)
		return
	}
	defer r.Body.Close()

	a, err := app.New()
	if err != nil {
		toServerError(w, err)
		return
	}
	nu, err := a.AddUsers(&u)
	if err != nil {
		toServerError(w, err)
		return
	}
	toCreated(w, nu)
}