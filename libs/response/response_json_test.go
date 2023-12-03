package response

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessRes(t *testing.T) {
	// argument
	type args struct {
		ctx             *fiber.Ctx
		responseCode    int
		responseMessage string
		debugParam      string
		data            interface{}
	}

	// test case
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// success scenario: test with nil data and nil error
		{
			name: "Success_With_Nil_Data_And_Error",
			args: args{
				responseCode:    http.StatusOK,
				responseMessage: "success",
				debugParam:      "",
				data:            nil,
			},
			wantErr: assert.NoError,
		},
		// success scenario: test with data and nil error
		{
			name: "Success_With_Data_And_Nil_Error",
			args: args{
				responseCode:    http.StatusOK,
				responseMessage: "success",
				debugParam:      "",
				data: struct {
					ID   string
					Name string
				}{
					ID:   "4624712e-d1a7-428c-8a72-84ec4ad79ab9",
					Name: "Unit Testing",
				},
			},
			wantErr: assert.NoError,
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fiber.New()
			f.Get("/", func(ctx *fiber.Ctx) error {
				err := SuccessRes(ctx, tt.args.responseCode, tt.args.responseMessage, tt.args.debugParam, tt.args.data)
				return err
			})
			res, err := f.Test(httptest.NewRequest(http.MethodGet, "/", nil))
			// error
			assert.NoError(t, err)
			// header
			assert.Equal(t, res.StatusCode, tt.args.responseCode)
			assert.NotNil(t, res.Header.Get("date"))
			// body
			resBody, err := readBodyRes(res)
			assert.NoError(t, err)
			assert.Equal(t, resBody.Meta.Code, tt.args.responseCode)
			assert.Equal(t, resBody.Meta.Message, tt.args.responseMessage)
			assert.Equal(t, resBody.Meta.DebugParam, tt.args.debugParam)
			assert.NotNil(t, resBody.Meta.ServerTime)
			if tt.args.data != nil {
				assert.NotNil(t, resBody.Data)
			} else {
				assert.Nil(t, resBody.Data)
			}
		})
	}
}

