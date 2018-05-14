package main

import (
	"github.com/zxhaaa6/gin-demo/util"
	"github.com/zxhaaa6/gin-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/system"
	log "github.com/sirupsen/logrus"
	"github.com/zxhaaa6/gin-demo/middleware"
)

func main() {
	config := util.InitConf()
	util.LoadEnv()

	middleware.InitLog()

	system.ConnectMongodbServer()

	gin.SetMode(util.GetEnv("GIN_MODE"))

	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	route.InitRouter(app)

	port := config.Port

	log.Infoln("[GIN]Service Listening on port:", port)
	app.Run(":" + port)
}
