package interfaces

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/megaminx/white/cmd/business/user"
)

type (
	CryptoHandler struct {
		SecretKey string
	}
	invitationClaims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}
)

func NewCryptoHandler(secretKey string) user.CryptoHandler {
	crypto := &CryptoHandler{SecretKey: secretKey}
	return crypto
}

func (ch *CryptoHandler) VerifyInvitationToken(token string) (string, error) {

	t, err := jwt.ParseWithClaims(token, &invitationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ch.SecretKey), nil
	})
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return "", user.ErrExpiredInvitationToken
		}
	}
	if err != nil {
		return "", user.ErrInvalidInvitationToken
	}

	c := t.Claims.(*invitationClaims)
	return c.Email, nil
}

func (ch *CryptoHandler) GetInvitationToken(email string) (string, error) {
	k := []byte(ch.SecretKey)

	claims := &invitationClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "www.mahan.team",
			Subject:   "registration",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	st, err := t.SignedString(k)
	if err != nil {
		return "", err
	}
	return st, nil

}

func (ch *CryptoHandler) GetRandomSalt() string {
	bs := make([]byte, 16)
	rand.Read(bs)
	return hex.EncodeToString(bs)
}

func (ch *CryptoHandler) HashPassword(password, salt string) string {
	saltedPass := password + salt
	h := hmac.New(sha256.New, []byte(ch.SecretKey))
	h.Write([]byte(saltedPass))
	return hex.EncodeToString(h.Sum(nil))
}
