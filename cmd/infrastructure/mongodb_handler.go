package infrastructure

import (
	"github.com/megaminx/white/cmd/interfaces"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	MongodbHandler struct {
		*mgo.Database
	}
)

func (mh *MongodbHandler) Store(data interface{}) error {
	if err := mh.C("invitations").Insert(data); err != nil {
		return err
	}
	return nil
}

func (mh *MongodbHandler) FindOne(conditions map[string]interface{}, result interface{}) error {
	err := mh.C("invitations").Find(bson.M(conditions)).One(result)
	if err == mgo.ErrNotFound {
		return interfaces.ErrNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func NewMongodbHandler(url, dbName string) (*MongodbHandler, error) {
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	s.SetMode(mgo.Monotonic, true)

	mh := &MongodbHandler{
		Database: s.Clone().DB(dbName),
	}
	return mh, nil
}
