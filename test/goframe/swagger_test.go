package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
3、编写swagger注释
4、启动代码绑定swagger插件：
s := g.Server()
s.Plugin(&swagger.Swagger{})
5、项目根目录下执行gf命令；gf swagger --packed   注意如果gf-cli安装目录不和项目在同一个包下要配置环境变量
*/

type Person struct {
	Id   int
	Name string
}

type PersonController struct {
}

//@summary 人接口
//@tags 人服务
//@produce  json
//@param
func (pcr *PersonController) call(r *ghttp.Request) {

}

func TestSwagger(t *testing.T) {

	server := g.Server()

}
