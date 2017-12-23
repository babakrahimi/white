package app

import (
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"

type User struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	FirstName       string        `json:"firstName" bson:"firstName"`
	LastName        string        `json:"lastName" bson:"lastName"`
	Username        string        `json:"username" bson:"username"`
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

func (a *App) AddUsers(users ...*User) ([]*User, error) {
	for _, u := range users {
		u.ID = bson.NewObjectId()
	}

	if err := a.Repository.db.C(collection).Insert(users); err != nil {
		return nil, err
	}
	return users, nil
}