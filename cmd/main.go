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

	crypto := NewCryptoHandler(c.JWTSecretKey)
	email := NewEmailHandler(c.Email)

	ops := infrastructure.Operators{
		Invitation: GetInvitationOperator(mongodb, crypto, email),
	}

	ws := infrastructure.NewWebServer(ops, c.AllowedOrigins)
	log.Fatal(http.ListenAndServe(":"+c.Port, ws))
}

func NewCryptoHandler(secretKey string) interfaces.CryptoHandler {
	crypto := interfaces.CryptoHandler{SecretKey: secretKey}
	return crypto
}

func NewEmailHandler(ec *infrastructure.EmailConfig) interfaces.EmailHandler {
	eh := interfaces.EmailHandler{
		Provider: &infrastructure.EmailServer{
			ServerAddress: ec.ServerAddress,
			Username:      ec.Username,
			Password:      ec.Password,
		},
		RegistrationRedirectURL: ec.RegistrationRedirectURL,
	}
	return eh
}

func GetInvitationOperator(mh interfaces.DBHandler, ch interfaces.CryptoHandler, eh interfaces.EmailHandler) user.InvitationOperator {
	invRepo := interfaces.InvitationRepo{
		DB:          mh,
		StorageName: "invitations",
	}

	ia := user.InvitationAgent{
		Repository:    &invRepo,
		CryptoHandler: &ch,
		EmailHandler:  &eh,
	}

	return ia
}
