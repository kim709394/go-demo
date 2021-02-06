package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/kim709394/go-demo/goframe/pojo"
	"testing"
)

/*
@Author kim
@Description   静态文件服务模板
@date 2021-2-2 17:44
*/

/*#静态模板配置
[viewer]
Paths   = ["D:\\go\\goworks\\go-demo\\goframe\\static"]    #前端模板文件放置路径
DefaultFile ="index.html"         #默认解析的模板引擎文件:index.html
Delimiters  =  ["${", "}"]          # 模板引擎变量分隔符号。默认为 ["{{", "}}"]
AutoEncode  = false                 # 是否默认对变量内容进行XSS编码。默认为false*/

func TestStatic(t *testing.T) {

	s := g.Server()
	s.BindHandler("/static", func(r *ghttp.Request) {
		group := new(pojo.Group)
		group.Name = "kim"
		group.Id = 1
		group.CreatedAt = gtime.Now()
		persons := g.Slice{&pojo.Person{2, "小米"}, &pojo.Person{3, "华为"}, &pojo.Person{4, "三星"}}

		r.Response.WriteTpl("temp.html", g.Map{
			"val":    "普通变量",
			"simple": group,
			"slice":  persons,
			"myTemp": "myTemp.html",
		})
	})
	s.SetPort(8081)
	s.Run()

	/*
		and: {{and .X .Y .Z}}会逐一判断每个参数，将返回第一个为空的参数，否则就返回最后一个非空参数
		index: index支持map, slice, array, string，读取指定类型对应下标的值。{{index .Maps "name"}}
		len :  {{printf "The content length is %d" (.Content|len)}}   返回对应类型的长度，支持类型：map, slice, array, string, chan。
		not: 返回输入参数的否定值。 {{if not .Var}} {{end}}
		or: {{or .X .Y .Z}}  or会逐一判断每个参数，将返回第一个非空的参数，否则就返回最后一个参数。
	*/

}