func TestPaginationRes(t *testing.T) {
	// argument
	type args struct {
		ctx             *fiber.Ctx
		responseCode    int
		responseMessage string
		debugParam      string
		page            int
		perPage         int
		count           int64
		data            interface{}
	}
	type resPagination struct {
		currentPage int
		prevPage    bool
		nextPage    bool
		countFrom   int64
		countUntil  int64
		perPage     int
		pageCount   int
		totalRecord int64
	}

	// test case
	tests := []struct {
		name              string
		args              args
		wantResPagination resPagination
		wantErr           assert.ErrorAssertionFunc
	}{
		// success scenario: test with nil data, nil error, and first page
		{
			name: "Success_With_Nil_Data_And_Error_First_Page",
			args: args{
				responseCode:    http.StatusOK,
				responseMessage: "success",
				debugParam:      "",
				page:            1,
				perPage:         10,
				count:           int64(25),
				data:            nil,
			},
			wantResPagination: resPagination{
				currentPage: 1,
				prevPage:    false,
				nextPage:    true,
				countFrom:   int64(0),
				countUntil:  int64(10),
				perPage:     10,
				pageCount:   3,
				totalRecord: int64(25),
			},
			wantErr: assert.NoError,
		},
		// success scenario: test with nil data, nil error, and second page
		{
			name: "Success_With_Nil_Data_And_Error_Second_Page",
			args: args{
				responseCode:    http.StatusOK,
				responseMessage: "success",
				debugParam:      "",
				page:            2,
				perPage:         10,
				count:           int64(25),
				data:            nil,
			},
			wantResPagination: resPagination{
				currentPage: 2,
				prevPage:    true,
				nextPage:    true,
				countFrom:   int64(10),
				countUntil:  int64(20),
				perPage:     10,
				pageCount:   3,
				totalRecord: int64(25),
			},
			wantErr: assert.NoError,
		},
		// success scenario: test with nil data, nil error, and third page
		{
			name: "Success_With_Nil_Data_And_Error_Third_Page",
			args: args{
				responseCode:    http.StatusOK,
				responseMessage: "success",
				debugParam:      "",
				page:            3,
				perPage:         10,
				count:           25,
				data:            nil,
			},
			wantResPagination: resPagination{
				currentPage: 3,
				prevPage:    true,
				nextPage:    false,
				countFrom:   int64(20),
				countUntil:  int64(25),
				perPage:     10,
				pageCount:   3,
				totalRecord: int64(25),
			},
			wantErr: assert.NoError,
		},
		// success scenario: test with data, nil error, and third page
		{
			name: "Success_With_Data_And_Nil_Error_Third_Page",
			args: args{
				responseCode:    http.StatusOK,
				responseMessage: "success",
				debugParam:      "",
				page:            3,
				perPage:         10,
				count:           int64(25),
				data: []struct {
					ID   string
					Name string
				}{
					{
						ID:   "4624712e-d1a7-428c-8a72-84ec4ad79ab9",
						Name: "Unit Testing",
					},
					{
						ID:   "4624712e-d1a7-428c-8a72-84ec4ad79ab9",
						Name: "Unit Testing",
					},
				},
			},
			wantResPagination: resPagination{
				currentPage: 3,
				prevPage:    true,
				nextPage:    false,
				countFrom:   int64(20),
				countUntil:  int64(25),
				perPage:     10,
				pageCount:   3,
				totalRecord: int64(25),
			},
			wantErr: assert.NoError,
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fiber.New()
			f.Get("/", func(ctx *fiber.Ctx) error {
				err := PaginationRes(ctx, tt.args.responseCode, tt.args.responseMessage, tt.args.debugParam, tt.args.page, tt.args.perPage, tt.args.count, tt.args.data)
				return err
			})
			resp, err := f.Test(httptest.NewRequest(http.MethodGet, "/", nil))
			// error
			assert.NoError(t, err)
			// header
			assert.Equal(t, resp.StatusCode, tt.args.responseCode)
			assert.NotNil(t, resp.Header.Get("date"))
			// body
			resBody, err := readPaginationBodyRes(resp)
			assert.NoError(t, err)
			assert.Equal(t, resBody.Meta.Code, tt.args.responseCode)
			assert.Equal(t, resBody.Meta.Message, tt.args.responseMessage)
			assert.Equal(t, resBody.Meta.DebugParam, tt.args.debugParam)
			assert.NotNil(t, resBody.Meta.ServerTime)
			if tt.args.data != nil {
				assert.NotNil(t, resBody.Data)
			} else {
				assert.Nil(t, resBody.Data)
			}
			assert.Equal(t, resBody.Pagination.CurrentPage, &tt.wantResPagination.currentPage)
			assert.Equal(t, resBody.Pagination.PrevPage, &tt.wantResPagination.prevPage)
			assert.Equal(t, resBody.Pagination.NextPage, &tt.wantResPagination.nextPage)
			assert.Equal(t, resBody.Pagination.CountFrom, &tt.wantResPagination.countFrom)
			assert.Equal(t, resBody.Pagination.CountUntil, &tt.wantResPagination.countUntil)
			assert.Equal(t, resBody.Pagination.PerPage, &tt.wantResPagination.perPage)
			assert.Equal(t, resBody.Pagination.PageCount, &tt.wantResPagination.pageCount)
			assert.Equal(t, resBody.Pagination.TotalRecord, &tt.wantResPagination.totalRecord)
		})
	}
}

func TestErrorRes(t *testing.T) {
	// argument
	type args struct {
		ctx             *fiber.Ctx
		responseCode    int
		responseMessage string
		debugParam      string
	}

	// test case
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// success scenario: test with nil data and error
		{
			name: "Success_With_Nil_Data_And_Not_Nil_Error",
			args: args{
				responseCode:    http.StatusInternalServerError,
				responseMessage: "failed",
				debugParam:      "unit testing failed",
			},
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fiber.New()
			f.Get("/", func(ctx *fiber.Ctx) error {
				err := ErrorRes(ctx, tt.args.responseCode, tt.args.responseMessage, tt.args.debugParam)
				return err
			})
			res, err := f.Test(httptest.NewRequest(http.MethodGet, "/", nil))
			// error
			assert.NoError(t, err)
			// header
			assert.Equal(t, res.StatusCode, tt.args.responseCode)
			assert.NotNil(t, res.Header.Get("date"))
			// body
			resBody, err := readBodyRes(res)
			assert.NoError(t, err)
			assert.Equal(t, resBody.Meta.Code, tt.args.responseCode)
			assert.Equal(t, resBody.Meta.Message, tt.args.responseMessage)
			assert.Equal(t, resBody.Meta.DebugParam, tt.args.debugParam)
			assert.NotNil(t, resBody.Meta.ServerTime)
			assert.Nil(t, resBody.Data)
		})
	}
}

// readBodyRes is
func readBodyRes(res *http.Response) (resBody Response, err error) {
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return resBody, err
	}
	err = json.Unmarshal(response, &resBody)
	if err != nil {
		return resBody, err
	}
	return resBody, nil
}

// readPaginationBodyRes is
func readPaginationBodyRes(res *http.Response) (resBody PaginationResponse, err error) {
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return resBody, err
	}
	err = json.Unmarshal(response, &resBody)
	if err != nil {
		return resBody, err
	}
	return resBody, nil
}
