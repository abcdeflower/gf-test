package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Auth(r *ghttp.Request) {
	// token参数获取
	token := r.Header.Get("token")
	// token判断逻辑
	if token == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
