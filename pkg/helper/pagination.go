package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationInput struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Search  string `json:"search"`
	OrderBy string `json:"orderBy"`
}

func Pagination(ctx *gin.Context) PaginationInput {
	limit := 10
	page := 1
	search := ""
	orderBy := "created_at desc"

	query := ctx.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "search":
			search = queryValue
		case "orderBy":
			orderBy = queryValue
		}
	}

	return PaginationInput{
		Limit:   limit,
		Page:    page,
		Search:  search,
		OrderBy: orderBy,
	}
}
