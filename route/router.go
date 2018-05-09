package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/module/user"
)

func InitRouter(router *gin.Engine) {
	user.InitController(router.Group("user"))
}
