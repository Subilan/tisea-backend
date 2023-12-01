package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, ok bool, finishMessage string, errorCause string, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": ok,
		"message": finishMessage,
		"cause": errorCause,
		"data": data,
		"timestamp": time.Now().UnixMilli(),
	})
}

func OK(c *gin.Context, finishMessage string, data interface{}) {
	Respond(c, true, finishMessage, "", data)
}

func NG(c *gin.Context, errorCause string, data interface{}) {
	Respond(c, false, "", errorCause, data)
}

