package infrastructure

import (
	"fmt"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
)

type (
	EmailServer struct {
		ServerAddress, Username, Password string
	}
)

func (es *EmailServer) Send(to, subject, content string) error {
	sender := mail.Address{Name: "تیم نرم‌افزار ماهان", Address: es.Username}
	recipient := mail.Address{Name: "", Address: to}

	headers := make(map[string]string)
	headers["From"] = sender.String()
	headers["Subject"] = mime.QEncoding.Encode("utf-8", subject)
	headers["To"] = recipient.String()
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=utf-8"

	body := ""
	for k, v := range headers {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	body += content

	host, _, _ := net.SplitHostPort(es.ServerAddress)
	auth := smtp.PlainAuth(
		"",
		es.Username,
		es.Password,
		host,
	)
	if err := smtp.SendMail(
		es.ServerAddress,
		auth,
		sender.Address,
		[]string{recipient.Address},
		[]byte(body),
	); err != nil {
		return err
	}
	return nil
}
