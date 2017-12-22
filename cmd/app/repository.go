package app

import (
	"gopkg.in/mgo.v2"
	"os"
	"errors"
)

type repository struct {
	db *mgo.Database
}

func newRepository() (*repository, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, errors.New("$DB_NAME must set")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("$MONGODB_URI must set")
	}

	s, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}

	s.SetMode(mgo.Monotonic, true)
	defer s.Close()

	r := &repository{
		db: s.Clone().DB(dbName),
	}

	return r, nil
}
