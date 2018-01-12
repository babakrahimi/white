package infrastructure

import (
	"net/http"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/business/user"
	"github.com/rs/cors"
)

type (
	Operators struct {
		Invitation user.InvitationOperator
	}

	WebServerHandler struct {
		Router *httprouter.Router
	}
)

func (ws *WebServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.Router.ServeHTTP(w, r)
}

func NewWebServer(ops Operators, allowedOrigins []string) http.Handler {
	r := httprouter.New()

	r.POST("/invite", GetInviteUserHandler(ops))
	r.POST("/verify", GetVerifyInvitationHandler(ops))

	ws := &WebServerHandler{
		Router: r,
	}
	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(ws)
}

func GetVerifyInvitationHandler(ops Operators) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		token := r.PostFormValue("token")
		inv, err := ops.Invitation.VerifyInvitation(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if err := WriteJSON(w, inv); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetInviteUserHandler(ops Operators) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		email := r.PostFormValue("email")
		if err := ops.Invitation.InviteUser(email); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return nil
}
