package infrastructure

import (
	"fmt"
	"os"
	"strings"
)

type (
	ENVVariableNotSetError struct {
		Name string
	}
	Config struct {
		Port, JWTSecretKey string
		AllowedOrigins     []string
		Email              *EmailConfig
		DB                 *Database
	}
	Database struct {
		URL, DBName string
	}
	EmailConfig struct {
		ServerAddress           string
		RegistrationAddress     string
		RegistrationPassword    string
		RegistrationRedirectURL string
	}
)

func (e *ENVVariableNotSetError) Error() string {
	text := fmt.Sprintf("$%s should not be empty", e.Name)
	return text
}

func GetENVVariable(name string) (string, error) {
	name = strings.ToUpper(name)
	v := os.Getenv(name)
	if v == "" {
		return "", &ENVVariableNotSetError{name}
	}
	return v, nil
}

func GetConfig() (*Config, error) {
	p, err := GetENVVariable("port")
	if err != nil {
		return nil, err
	}
	ao, err := GetENVVariable("allowed_origins")
	if err != nil {
		return nil, err
	}
	jwt, err := GetENVVariable("jwt_secret")
	if err != nil {
		return nil, err
	}
	rru, err := GetENVVariable("registration_redirect_url")
	if err != nil {
		return nil, err
	}
	sa, err := GetENVVariable("email_server")
	if err != nil {
		return nil, err
	}
	ra, err := GetENVVariable("email_address_reg")
	if err != nil {
		return nil, err
	}
	rp, err := GetENVVariable("email_password_reg")
	if err != nil {
		return nil, err
	}
	du, err := GetENVVariable("mongodb_uri")
	if err != nil {
		return nil, err
	}
	dn, err := GetENVVariable("db_name")
	if err != nil {
		return nil, err
	}

	c := &Config{
		Port:           p,
		AllowedOrigins: strings.Split(ao, ","),
		JWTSecretKey:   jwt,
		Email: &EmailConfig{
			ServerAddress:           sa,
			RegistrationAddress:     ra,
			RegistrationPassword:    rp,
			RegistrationRedirectURL: rru,
		},
		DB: &Database{
			URL:    du,
			DBName: dn,
		},
	}
	return c, nil
}
