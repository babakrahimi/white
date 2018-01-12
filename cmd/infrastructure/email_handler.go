package infrastructure

import "github.com/megaminx/white/cmd/interfaces"

func NewEmailHandler(c *EmailConfig) *interfaces.EmailHandler {
	ms := interfaces.MailServer{
		ServerAddress: c.ServerAddress,
		Username:      c.RegistrationAddress,
		Password:      c.RegistrationPassword,
	}
	eh := &interfaces.EmailHandler{
		MailServer:              ms,
		Sender:                  c.RegistrationAddress,
		RegistrationRedirectURL: c.RegistrationRedirectURL,
	}
	return eh
}
