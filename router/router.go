package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"web_test/app/api/chat"
	"web_test/app/service/middleware"
)

func init() {
	s := g.Server()

	// 全局middleware
	s.Use(middleware.CORS, middleware.Log)
	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlChat := new(chat.Controller)
		group.Group("/test", func(group *ghttp.RouterGroup) {
			// 局部middleware
			// group.Middleware(middleware.Auth)
			group.ALL("/chat", ctlChat.WebSocket)
		})
	})
}
