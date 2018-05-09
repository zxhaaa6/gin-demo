package user

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Controller struct {
}

var service = Service{}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}

	route.GET("/all", controller.getAllUsers)
}

func (r *Controller) getAllUsers(c *gin.Context) {
	users, err := service.getAllUsers()
	if err != nil {
		log.Fatal("err: ", err)
	}

	c.JSON(200, users)
}
