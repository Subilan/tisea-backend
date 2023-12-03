package auth

import "github.com/gin-gonic/gin"

func Bind(gin *gin.Engine) {
	group := gin.Group("/auth")
	group.POST("/login", login)
	group.POST("/register", register)
}
