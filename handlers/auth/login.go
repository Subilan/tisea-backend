package auth

import (
	"fmt"
	"tisea-backend/structs"
	"tisea-backend/utils/database"
	"tisea-backend/utils/response"
	"tisea-backend/utils/security"

	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	var request structs.LoginUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	if len(request.Username) == 0 || len(request.Password) == 0 {
		response.NG(ctx, fmt.Errorf("NOT_ENOUGH_ARGUMENT"), nil)
		return
	}

	user, userErr := database.GetUserByUsername(request.Username)

	if user == nil {
		if userErr == nil {
			response.NG(ctx, fmt.Errorf("NOT_FOUND"), nil)
		} else {
			response.NG(ctx, userErr, nil)
		}
		return
	}

	if security.CompareHash(request.Password, user.Hash) {
		token, tokenErr := security.GenerateTokenForUser(request.Username, request.Remembered)
		if tokenErr != nil {
			response.NG(ctx, tokenErr, nil)
		} else {
			response.OK(ctx, "User successfully logged in.", token)
		}
	} else {
		response.NG(ctx, fmt.Errorf("INVALID_CREDENTIALS"), nil)
	}
}
