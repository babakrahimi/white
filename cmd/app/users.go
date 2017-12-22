package app

import (
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "users"

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Username  string        `json:"username" bson:"username"`
}

func (a *App) GetUser(username string) (*User, error) {
	user := &User{}
	err := a.Repository.db.C(COLLECTION).Find(bson.M{"username": username}).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *App) GetUsers() ([]User, error) {
	var users []User
	err := a.Repository.db.C(COLLECTION).Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
