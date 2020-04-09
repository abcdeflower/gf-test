package main

import (
	"github.com/gogf/gf/frame/g"
	_ "web_test/router"
)

func main() {
	s := g.Server()
	s.SetPort(8888)
	s.Run()
}
