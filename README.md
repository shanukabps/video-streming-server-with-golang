# video-streming-server-with-golang
vide streaming server with golang, gin freamwork, MVC pattern

## this go lang API write using gin and used MVC design pattern


### put the video file  the videoContes folder rename video as test
go version 1.1.8

for run the project 
go run main.go
build the project 

go build -o bps

# {{}}/video reponse is 304 
A 304 response status code is returned when the client's cached version of a requested resource is still considered valid and can be used. It indicates that the requested resource has not been modified since the client last accessed it, so there is no need to send the entire response again. Instead, the server can respond with a 304 status code and an empty response body, instructing the client to use its cached version.

In the context of your "/video" route, it seems that the client is sending an "If-Modified-Since" header in the request. This header contains a timestamp indicating the last time the client accessed the resource. The server checks this timestamp against the actual modification time of the resource. If the resource has not been modified since the provided timestamp, the server can respond with a 304 status code.

!! If you want to ensure that the server always sends the complete video content with a 200 status code, regardless of the client's cache, you can modify your code to disable cache control for the "/video" route. Here's an example:
func (vc *VideoController) GetVideo(c *gin.Context) {
	// ...

	// Disable cache control for the response
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	// Stream the file content
	// ...
}

bpsshanuka...
