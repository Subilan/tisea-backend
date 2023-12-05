package dynamic

import (
	"fmt"
	"strconv"
	"tisea-backend/utils/database"
	"tisea-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func delete(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()

	id := queries.Get("id")

	if len(id) == 0 {
		response.NG(ctx, fmt.Errorf("NOT_ENOUGH_ARGUMENT"), nil)
		return
	}

	// The type of id - uint64
	// See in structs/dynamic
	parsed, parseErr := strconv.ParseUint(id, 10, 64)

	if parseErr != nil {
		response.NG(ctx, fmt.Errorf("INVALID_ARGUMENT"), nil)
		return
	}

	if err := database.DeleteDynamic(parsed); err != nil {
		response.NG(ctx, err, nil)
		return
	}

	response.OK(ctx, fmt.Sprintf("Successfully deleted the dynamic post with ID %s", id), nil)
}