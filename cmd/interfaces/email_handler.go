package interfaces

import (
	"bytes"
	"fmt"
	"html/template"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
)

type EmailHandler struct {
	MailServer MailServer
	Sender     string
	BaseUrl    string
}
type MailServer struct {
	ServerAddress string
	Username      string
	Password      string
}

func (mp *EmailHandler) SendInvitationEmail(to, token string) error {
	fromAddr := mail.Address{Name: "تیم نرم‌افزار ماهان", Address: mp.Sender}
	toAddr := mail.Address{Address: to}

	headers := make(map[string]string)
	headers["From"] = fromAddr.String()
	headers["Subject"] = mime.QEncoding.Encode("utf-8", "عضویت در سامانه")
	headers["To"] = toAddr.String()
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=utf-8"
	body := ""
	for k, v := range headers {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	content, err := getInvitationEmailContent(mp.BaseUrl, token)
	if err != nil {
		return err
	}
	body += content

	host, _, _ := net.SplitHostPort(mp.MailServer.ServerAddress)
	auth := smtp.PlainAuth(
		"",
		mp.MailServer.Username,
		mp.MailServer.Password,
		host,
	)
	if err := smtp.SendMail(
		mp.MailServer.ServerAddress,
		auth,
		fromAddr.Address,
		[]string{toAddr.Address},
		[]byte(body),
	); err != nil {
		return err
	}

	return nil
}

func getInvitationEmailContent(url, token string) (string, error) {
	msg := `
<!doctype html>
<html lang="fa">
<head>
	<meta charset="UTF-8">
</head>
<body>
	<div dir="rtl">
		<p>سلام همکار عزیز،</p>
		<p>لطفا برای عضویت و تکمیل مشخصات فردی خود در سامانه تیم نرم‌افزار ماهان، بر روی لینک زیر کلیک کنید.</p>
		<p dir="ltr">
			<a href="{{.URL}}?t={{.Token}}">{{.URL}}</a>
		</p>
		<br>
		<p>
		باتشکر
		<br>
		سامانه تیم نرم‌افزار ماهان
		</p>
	</div>  
</body>
</html>
`
	t := template.Must(template.New("email").Parse(msg))
	b := bytes.NewBufferString("")
	err := t.Execute(b, struct {
		URL, Token string
	}{
		URL:   url,
		Token: token,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
