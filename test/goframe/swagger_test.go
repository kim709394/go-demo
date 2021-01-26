package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"
	"github.com/kim709394/go-demo/goframe/pojo"
	"testing"
)

/*
@Author kim
@Description
@date 2021-1-22 18:43
*/

/**
1、go.mod添加swagger依赖:github.com/gogf/swagger v1.2.0，注意mod依赖库的代理设置
2、下载gf-cli工具:https://github.com/gogf/gf-cli
3、编写swagger注释,说明文档地址:https://github.com/swaggo/swag/blob/master/README_zh-CN.md
4、启动代码绑定swagger插件：
s := g.Server()
s.Plugin(&swagger.Swagger{})
5、项目根目录下执行gf命令；gf swagger --packed   注意如果gf-cli安装目录不和项目在同一个包下要配置环境变量
执行过程：寻找项目根目录下的main.go文件的启动代码，然后自动识别对应的实体结构体，然后在项目根目录下生成/swagger/swagger.json文件
6、输入地址查看swagger页面:   http://localhost:8082/swagger/
*/

func TestSwagger(t *testing.T) {

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
