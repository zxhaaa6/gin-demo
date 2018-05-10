package main

import (
	"github.com/zxhaaa6/gin-demo/util"
	"github.com/zxhaaa6/gin-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/system"
	"log"
)

func main() {
	config := util.InitConf()
	util.LoadEnv()

	system.ConnectMongodbServer()

	gin.SetMode(util.GetEnv("GIN_MODE"))
	app := gin.Default()

	route.InitRouter(app)

	port := config.Port
	log.Println("[GIN]Service Listening on port:", port)
	app.Run(":" + port)
}
