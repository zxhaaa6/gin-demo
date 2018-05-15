package main

import (
	"github.com/zxhaaa6/gin-demo/util"
	"github.com/zxhaaa6/gin-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/system"
	log "github.com/sirupsen/logrus"
	"github.com/zxhaaa6/gin-demo/middleware"
)

func main() {
	// ============== init < config | env | log > ===============
	config := util.InitConf()
	util.LoadEnv()
	middleware.InitLog()

	// ================= connect mongodb | etc. =================
	system.ConnectMongodbServer()

	// ====================== init GIN APP ======================
	gin.SetMode(util.GetEnv("GIN_MODE"))

	app := gin.New()

	// ======================= middleware =======================
	if util.GetEnv("GIN_ENV") != "production" {
		app.Use(gin.Logger())
	}
	app.Use(gin.Recovery())

	// ========================= router =========================
	route.InitRouter(app)

	// ======================== startUp =========================
	port := config.Port
	log.Infoln("[GIN]Service Listening on port:", port)

	app.Run(":" + port)
}
