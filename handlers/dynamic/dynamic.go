package dynamic

import "github.com/gin-gonic/gin"

func Bind(gin *gin.Engine) {
	dynamics := gin.Group("/dynamics")

	dynamics.POST("/", create)
	dynamics.DELETE("/", delete)
	dynamics.POST("/hide")
	dynamics.GET("/")
}