// 封装 web 请求通用响应
package response

import (
	"encoding/json"
	"go_web_template/internal/businesserr"
)

// R 通用结果响应
type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// NewByError 根据 error 初始化返回响应
func NewByError(err error) *R[string] {
	// 如果是业务异常, 自动进行转换
	if busErr, ok := err.(*businesserr.E); ok {
		return &R[string]{Code: busErr.Code, Message: busErr.Message}
	}
	return &R[string]{Code: businesserr.EnumSystemError.Code, Message: err.Error()}
}

func Ok() *R[string] {
	return NewByError(businesserr.EnumSuccess)
}

func OkWithData[T any](data T) *R[T] {
	return &R[T]{
		Code:    businesserr.EnumSuccess.Code,
		Message: businesserr.EnumSuccess.Message,
		Data:    data,
	}
}

func Error(busErr *businesserr.E) *R[string] {
	return NewByError(busErr)
}

func ErrorWithMessage(busErr *businesserr.E, message string) *R[string] {
	return CustomError(busErr.Code, message)
}

func CustomError(code int, message string) *R[string] {
	return &R[string]{Code: code, Message: message}
}

func (r *R[T]) String() string {
	res, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(res)
}
