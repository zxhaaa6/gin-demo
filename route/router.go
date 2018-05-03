package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/model/user"
)

func InitRouter(router *gin.Engine) {
	rUser := router.Group("user")
	user.InitController(rUser)
}
