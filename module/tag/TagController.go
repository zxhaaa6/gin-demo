package tag

import (
	"github.com/zxhaaa6/gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}
	route.GET("/all", controller.getAllTags)
}

func (r *Controller) getAllTags(c *gin.Context) {

	c.JSON(200, middleware.GenSimpleJson("all tag"))
}

