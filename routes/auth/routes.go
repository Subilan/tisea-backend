package auth

import "github.com/gin-gonic/gin"

func Bind(gin *gin.Engine) {
	group := gin.Group("/auth")
	group.POST("/auth/login", login)
	group.POST("/auth/register", register)
}
