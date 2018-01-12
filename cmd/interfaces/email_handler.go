package interfaces

import (
	"bytes"
	"html/template"
)

type (
	EmailServerProvider interface {
		Send(to, subject, content string) error
	}

	EmailHandler struct {
		Provider                EmailServerProvider
		RegistrationRedirectURL string
	}
)

func (eh *EmailHandler) SendInvitationEmail(to, token string) error {
	subject := "عضویت در سامانه"
	content, err := GetInvitationEmailContent(eh.RegistrationRedirectURL, token)
	if err != nil {
		return err
	}

	if err := eh.Provider.Send(to, subject, content); err != nil {
		return err
	}
	return nil
}

func GetInvitationEmailContent(url, token string) (string, error) {
	body := `
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
`
	t := template.Must(template.New("email").Parse(body))
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
