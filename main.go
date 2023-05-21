package main

import (
	"log"
	"net/http"

	"bps/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configure proxy trust
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment

	// Initialize the Gin router
	router := gin.Default()
	routes.DefineRoutes(router)
	log.Println("Server listening on port 8080...")
	log.Fatal(router.Run(":8080"))
}
