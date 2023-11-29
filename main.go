package main

import "github.com/gin-gonic/gin"
import "tisea-backend/middlewares"
import "tisea-backend/routes/auth"

func main() {
	gin := gin.New()
	gin.Use(middlewares.WithGlobalHeaders())
	gin.Use(middlewares.WithLogger())

	auth.AddRoutes(gin)
}