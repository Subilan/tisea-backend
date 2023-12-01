package main

import (
	"database/sql"
	"fmt"
	"time"
	"tisea-backend/middlewares"
	"tisea-backend/utils/config"
	"tisea-backend/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.New()
	gin.Use(middlewares.WithGlobalHeaders())
	gin.Use(middlewares.WithLogger())

	cfg := config.GetConfiguration()
	pool, poolErr := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Dbname))
	
	if poolErr != nil {
		panic(poolErr)
	}
	
	database.Pool = pool
	connErr := database.Pool.Ping()
	if connErr != nil {
		panic(connErr)
	}
	database.Pool.SetConnMaxLifetime(3 * time.Minute)
	database.Pool.SetMaxOpenConns(10)
	database.Pool.SetMaxIdleConns(10)
}