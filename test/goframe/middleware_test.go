package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

/*
@Author kim
@Description  中间件定义，类似java的过滤器，springmvc的拦截器
@date 2020-12-9 16:09
*/

//中间类型、前置中间件、后置中间件
//前置中间件
func PreMiddleware(r *ghttp.Request) {
	//前置业务处理
	r.Middleware.Next()
}

func AfterMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	//后置业务处理
}

//跨域中间件
func MiddlewareCors(r *ghttp.Request) {
	r.Response.CORSDefault()
	/*
		 个性化设置
		//设置允许的请求域名
		corsOptions := r.Response.DefaultCORSOptions()
		corsOptions.AllowDomain = []string{"www.baidu.com", "www.jd.com"}
		r.Response.CORS(corsOptions)
		r.Middleware.Next()

	*/
	r.Middleware.Next()
}

//鉴权中间件
func AuthMiddleware(r *ghttp.Request) {
	//配置不需要鉴权的路由url
	anon := make([]string, 5)
	anon = append(anon, "/goframe/goods/add", "/goframe/user/list")
	for _, uri := range anon {
		if uri == r.RequestURI {
			r.Middleware.Next()
			/*
				Exit: 仅退出当前执行的逻辑方法，不退出后续的请求流程，可用于替代return。
				ExitAll: 强行中断当前执行流程，当前执行方法的后续逻辑以及后续所有的逻辑方法将不再执行，常用于权限控制
			*/
			//return
			r.Exit()
		}
	}
	//进行token鉴权
	token := r.Header.Get("token")
	if token == "" {
		token, _ = r.Get("token").(string)
	}
	if token == "709394" {
		r.Middleware.Next()
	} else {
		result := ResultVO{4001, "你无权限访问", ""}
		r.Response.WriteJson(result)
	}

}

//统一异常处理处理中间件
func GlobalError(r *ghttp.Request) {
	r.Middleware.Next()
	//获取错误信息
	err := r.GetError()
	if err != nil {
		r.Context()
		r.Response.ClearBuffer()
		r.Response.WriteJson(ResultVO{5001, "服务端错误", err.Error()})
	}
	/*if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.WriteJson(ResultVO{5001, "服务端错误", err.Error()})
	}*/

}

//全局中间件以及分组中间件
func TestGroupMiddleware(t *testing.T) {
	/*
		全局中间件
		全局中间件仅对动态请求拦截有效，无法拦截静态文件请求。
		由于全局中间件也是通过路由规则执行，那么也会存在执行优先级：
		首先，由于全局中间件是基于模糊路由匹配，因此当同一个路由匹配到多个中间件时，会按照路由的深度优先规则执行，具体请查看路由章节；
		其次，同一个路由规则下，会按照中间件的注册先后顺序执行，中间件的注册方法也支持同时按照先后顺序注册多个中间件；
		最后，为避免优先级混淆和后续管理，建议将所有中间件放到同一个地方进行先后顺序注册来控制执行优先级；
	*/
	s := g.Server()
	//全局中间件，模糊匹配路由
	s.BindMiddleware("/goframe/user/*", func(r *ghttp.Request) {
		r.Response.Write("用户模块全局中间件")
		r.Middleware.Next()
	})
	//全局中间件，默认匹配/*路由
	s.BindMiddlewareDefault(func(r *ghttp.Request) {
		r.Response.Write("默认全局中间件")
		r.Middleware.Next()
	})
	//默认全局中间别名
	s.Use(func(r *ghttp.Request) {
		r.Response.Write("默认全局中间件别名")
		r.Middleware.Next()
	})
	//跨域中间件
	s.Use(MiddlewareCors)
	//鉴权中间件
	s.Use(AuthMiddleware)
	//错误处理器
	s.Use(GlobalError)
	//实例化控制器
	u := new(UserController)
	g := new(GoodsController)

	//层级分组路由，进行项目名到模块名再到接口名的路由层级注册
	s.Group("/goframe", func(group *ghttp.RouterGroup) {
		//user模块接口
		group.Group("/user", func(group *ghttp.RouterGroup) {
			//分层内设置中间件,只对该层级的所有路由生效
			group.Middleware(func(r *ghttp.Request) {
				r.Response.Write("用户模块内中间件")
				r.Middleware.Next()
			})
			group.ALL("/list/:id", u, "List")
			group.GET("/get", u, "GetUser")
			group.DELETE("/delete", u, "DeleteUser")
			group.PUT("/update", u, "UpdateUser")
			group.POST("/add", u, "AddUser")
		})
		//goods模块接口
		group.Group("/goods", func(group *ghttp.RouterGroup) {
			group.ALL("/list", g, "List")
			group.GET("/get", g, "GetGoods")
			group.DELETE("/delete", g, "DeleteGoods")
			group.PUT("/update", g, "UpdateGoods")
			group.POST("/add", g, "AddGoods")
		})
	})
	s.SetPort(8081)
	s.Run()
}

//user模块controller
type UserController struct {
}

func (u *UserController) List(r *ghttp.Request) {
	//panic(&hello.MyError{Code: 1, Msg: "报错"})
	url := r.URL
	g.Dump(r.Router.Uri)
	g.Dump("url", url)
	r.Response.Write("listUser")
}

func (u *UserController) GetUser(r *ghttp.Request) {
	r.Response.Write("getUser")
}

func (u *UserController) DeleteUser(r *ghttp.Request) {
	r.Response.Write("deleteUser")
}

func (u *UserController) UpdateUser(r *ghttp.Request) {
	r.Response.Write("updateUser")
}
func (u *UserController) AddUser(r *ghttp.Request) {
	r.Response.Write("addUser")
}

//商品模块controller
type GoodsController struct {
}

func (g *GoodsController) List(r *ghttp.Request) {
	if r.Get("err") != nil {
		panic("报错了")
	}
	r.Response.Write(ResultVO{0, "成功返回", "listGoods"})
}

func (g *GoodsController) GetGoods(r *ghttp.Request) {
	r.Response.Write("getGoods")
}

func (g *GoodsController) DeleteGoods(r *ghttp.Request) {
	r.Response.Write("deleteGoods")
}

func (g *GoodsController) UpdateGoods(r *ghttp.Request) {
	r.Response.Write("updateGoods")
}
func (g *GoodsController) AddGoods(r *ghttp.Request) {
	r.Response.Write("addGoods")
}

//统一返回值
type ResultVO struct {
	Code int
	Msg  string
	Data interface{}
}
