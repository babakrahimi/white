package business

import "github.com/megaminx/white/cmd/core"

type UserRepository interface {
	Store(user *User) error
	FindUserByEmail(email string) (*User, error)
	SetUserStateToVerified(email string) error
}

type MailProvider interface {
	SendMailVerification(to string, token string) error
}

type CryptoProvider interface {
	GetRandomSalt() string
	GetEmailVerificationToken(data string) string
	VerifyToken(token string) error
	GetEmailFromToken(token string) string
	HashPassword(password, salt string) string
}

type UserActor interface {
	SignUp(email, password string) error
	VerifyEmail(token string) error
	SignIn(username, password string) error
	ChangePassword(username, oldPassword, newPassword string) error
}

type UserInteractor struct {
	UserRepository UserRepository
	MailProvider   MailProvider
	CryptoProvider CryptoProvider
}

type User struct {
	ID           string        `json:"_" bson:"_id"`
	Employee     core.Employee `json:"employee" bson:"employee"`
	IsAdmin      bool          `json:"isAdmin" bson:"isAdmin"`
	Username     string        `json:"username" bson:"username"`
	Email        string        `json:"email" bson:"email"`
	Password     string        `json:"_" bson:"password"`
	PasswordSalt string        `json:"_" bson:"passwordSalt"`
	Verified     bool          `json:"verified" bson:"verified"`
}

func (ui UserInteractor) SignUp(email, password string) error {
	u, err := ui.UserRepository.FindUserByEmail(email)
	if err != nil {
		return err
	}

	if u != nil {
		// Say nothing to user, for sake of security
		return nil
	}

	s := ui.CryptoProvider.GetRandomSalt()
	p := ui.CryptoProvider.HashPassword(password, s)
	user := &User{
		Username:     "",
		Email:        email,
		PasswordSalt: s,
		Password:     p,
		Verified:     false,
	}

	t := ui.CryptoProvider.GetEmailVerificationToken(email)
	if err = ui.UserRepository.Store(user); err != nil {
		return err
	}

	if err = ui.MailProvider.SendMailVerification(email, t); err != nil {
		return err
	}

	return nil
}

func (ui UserInteractor) VerifyEmail(token string) error {
	if err := ui.CryptoProvider.VerifyToken(token); err != nil {
		return err
	}
	email := ui.CryptoProvider.GetEmailVerificationToken(token)
	if err := ui.UserRepository.SetUserStateToVerified(email); err != nil {
		return err
	}
	return nil
}

func (ui UserInteractor) SignIn(username, password string) error {
	return nil
}

func (ui UserInteractor) ChangePassword(username, oldPassword, newPassword string) error {
	return nil
}
