package response

import (
	"net/http"
)

// APIResponse 基础的HttpApi返回结构
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"result"`
}

// ForbiddenRequest 没有权限
func ForbiddenRequest() APIResponse {
	return APIResponse{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
		Data:    nil,
	}
}

// BadRequest 错误的请求, 一般是逻辑上的，比如说传少参数了
func BadRequest(message error) APIResponse {
	response := APIResponse{
		Code:    http.StatusBadRequest,
		Message: message.Error(),
		Data:    nil,
	}

	return response
}

// Success 成功的请求
func Success(data interface{}) APIResponse {
	return APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

// SystemError 系统错误
func SystemError(message error) APIResponse {
	return APIResponse{
		Code:    http.StatusInternalServerError,
		Message: message.Error(),
		Data:    nil,
	}
}
