package dynamic

import "github.com/gin-gonic/gin"

func Bind(gin *gin.Engine) {
	dynamics := gin.Group("/dynamics")

	dynamics.POST("/create")
	dynamics.POST("/delete")
	dynamics.POST("/hide")
	dynamics.GET("/")
}