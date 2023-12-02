package cli

import (
	"database/sql"
	"fmt"
	"time"
	"tisea-backend/middlewares"
	"tisea-backend/routes/auth"
	"tisea-backend/utils/config"
	"tisea-backend/utils/database"

	"github.com/gin-gonic/gin"
)

func RunBackend() error {
	cfg := config.GetConfiguration()
	pool, poolErr := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Dbname))

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

	auth.Bind(gin)

	return nil
}
