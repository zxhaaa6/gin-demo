package user

import (
	"github.com/zxhaaa6/gin-demo/system"
	"gopkg.in/mgo.v2/bson"
	"github.com/zxhaaa6/gin-demo/model"
)

type Dao struct {
}

func (r *Dao) getAllUsers() ([]model.User, error) {
	collection := system.GetDB().C("users")
	var users []model.User
	err :=collection.Find(bson.M{}).All(&users)
	return users, err
}
