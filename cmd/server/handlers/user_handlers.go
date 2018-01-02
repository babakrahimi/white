package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/app"
)

func GetUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	a, err := app.New()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}
	user, err := a.GetUser(ps.ByName("username"))
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}
	writeJSONValue(w, user, http.StatusOK)
}

func GetUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	a, err := app.New()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	users, err := a.GetUsers()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}
	writeJSONValue(w, users, http.StatusOK)
}

func PostUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := r.PostFormValue("username")
	p := r.PostFormValue("password")

	a, err := app.New()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	if err = a.CreateUser(u, p); err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
