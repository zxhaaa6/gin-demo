package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/module/user"
	"github.com/zxhaaa6/gin-demo/module/tag"
)

func InitRouter(router *gin.Engine) {
	user.InitController(router.Group("user"))
	tag.InitController(router.Group("tag"))
}
