package routes

import "github.com/gin-gonic/gin"

func DefineRoutes(router *gin.Engine) {
	registerVideoRoutes(router)
}
