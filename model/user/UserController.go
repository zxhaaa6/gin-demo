package user

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}

	route.GET("/all", controller.getAllUsers)
}

func (r *Controller) getAllUsers(c *gin.Context) {
	c.String(200, "Success")
}
