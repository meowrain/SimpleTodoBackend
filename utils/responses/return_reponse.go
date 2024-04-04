package responses

import (
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func SuccessResponse(data any, message string) Response {
	return Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  message,
	}
}
func NotFoundResponse(message string) Response {
	return Response{
		Code: http.StatusNotFound,
		Data: nil,
		Msg:  "Not Found",
	}
}
func ErrorResponse(data any, message string) Response {
	return Response{
		Code: http.StatusBadRequest,
		Data: data,
		Msg:  message,
	}
}
