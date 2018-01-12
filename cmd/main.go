package main

import (
	"log"
	"net/http"

	"github.com/megaminx/white/cmd/business/user"
	"github.com/megaminx/white/cmd/infrastructure"
	"github.com/megaminx/white/cmd/interfaces"
)

func main() {
	c, err := infrastructure.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	mh, err := infrastructure.NewMongodbHandler(c.DB.URL, c.DB.DBName)
	if err != nil {
		log.Panic(err)
	}

	repo := interfaces.InvitationRepo{
		DB: mh,
	}

	ch := interfaces.CryptoHandler{SecretKey: c.JWTSecretKey}

	eh := interfaces.EmailHandler{
		Provider: &infrastructure.EmailServer{
			ServerAddress: c.Email.ServerAddress,
			Username:      c.Email.Username,
			Password:      c.Email.Password,
		},
		RegistrationRedirectURL: c.Email.RegistrationRedirectURL,
	}

	ia := user.InvitationAgent{
		Repository:    &repo,
		CryptoHandler: &ch,
		EmailHandler:  &eh,
	}

	agents := infrastructure.Agents{
		Invitation: &ia,
	}

	ws := infrastructure.NewWebService(agents, c.AllowedOrigins)
	log.Fatal(http.ListenAndServe(":"+c.Port, ws))
}
