package main

import (
	"github.com/zxhaaa6/gin-demo/util"
	"github.com/zxhaaa6/gin-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/system"
)

func main() {
	config := util.InitConf()

	system.ConnectMongodbServer()

	app := gin.Default()

	route.InitRouter(app)

	port := config.Port
	app.Run(":" + port)
}
