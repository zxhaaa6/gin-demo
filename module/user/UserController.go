package user

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Controller struct {
	UserService Service
}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}
	controller.UserService = InitService()
	route.GET("/all", controller.getAllUsers)
}

func (r *Controller) getAllUsers(c *gin.Context) {
	users, err := r.UserService.getAllUsers()
	if err != nil {
		log.Fatal("err: ", err)
	}

	c.JSON(200, users)
}
