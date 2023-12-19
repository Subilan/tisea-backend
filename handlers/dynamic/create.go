package dynamic

import (
	"tisea-backend/structs"
	"tisea-backend/utils/database"
	"tisea-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var dynamic structs.CreateDynamicRequest

	if err := ctx.BindJSON(&dynamic); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	if err := database.InsertPostingDynamic(dynamic); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	response.OK(ctx, "Successfully posted the dynamic.", nil)
}
