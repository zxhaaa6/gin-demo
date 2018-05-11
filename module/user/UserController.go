package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/zxhaaa6/gin-demo/middleware"
	"github.com/zxhaaa6/gin-demo/util"
)

type Controller struct {
	UserService Service
}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}
	controller.UserService = InitService()
	route.GET("/all", controller.getAllUsers)
	route.GET("/one/:id", controller.getOneUser)
}

func (r *Controller) getOneUser(c *gin.Context) {
	id := c.Param("id")
	user, err := r.UserService.findUserById(id)
	if err != nil {
		log.Println("err: ", err)
	}
	c.JSON(200, middleware.GenSimpleJson(user))
}

func (r *Controller) getAllUsers(c *gin.Context) {
	users, err := r.UserService.getAllUsers()
	if err != nil {
		log.Println("err: ", err)
		c.JSON(200, middleware.GenErrorJson(util.GenUniError(500, err.Error())))
		return
	}

	c.JSON(200, middleware.GenSimpleJson(users))
}
