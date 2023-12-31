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
	soft := queries.Get("soft") == "1"

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

	var err error
	if soft {
		err = database.SetHiddenDynamic(parsed, true)
	} else {
		err = database.DeleteDynamic(parsed)
	}

	if err != nil {
		response.NG(ctx, err, nil)
		return
	}

	response.OK(ctx, fmt.Sprintf("Successfully deleted the dynamic post with ID %s", id), nil)
}
