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

	mongodb, err := infrastructure.NewMongodbHandler(c.DB.URL, c.DB.DBName)
	if err != nil {
		log.Panic(err)
	}

	crypto := interfaces.NewCryptoHandler(c.JWTSecretKey)
	es := &infrastructure.EmailServer{
		ServerAddress: c.Email.ServerAddress,
		Username:      c.Email.Username,
		Password:      c.Email.Password,
	}
	email := interfaces.NewEmailHandler(es, c.Email.SignUpURL)

	ops := infrastructure.Operators{
		Invitation: GetInvitationOperator(mongodb, crypto, email),
	}

	ws := infrastructure.NewWebServer(ops, c.AllowedOrigins)
	log.Fatal(http.ListenAndServe(":"+c.Port, ws))
}

func GetInvitationOperator(mh interfaces.DBHandler, ch user.CryptoHandler, eh user.EmailHandler) user.InvitationOperator {
	invRepo := &interfaces.InvitationRepo{
		DB:          mh,
		StorageName: "invitations",
	}

	a := &user.InvitationAgent{
		Repository:    invRepo,
		CryptoHandler: ch,
		EmailHandler:  eh,
	}

	return a
}
