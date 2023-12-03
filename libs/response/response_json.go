package response

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"net/http"
	"time"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type PaginationResponse struct {
	Meta       Meta        `json:"meta"`
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}

type Meta struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DebugParam string `json:"debug_param"`
	ServerTime string `json:"server_time"`
}

type Pagination struct {
	CurrentPage *int   `json:"current_page,omitempty"`
	PrevPage    *bool  `json:"prev_page,omitempty"`
	NextPage    *bool  `json:"next_page,omitempty"`
	CountFrom   *int64 `json:"count_from,omitempty"`
	CountUntil  *int64 `json:"count_until,omitempty"`
	PerPage     *int   `json:"per_page,omitempty"`
	PageCount   *int   `json:"page_count,omitempty"`
	TotalRecord *int64 `json:"total_record,omitempty"`
}

// SuccessRes is
func SuccessRes(ctx *fiber.Ctx, responseCode int, responseMessage, debugParam string, data interface{}) error {
	date := time.Now().Format(time.RFC1123)
	res := Response{
		Meta: Meta{
			Code:       responseCode,
			Message:    responseMessage,
			DebugParam: debugParam,
			ServerTime: date,
		},
		Data: data,
	}

	ctx.Set("date", date)
	return ctx.Status(responseCode).JSON(res)
}

// PaginationRes is
func PaginationRes(ctx *fiber.Ctx, responseCode int, responseMessage, debugParam string, page, perPage int, count int64, data interface{}) error {
	var (
		prevPage   bool
		nextPage   bool
		pageCount  int
		countFrom  int64
		countUntil int64
	)

	if page > 1 {
		prevPage = true
	}
	countFrom = (int64(page) - 1) * int64(perPage)
	countUntil = countFrom + int64(perPage)
	if countUntil >= count {
		countUntil = count
	}
	if countUntil < count {
		nextPage = true
	}
	pageCount = int(math.Ceil(float64(count) / float64(perPage)))

	date := time.Now().Format(time.RFC1123)
	res := PaginationResponse{
		Meta: Meta{
			Code:       responseCode,
			Message:    responseMessage,
			DebugParam: debugParam,
			ServerTime: date,
		},
		Data: data,
		Pagination: Pagination{
			CurrentPage: &page,
			PrevPage:    &prevPage,
			NextPage:    &nextPage,
			CountFrom:   &countFrom,
			CountUntil:  &countUntil,
			PerPage:     &perPage,
			PageCount:   &pageCount,
			TotalRecord: &count,
		},
	}

	ctx.Set("date", date)
	return ctx.Status(http.StatusOK).JSON(res)
}

// ErrorRes is
func ErrorRes(ctx *fiber.Ctx, responseCode int, responseMessage, debugParam string) error {
	date := time.Now().Format(time.RFC1123)
	res := Response{
		Meta: Meta{
			Code:       responseCode,
			Message:    responseMessage,
			DebugParam: debugParam,
			ServerTime: date,
		},
		Data: nil,
	}

	ctx.Set("date", date)
	return ctx.Status(responseCode).JSON(res)
}
