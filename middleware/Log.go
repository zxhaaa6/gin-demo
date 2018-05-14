package middleware

import (
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/gin-gonic/gin"
)

func InitLog() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})

	// or file
	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
