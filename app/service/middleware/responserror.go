package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"web_test/common"
)

func ResponsErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln(common.JsonResponse{
			Code:    0,
			State:   0,
			Data:    nil,
			Message: "服务器居然开小差，请稍后再试！",
		})
	}
}
