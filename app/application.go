package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nictes1/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8082")
}
