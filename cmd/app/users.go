package app

import (
	"gopkg.in/mgo.v2/bson"
	"crypto/sha256"
	"encoding/hex"
	"crypto/rand"
	"gopkg.in/mgo.v2"
	"fmt"
)

const (
	collection = "users"
)

type DuplicateUsernameError struct {
	Username string
}

func (e DuplicateUsernameError) Error() string {
	return fmt.Sprintf("نام کاربری %v قبلا در سیستم ثبت شده", e.Username)
}

type User struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	FirstName       string        `json:"firstName" bson:"firstName"`
	LastName        string        `json:"lastName" bson:"lastName"`
	Username        string        `json:"username" bson:"username"`
	Password        string        `json:"password" bson:"password"`
	PasswordSalt    string        `json:"_" bson:"passwordSalt"`
	Email           string        `json:"email" bson:"email"`
	BankCardNumbers []string      `json:"bankCardNumbers" bson:"bankCardNumbers"`
}

func (a *App) GetUser(username string) (*User, error) {
	user := &User{}
	err := a.Repository.db.C(collection).Find(bson.M{"username": username}).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *App) GetUsers() ([]User, error) {
	var users []User
	err := a.Repository.db.C(collection).Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *App) CreateUser(username, password string) error {
	ps := makePasswordSalt()
	p := makeSecurePassword(password, ps)

	user := &User{
		ID:           bson.NewObjectId(),
		Username:     username,
		PasswordSalt: ps,
		Password:     p,
	}

	if err := a.Repository.db.C(collection).Insert(user); err != nil {
		if mgo.IsDup(err) {
			return DuplicateUsernameError{ username}
		}
		return err
	}
	return nil
}

func makeSecurePassword(password, salt string) string {
	saltedPass := password + salt

	h := sha256.New()
	h.Write([]byte(saltedPass))

	hashedPassword := hex.EncodeToString(h.Sum(nil))

	return hashedPassword
}

func makePasswordSalt() string {
	bs := make([]byte, 16)
	rand.Read(bs)
	salt := hex.EncodeToString(bs)
	return salt
}
