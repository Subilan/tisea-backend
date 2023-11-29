package auth

import (
	"encoding/json"
	"tisea-backend/structs"
	"tisea-backend/utils"

	"github.com/gin-gonic/gin"
)

func handle(ctx *gin.Context) {
	var request structs.RUserRegister

	if err := ctx.BindJSON(&request); err != nil {
		utils.NG(ctx, err.Error(), nil)
		return
	}

	if len(request.Username) == 0 || len(request.Email) == 0 || len(request.Password) == 0 {
		utils.NG(ctx, "Not enough argument.", nil)
	}

	selectUser := 
}
