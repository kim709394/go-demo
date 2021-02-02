package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

/*
@Author kim
@Description
@date 2020-12-8 10:39
*/

//简单的http服务
func TestHttp(t *testing.T) {
	//开启一个server实例
	server := g.Server()

	//绑定路由，传入执行处理的函数
	server.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello goframe")
	})
	//设置端口号，不设置默认监听80
	server.SetPort(8081)
	//设置多端口监听
	//server.SetPort(8081,8082,8083)
	//静态资源渲染
	//设置允许访问静态资源，默认是false
	server.SetIndexFolder(true)
	//设置静态资源路径,访问/index.html即是访问/goframe/static/index.html
	server.SetServerRoot("/goframe/static/")
	//静态路由重写,访问/index=访问/index.html
	server.SetRewrite("/index", "/index.html")
	//map方式设置多个静态路由重写
	server.SetRewriteMap(g.MapStrStr{
		"/index1": "/index.html",
		"index2":  "index2.html",
	})
	//跑起来
	server.Run()
}

//多实例
func TestMultiServer(t *testing.T) {
	//多实例
	s1 := g.Server("s1")
	s1.BindHandler("/s1", func(r *ghttp.Request) {
		r.Response.Write("hello s1")
	})
	s1.SetPort(8081)
	s1.Start()

	s2 := g.Server("s2")
	s2.BindHandler("/s2", func(r *ghttp.Request) {
		r.Response.Write("hello s2")
	})
	s2.SetPort(8082)
	s2.Start()
	g.Wait()
}

//域名绑定
func TestDomain(t *testing.T) {

	s := g.Server()
	s.Domain("www.baidu.com").BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello baidu")
	})
	s.Domain("www.jd.com").BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hell jd")
	})
	s.SetPort(8081)
	s.Run()
}

//请求方式
func TestHttpMethod(t *testing.T) {

	s := g.Server()
	//get请求
	s.BindHandler("GET:/", func(r *ghttp.Request) {
		r.Response.Write("hello get")
	})
	//post请求
	s.BindHandler("POST:/", func(r *ghttp.Request) {
		r.Response.Write("hello post")
	})
	//put请求
	s.BindHandler("PUT:/", func(r *ghttp.Request) {
		r.Response.Write("hello put")
	})
	//delete请求
	s.BindHandler("DELETE:/", func(r *ghttp.Request) {
		r.Response.Write("hello delete")
	})
	s.SetPort(8081)
	s.Run()

}

//动态路由规则
func TestDynamicRouter(t *testing.T) {
	s := g.Server()
	//精准匹配规则,100%匹配
	s.BindHandler("/user/login", func(r *ghttp.Request) {
		r.Response.Write("精准匹配")
	})
	/*命名匹配规则
	*使用:name方式进行匹配(name为自定义的匹配名称)，对URI指定层级的参数进行命名匹配（类似正则([^/]+)，该URI层级必须有值），
	*对应匹配参数会被解析为Router参数并传递给注册的服务接口使用
	 */
	/*
		*匹配示例1：
		rule: /user/:user

		/user/john                match
		/user/you                 match
		/user/john/profile        no match
		/user/                    no match

		匹配示例2：
		rule: /:name/action
		/john/name                no match
		/john/action              match
		/smith/info               no match
		/smith/info/age           no match
		/smith/action             match
		匹配示例3：

		rule: /:name/:action

		/john/name                match
		/john/info                match
		/smith/info               match
		/smith/info/age           no match
		/smith/action/del         no match
	*/
	s.BindHandler("/user/:myName", func(r *ghttp.Request) {
		//获取匹配参数
		myName := r.Get("myName")
		r.Response.Write("命名匹配:", myName)
	})

	/*模糊匹配,使用*any方式进行匹配(any为自定义的匹配名称)，对URI指定位置之后的参数进行模糊匹配（类似正则(.*)，该URI层级可以为空），
	*并将匹配参数解析为Router参数并传递给注册的服务接口使用。
	 */
	/*
		*匹配示例1：

		rule: /src/*path

		/src/                     match
		/src/somefile.go          match
		/src/subdir/somefile.go   match
		/user/                    no match
		/user/john                no match
		匹配示例2：

		rule: /src/*path/:action

		/src/                     no match
		/src/somefile.go          match
		/src/somefile.go/del      match
		/src/subdir/file.go/del   match
		匹配示例3：

		rule: /src/*path/show

		/src/                     no match
		/src/somefile.go          no match
		/src/somefile.go/del      no match
		/src/somefile.go/show     match
		/src/subdir/file.go/show  match
		/src/show                 match
	*/
	s.BindHandler("/user/*path", func(r *ghttp.Request) {
		//获取匹配路由参数
		path := r.Get("path")
		r.Response.Write("模糊匹配:", path)
	})

	/*字段匹配
	使用{field}方式进行匹配(field为自定义的匹配名称)，可对URI任意位置的参数进行截取匹配（类似正则([\w\.\-]+)，
	该URI层级必须有值，并且可以在同一层级进行多个字段匹配），并将匹配参数解析为Router参数并传递给注册的服务接口使用。
	*/
	/*
		匹配示例1：

		rule: /order/list/{page}.php

		/order/list/1.php          match
		/order/list/666.php        match
		/order/list/2.php5         no match
		/order/list/1              no match
		/order/list                no match
		匹配示例2：

		rule: /db-{table}/{id}

		/db-user/1                     match
		/db-user/2                     match
		/db/user/1                     no match
		/db-order/100                  match
		/database-order/100            no match
		匹配示例3：

		rule: /{obj}-{act}/*param

		/user-delete/10                match
		/order-update/20               match
		/log-list                      match
		/log/list/1                    no match
		/comment/delete/10             no match
	*/
	s.BindHandler("/user/{any}.go", func(r *ghttp.Request) {
		any := r.Get("any")
		r.Response.Write("字段匹配:", any)
	})

	s.SetPort(8081)
	s.Run()
	/*
		路由优先级控制
		优先级控制按照深度优先策略，最主要的几点因素：

		层级越深的规则优先级越高；
		同一层级下，精准匹配优先级高于模糊匹配；
		同一层级下，模糊匹配优先级：字段匹配 > 命名匹配 > 模糊匹配；
	*/
}
