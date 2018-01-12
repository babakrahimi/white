package infrastructure

import (
	"net/http"

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

	ws := &WebServerHandler{
		Router: r,
	}
	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(ws)
}

func GetInviteUserHandler(ops Operators) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		email := r.PostFormValue("email")
		err := ops.Invitation.InviteUser(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
