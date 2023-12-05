package dynamic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tisea-backend/utils/database"
	"tisea-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()

	_skip := queries.Get("skip") // Required. Amount of items that need to be skipped before selecting
	_limit := queries.Get("limit") // Required. Amount of items to be selected
	by := queries.Get("by") // Required. The classification standard type of the selection. Must be either 'author' or 'id'
	with := queries.Get("with") // Required. The classification standard value of the selection. For 'author', must be a case-sensitive string. For 'id', must be a valid uint64.
	order := queries.Get("order") // Optional. The order of the selected values. If provided, it must be either 'asc' (stands for ascending) or 'desc' (stands for descending). If not provided, default to 'desc'.

	if by != "author" && by != "id" {
		response.NG(ctx, fmt.Errorf("INVALID_ARGUMENT `by`"), nil)
		return
	}

	if len(order) > 0 {
		if order != "asc" && order != "desc" {
			response.NG(ctx, fmt.Errorf("INVALID_ARGUMENT `order`"), nil)
			return
		}
	} else {
		order = "desc" // default value
	}

	skip, err := strconv.Atoi(_skip)

	if err != nil {
		response.NG(ctx, fmt.Errorf("INVALID_ARGUMENT `skip`"), nil)
		return
	}

	limit, err := strconv.Atoi(_limit)

	if err != nil {
		response.NG(ctx, fmt.Errorf("INVALID_ARGUMENT `limit`"), nil)
		return
	}

	if by == "author" {
		results, resultErr := database.GetDynamicsByAuthor(with, limit, skip, order == "desc")
		if resultErr != nil {
			response.NG(ctx, resultErr, nil)
		}

		var finalResults []string

		for i := range results {
			marshed, err := json.Marshal(*results[i])
			if err != nil {
				response.NG(ctx, err, nil)
				return
			}
			finalResults = append(finalResults, string(marshed))
		}

		response.OK(ctx, "Got", finalResults)
	} else {
		// Convert `with` to uint64 number.
		with, err := strconv.ParseUint(with, 10, 64)

		if err != nil {
			response.NG(ctx, fmt.Errorf("INVALID_ARGUMENT `with`"), nil)
			return
		}

		// Get single result
		result, resultErr := database.GetDynamicsByID(with)

		if resultErr != nil {
			response.NG(ctx, resultErr, nil)
			return
		}

		finalResult, err := json.Marshal(*result)

		if err != nil {
			response.NG(ctx, err, nil)
			return
		}

		response.OK(ctx, "Got", finalResult)
	}
}