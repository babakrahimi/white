package infrastructure

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/megaminx/white/cmd/business/user"
	"github.com/rs/cors"
)

type (
	Agents struct {
		Invitation *user.InvitationAgent
	}

	WebServiceHandler struct {
		Router *httprouter.Router
	}
)

func (ws *WebServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.Router.ServeHTTP(w, r)
}

func NewWebService(agents Agents, allowedOrigins []string) http.Handler {
	r := httprouter.New()

	r.POST("/invite", GetInviteUserHandler(agents))

	ws := &WebServiceHandler{
		Router: r,
	}
	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(ws)
}

func GetInviteUserHandler(agents Agents) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		email := r.PostFormValue("email")
		err := agents.Invitation.InviteUser(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
