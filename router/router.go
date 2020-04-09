package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"web_test/app/api/chat"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlChat := new(chat.Controller)
		group.Group("/test", func(group *ghttp.RouterGroup) {
			group.ALL("/chat", ctlChat.WebSocket)
		})
	})
}
