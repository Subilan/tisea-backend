package dynamic

import (
	"tisea-backend/structs"
	"tisea-backend/utils/database"
	"tisea-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var request structs.PostingDynamic

	if err := ctx.BindJSON(&request); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	dynamic := database.MakePostingDynamic(request.Title, request.Content, request.Author, request.Categories, request.Tags)

	if err := database.InsertPostingDynamic(*dynamic); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	response.OK(ctx, "Successfully posted the dynamic.", nil)
}