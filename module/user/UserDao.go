package user

import (
	"github.com/zxhaaa6/gin-demo/system"
	"gopkg.in/mgo.v2/bson"
	"github.com/zxhaaa6/gin-demo/model"
	"gopkg.in/mgo.v2"
)

type Dao struct {
	Collection *mgo.Collection
}

func InitDao() Dao {
	dao := Dao{}
	dao.Collection = system.GetDB().C("users")
	return dao
}

func (r *Dao) getAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.Collection.Find(bson.M{}).All(&users)
	return users, err
}
