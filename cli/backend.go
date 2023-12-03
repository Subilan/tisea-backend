package cli

import (
	"database/sql"
	"fmt"
	"time"
	"tisea-backend/handlers/auth"
	"tisea-backend/middlewares"
	"tisea-backend/utils/config"
	"tisea-backend/utils/database"

	"github.com/gin-gonic/gin"
)

func RunBackend() error {
	cfg := config.GetConfiguration()

	// Database initialization

	pool, poolErr := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Dbname))

	if poolErr != nil {
		return poolErr
	}

	database.Pool = pool
	connErr := database.Pool.Ping()
	if connErr != nil {
		return connErr
	}
	database.Pool.SetConnMaxLifetime(3 * time.Minute)
	database.Pool.SetMaxOpenConns(10)
	database.Pool.SetMaxIdleConns(10)

	gin := gin.New()
	gin.Use(middlewares.WithGlobalHeaders())
	gin.Use(middlewares.WithLogger())

	// Handler bindings

	auth.Bind(gin)

	// Necessary to block the calling goroutine
	gin.Run(":3000")

	return nil
}
