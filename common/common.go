package common

import (
	"github.com/gogf/gf/net/ghttp"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	State   int         `json:"state"`
}

func Json(r *ghttp.Request, code int, message string, state int, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
		State:   state,
	})
}

func JsonExit(r *ghttp.Request, code int, msg string, state int, data ...interface{}) {
	Json(r, code, msg, state, data)
	r.Exit()
}
