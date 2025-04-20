// common/util/response.go
package util

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Response 统一响应结构
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseSuccess 成功响应
func HttpOk(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, &Response{
		Status:  1,
		Message: "",
		Data:    data,
	})
}

// ResponseError 错误响应
func HttpError(w http.ResponseWriter, message string) {
	httpx.OkJson(w, &Response{
		Status:  0,
		Message: message,
		Data:    map[string]interface{}{},
	})
}

func HttpSend(w http.ResponseWriter, err error, data interface{}) {
	if err != nil {
		// 错误响应
		httpx.OkJson(w, &Response{
			Status:  0,
			Message: err.Error(),
			Data:    map[string]interface{}{},
		})
		return
	}

	// 成功响应
	httpx.OkJson(w, &Response{
		Status:  1,
		Message: "",
		Data:    data,
	})
}
