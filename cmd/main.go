package main

import (
	"log"
	"net/http"
	"os"

	"strings"

	"github.com/megaminx/white/cmd/business/user"
	"github.com/megaminx/white/cmd/infrastructure"
	"github.com/megaminx/white/cmd/interfaces"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panic("$Port is not set")
	}

	dbURL := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("JWT_SECRET")
	mh, err := infrastructure.NewMongodbHandler(dbURL, dbName)
	if err != nil {
		log.Panic(err)
	}

	repo := interfaces.InvitationRepo{
		DB: mh,
	}

	emailServer := os.Getenv("EMAIL_SERVER")
	emailAddress := os.Getenv("EMAIL_ADDRESS_REG")
	emailPassword := os.Getenv("EMAIL_PASSWORD_REG")
	baseURL := os.Getenv("REGISTRATION_REDIRECT_URL")
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	ms := interfaces.MailServer{
		ServerAddress: emailServer,
		Username:      emailAddress,
		Password:      emailPassword,
	}
	eh := interfaces.EmailHandler{
		MailServer: ms,
		Sender:     emailAddress,
		BaseUrl:    baseURL,
	}
	ch := interfaces.CryptoHandler{SecretKey: secretKey}

	ia := user.InvitationAgent{
		Repository:    &repo,
		CryptoHandler: &ch,
		EmailHandler:  &eh,
	}

	agents := infrastructure.Agents{
		Invitation: &ia,
	}

	ws := infrastructure.NewWebService(agents, allowedOrigins)
	log.Fatal(http.ListenAndServe(":"+port, ws))
}
