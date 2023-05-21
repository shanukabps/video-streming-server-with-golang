package controller

import (
	"log"
	"net/http"

	"bps/services"
	"bps/utils/response"

	"github.com/gin-gonic/gin"
)

// VideoController handles video-related operations
type VideoController struct{}

// GetVideo is the controller action for serving video content
func (vc *VideoController) GetVideo(c *gin.Context) {
	videoService := &services.VideoService{
		Path:        "./videoContents/test.mp4",
		ContentType: "video/mp4",
	}

	// Stream the video content using the VideoService
	err := videoService.StreamVideo(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}
