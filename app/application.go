package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()
	router.Run(":8080")
	fmt.Println("Hello World")
}