package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/route"
	"github.com/zxhaaa6/gin-demo/util"
	"fmt"
)

func main() {
	config := util.GetConf()
	fmt.Println(*config)
	app := gin.Default()

	route.InitRouter(app)

	port := "3000"
	app.Run(":" + port)
}
