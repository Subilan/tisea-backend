package auth

import (
	"fmt"
	"tisea-backend/structs"
	"tisea-backend/utils/database"
	"tisea-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	var request structs.RegisterUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	if len(request.Username) == 0 || len(request.Email) == 0 || len(request.Password) == 0 {
		response.NG(ctx, fmt.Errorf("NOT_ENOUGH_ARGUMENT"), nil)
		return
	}

	if user, userErr := database.GetUserByUsername(request.Username); user != nil {
		if userErr == nil {
			response.NG(ctx, fmt.Errorf("DUPLICATE"), nil)
		} else {
			response.NG(ctx, userErr, nil)
		}
		return
	}

	registering, registeringErr := database.MakeRegisteringUser(request.Username, request.Email, request.Password)

	if registeringErr != nil {
		response.NG(ctx, registeringErr, nil)
		return
	}

	insertErr := database.InsertRegisteringUser(*registering)

	if insertErr != nil {
		response.NG(ctx, insertErr, nil)
		return
	}

	response.OK(ctx, "Registration finished successfully.", nil)
}
