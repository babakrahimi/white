package interfaces

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type CryptoHandler struct {
	SecretKey string
}

func (ch *CryptoHandler) GetInvitationToken(email string) (string, error) {
	k := []byte(ch.SecretKey)

	claims := struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Duration(time.Hour*48) / time.Second),
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

func (ch *CryptoHandler) GetEmailVerificationToken(data string) string {
	panic("implement me")
}

func (ch *CryptoHandler) VerifyToken(token string) error {
	panic("implement me")
}

func (ch *CryptoHandler) GetEmailFromToken(token string) string {
	panic("implement me")
}

func (ch *CryptoHandler) HashPassword(password, salt string) string {
	saltedPass := password + salt
	h := hmac.New(sha256.New, []byte(ch.SecretKey))
	h.Write([]byte(saltedPass))
	return hex.EncodeToString(h.Sum(nil))
}