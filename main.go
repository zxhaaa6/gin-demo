package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/route"
)

func main() {
	app := gin.Default()

	route.InitRouter(app)

	port := "3000"
	app.Run(":" + port)
}
