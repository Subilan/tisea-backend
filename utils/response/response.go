package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, ok bool, finishMessage string, errorObj error, data interface{}) {
	response := map[string]interface{}{
		"ok": ok,
		"message": finishMessage,
		"data": data,
		"timestamp": time.Now().UnixMilli(),
	}

	if errorObj != nil {
		response["error"] = errorObj.Error()
	} else {
		response["error"] = nil
	}

	c.JSON(http.StatusOK, response)
}

func OK(c *gin.Context, finishMessage string, data interface{}) {
	Respond(c, true, finishMessage, nil, data)
}

func NG(c *gin.Context, errorObj error, data interface{}) {
	Respond(c, false, "", errorObj, data)
}

