package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"
	"github.com/kim709394/go-demo/goframe/pojo"
)

/*
@Author kim
@Description
@date 2021-1-25 9:55
*/

func main() {
	pcr := new(pojo.PersonController)
	s := g.Server()
	s.Group("/swagger", func(group *ghttp.RouterGroup) {
		group.Group("/test", func(group *ghttp.RouterGroup) {
			group.POST("/post", pcr, "Call")
		})
	})
	s.Plugin(&swagger.Swagger{})
	s.SetPort(8081)
	s.Run()
}
