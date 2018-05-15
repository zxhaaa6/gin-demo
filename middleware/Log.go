package middleware

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/gin-demo/util"
	"os"
)

func InitLog() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", os.FileMode(0777))
	}

	ginEnv := util.GetEnv("GIN_ENV")

	if ginEnv != "production" {
		log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
		log.SetOutput(os.Stdout)
	} else {
		log.SetFormatter(&log.TextFormatter{})
		file, err := os.OpenFile("log/log.log", os.O_WRONLY|os.O_CREATE, os.FileMode(0755))
		if err != nil {
			log.Fatal("open log file fatal err: ", err)
		}
		log.SetOutput(file)
	}

	log.SetLevel(log.InfoLevel)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
