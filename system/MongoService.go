package system

import (
	"github.com/zxhaaa6/gin-demo/util"
	"gopkg.in/mgo.v2"
	"log"
)

var db *mgo.Database

func ConnectMongodbServer() {
	config := util.GetConf()
	mongoUrl := "mongodb://" + config.Mongodb.Host + ":" + config.Mongodb.Port
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		log.Println("[Mongodb]connect DB server fatal err: ", err)
		panic(err)
	}
	db = session.DB(config.Mongodb.DB)
	log.Println("[Mongodb]DB connection has been established successfully.")
}

func GetDB() *mgo.Database {
	return db
}
