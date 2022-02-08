package app

import (
	"userapi/logger"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8000")
}