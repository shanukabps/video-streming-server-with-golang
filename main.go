package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Video represents the video file
type Video struct {
	Path        string
	ContentType string
}

// VideoController handles video-related operations
type VideoController struct {
	Video *Video
}

// GetVideo is the controller action for serving video content
func (vc *VideoController) GetVideo(c *gin.Context) {
	// Open the video file
	file, err := os.Open(vc.Video.Path)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the Content-Type header
	c.Header("Content-Type", vc.Video.ContentType)

	/*
	file.Stat() is a method call on the file object, which returns a FileInfo object and an error. The FileInfo object 
	contains information about the file, such as its size, modification time, and mode.

	The purpose of retrieving the file information in this context is to obtain the modification time of the file.
	It is used later in the code when calling http.ServeContent() to serve the video content. The modification time 
	is passed as an argument to http.ServeContent(), which helps with caching and handling conditional requests.
	*/
	// Stream the file content
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	/*
	The empty string "" is passed as the third parameter, which represents the name of the file. In this case, an empty string is 
	used because the filename is not explicitly provided.fileInfo.ModTime() is the modification time of the video file. It is
	obtained from the fileInfo object, which was retrieved earlier using file.Stat(). The modification time is used to 
	set the Last-Modified header in the response, allowing clients to perform conditional requests and cache the video content.

	By calling http.ServeContent(), the video content is streamed from the file to the client's http.ResponseWriter.
	The function takes care of setting the appropriate headers, including the Content-Type based on the file's extension
	and Content-Length based on the file's size. It also handles partial content requests (range requests) and sends
	the appropriate status codes.
	*/
	http.ServeContent(c.Writer, c.Request, "", fileInfo.ModTime(), file)
}

func main() {
	// Configure proxy trust
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment
	// Create the video model
	video := &Video{
		Path:        "./video.mp4",
		ContentType: "video/mp4",
	}

	// Create the video controller
	videoController := &VideoController{
		Video: video,
	}

	// Initialize the Gin router
	router := gin.Default()

	// Register the video route
	router.GET("/video", videoController.GetVideo)

	log.Println("Server listening on port 8080...")
	log.Fatal(router.Run(":8080"))
}
