package dynamic

import (
	"tisea-backend/structs"
	"tisea-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var request structs.PostingDynamic

	if err := ctx.BindJSON(&request); err != nil {
		response.NG(ctx, err, nil)
		return
	}
}