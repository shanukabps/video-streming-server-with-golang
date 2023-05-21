package routes

import (
	"bps/controller"

	"github.com/gin-gonic/gin"
)

// Create the video controller
var videoController = &controller.VideoController{}

// registerVideoRoutes registers video routes on the given router
func registerVideoRoutes(router *gin.Engine) {
	router.GET("/video", videoController.GetVideo)
}
