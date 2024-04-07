package responses

import (
	"net/http"
)

// Response 结构表示API响应的基本结构
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// SuccessResponse 返回一个成功的响应
func SuccessResponse(data any, message string) Response {
	return Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  message,
	}
}

// NotFoundResponse 返回一个资源未找到的响应
func NotFoundResponse(message string) Response {
	return Response{
		Code: http.StatusNotFound,
		Data: nil,
		Msg:  "Not Found",
	}
}

// ErrorResponse 返回一个包含错误信息的响应
func ErrorResponse(data any, message string) Response {
	return Response{
		Code: http.StatusBadRequest,
		Data: data,
		Msg:  message,
	}
}
