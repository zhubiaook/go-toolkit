package ginresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhubiaook/go-toolkit/errno"
)

// ErrResponse 定义了发生错误时的返回消息.
// swagger:response errResponse
type ErrResponse struct {
	// 业务错误码, 正确返回 0, 错误返回对应的业务错误码
	ErrorCode int `json:"error_code"`
	// 业务错误码对应的说明
	ResponseCode string `json:"response_code"`
	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message"`
}

// DefaultResponse 定义了正确返回消息.
// swagger:response okResponse
type DefaultResponse struct {
	ErrResponse

	Data any `json:"data"`
}

func Default(data any) DefaultResponse {
	return DefaultResponse{
		ErrResponse: ErrResponse{
			ErrorCode:    errno.OK.ErrorCode,
			ResponseCode: errno.OK.ResponseCode,
			Message:      errno.OK.Message,
		},
		Data: data,
	}
}

type Pagination struct {
	// 总条数
	Total int `json:"total"`
	// 总页数
	TotalPages int `json:"total_pages"`
	// 当前页
	PageNumber int `json:"page_number"`
	// 每页条数
	PageSize int `json:"page_size"`
}

type PaginateResponse struct {
	ErrResponse
	Pagination

	Data any
}

func Paginate[T any](pageNumber, pageSize int, total int, data []T) PaginateResponse {

	ok := errno.OK
	errResponse := ErrResponse{
		ErrorCode:    ok.ErrorCode,
		ResponseCode: ok.ResponseCode,
		Message:      ok.Message,
	}

	totalPages := total / pageSize
	if total%pageSize != 0 {
		totalPages++
	}

	return PaginateResponse{
		ErrResponse: errResponse,
		Pagination: Pagination{
			Total:      total,
			TotalPages: totalPages,
			PageNumber: pageNumber,
			PageSize:   len(data),
		},
		Data: data,
	}
}

// WriteResponse 将错误或响应数据写入 HTTP 响应主体。
// WriteResponse 使用 errno.Decode 方法，根据错误类型，尝试从 err 中提取业务错误码和错误信息.
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		hcode, eCode, rCode, message := errno.Decode(err)

		c.JSON(hcode, ErrResponse{
			ErrorCode:    eCode,
			ResponseCode: rCode,
			Message:      message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
