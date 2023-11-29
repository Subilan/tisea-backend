package auth

import "github.com/gin-gonic/gin"

func AddRoutes(gin *gin.Engine) {
	group := gin.Group("/auth")
	group.POST("/auth/login")
	group.POST("/auth/register")
}