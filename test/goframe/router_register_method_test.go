package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

/*
@Author kim
@Description   web注册方式
@date 2020-12-8 16:50
*/

/*路由注册方式比较
在详细讲解每一种注册方式之前，先看看每种注册方式各自的优缺点，以便在不同的业务场景中选择更适合的注册方式。
如果暂时不理解这个表格没有关系，可以在了解完每一种注册方式之后再回过头来看，也许会更清晰。

注册方式	使用难度	安全系数	执行性能	内存消耗
回调函数注册	高	低	高	低
执行对象注册	中	中	中	中
控制器方式注册	低	高	低	高
比较指标说明：

使用难度：主要指对于执行流程以及数据管理维护的复杂度；
安全系数：主要指在异步多协程下的内部数据安全管理维护；
执行性能：执行性能，相对比较的结果；
内存消耗：内存消耗，相对比较的结果；
*/

//回调函数，在http_test.go已有示例

//对象注册方式
//控制器对象
type Controller struct {
}

//index方法，方法名要大写，否则框架无法在外部调用
func (c *Controller) Index(r *ghttp.Request) {
	r.Response.Write("index")
}

func (c *Controller) Add(r *ghttp.Request) {
	r.Response.Write("add")
}

func TestBindObj(t *testing.T) {
	s := g.Server()
	c := new(Controller)
	s.BindObject("/user", c)
	s.SetPort(8081)
	s.Run()
	/*控制器的每一个方法都是一个http的api，路由是/user拼接上控制器对象的方法名小写，
	  默认路由：/user=/user/index
	*/

	/*
		结果如下：
		SERVER  | DOMAIN  | ADDRESS | METHOD |    ROUTE    |                            HANDLER                            | MIDDLEWARE
		|---------|---------|---------|--------|-------------|---------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /user       | github.com/kim709394/go-demo/test/goframe.(*Controller).Index |
		|---------|---------|---------|--------|-------------|---------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /user/add   | github.com/kim709394/go-demo/test/goframe.(*Controller).Add   |
		|---------|---------|---------|--------|-------------|---------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /user/index | github.com/kim709394/go-demo/test/goframe.(*Controller).Index
	*/
}

//只绑定控制器对象中的几个方法
func TestBindObjOnly(t *testing.T) {
	s := g.Server()
	c := new(Controller)
	//第三个参数传入要暴露服务的方法名
	s.BindObject("/user", c, "Add")
	s.SetPort(8081)
	s.Run()
	/*
		SERVER  | DOMAIN  | ADDRESS | METHOD |   ROUTE   |                           HANDLER                           | MIDDLEWARE
		|---------|---------|---------|--------|-----------|-------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /user/add | github.com/kim709394/go-demo/test/goframe.(*Controller).Add |
		|---------|---------|---------|--------|-----------|-------------------------------------------------------------|------------|
	*/

}

//路由内置变量
type InnerVar struct {
}

func (i *InnerVar) Update(r *ghttp.Request) {
	r.Response.Write("update")
}

func TestInnerVar(t *testing.T) {
	s := g.Server()
	i := new(InnerVar)
	//内置变量路由规则，{.struct} = 对象名，小写驼峰转-，{.method} = 方法名小写
	s.BindObject("/{.struct}-{.method}", i)
	s.SetPort(8081)
	s.Run()
	/*
		* SERVER  | DOMAIN  | ADDRESS | METHOD |       ROUTE       |                           HANDLER                            | MIDDLEWARE
		|---------|---------|---------|--------|-------------------|--------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /inner-var-update | github.com/kim709394/go-demo/test/goframe.(*InnerVar).Update |
		|---------|---------|---------|--------|-------------------|--------------------------------------------------------------|------------|
	*/

}

//命名规则风格
type NameRule struct {
}

func (n *NameRule) List(r *ghttp.Request) {
	r.Response.Write("list")
}

func TestNameRule(t *testing.T) {

	/*
		URI_TYPE_DEFAULT  = 0 // （默认）全部转为小写，单词以'-'连接符号连接
		URI_TYPE_FULLNAME = 1 // 不处理名称，以原有名称构建成URI
		URI_TYPE_ALLLOWER = 2 // 仅转为小写，单词间不使用连接符号
		URI_TYPE_CAMEL    = 3 // 采用驼峰命名方式
		注意：需要在通过对象进行路由注册前进行该参数的设置，在路由注册后设置将不会生效，那么将使用默认规则
	*/
	s1 := g.Server("default")
	s2 := g.Server("full")
	s3 := g.Server("alllower")
	s4 := g.Server("camel")

	s1.SetNameToUriType(ghttp.URI_TYPE_DEFAULT)
	/*
		 SERVER  | DOMAIN  | ADDRESS | METHOD |      ROUTE      |                          HANDLER                           | MIDDLEWARE
		|---------|---------|---------|--------|-----------------|------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /name-rule-list | github.com/kim709394/go-demo/test/goframe.(*NameRule).List |
		|---------|---------|---------|--------|-----------------|------------------------------------------------------------|------------|
	*/

	s2.SetNameToUriType(ghttp.URI_TYPE_FULLNAME)
	/*
		SERVER | DOMAIN  | ADDRESS | METHOD |     ROUTE      |                          HANDLER                           | MIDDLEWARE
		|--------|---------|---------|--------|----------------|------------------------------------------------------------|------------|
		  full   | default | :8082   | ALL    | /NameRule-List | github.com/kim709394/go-demo/test/goframe.(*NameRule).List |
		|--------|---------|---------|--------|----------------|------------------------------------------------------------|------------|
	*/
	s3.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)
	/*
		 SERVER  | DOMAIN  | ADDRESS | METHOD |     ROUTE      |                          HANDLER                           | MIDDLEWARE
		|----------|---------|---------|--------|----------------|------------------------------------------------------------|------------|
		  alllower | default | :8083   | ALL    | /namerule-list | github.com/kim709394/go-demo/test/goframe.(*NameRule).List |
		|----------|---------|---------|--------|----------------|------------------------------------------------------------|------------|
	*/
	s4.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	/*
		  SERVER | DOMAIN  | ADDRESS | METHOD |     ROUTE      |                          HANDLER                           | MIDDLEWARE
		|--------|---------|---------|--------|----------------|------------------------------------------------------------|------------|
		  camel  | default | :8084   | ALL    | /nameRule-list | github.com/kim709394/go-demo/test/goframe.(*NameRule).List |
		|--------|---------|---------|--------|----------------|------------------------------------------------------------|------------|
	*/
	n := new(NameRule)
	s1.BindObject("/{.struct}-{.method}", n)
	s2.BindObject("/{.struct}-{.method}", n)
	s3.BindObject("/{.struct}-{.method}", n)
	s4.BindObject("/{.struct}-{.method}", n)
	s1.SetPort(8081)
	s2.SetPort(8082)
	s3.SetPort(8083)
	s4.SetPort(8084)
	s1.Start()
	s2.Start()
	s3.Start()
	s4.Start()
	g.Wait()
}

//restfull对象注册
type RestFull struct {
}

func (rest *RestFull) Get(r *ghttp.Request) {
	r.Response.Write("get")
}

func (rest *RestFull) Post(r *ghttp.Request) {
	r.Response.Write("post")
}

func (rest *RestFull) Put(r *ghttp.Request) {
	r.Response.Write("put")
}

func (rest *RestFull) Delete(r *ghttp.Request) {
	r.Response.Write("delete")
}

func TestRestFull(t *testing.T) {
	/*
		控制器对象的方法名必须首字母大写并且和http请求方法的名称一致，进行restful风格的url暴露服务，路由一样，请求方法不同
	*/
	s := g.Server()
	s.SetPort(8081)
	rest := new(RestFull)
	s.BindObjectRest("/rest", rest)
	s.Run()
	/*
		  SERVER  | DOMAIN  | ADDRESS | METHOD | ROUTE |                           HANDLER                            | MIDDLEWARE
		|---------|---------|---------|--------|-------|--------------------------------------------------------------|------------|
		  default | default | :8081   | DELETE | /rest | github.com/kim709394/go-demo/test/goframe.(*RestFull).Delete |
		|---------|---------|---------|--------|-------|--------------------------------------------------------------|------------|
		  default | default | :8081   | GET    | /rest | github.com/kim709394/go-demo/test/goframe.(*RestFull).Get    |
		|---------|---------|---------|--------|-------|--------------------------------------------------------------|------------|
		  default | default | :8081   | POST   | /rest | github.com/kim709394/go-demo/test/goframe.(*RestFull).Post   |
		|---------|---------|---------|--------|-------|--------------------------------------------------------------|------------|
		  default | default | :8081   | PUT    | /rest | github.com/kim709394/go-demo/test/goframe.(*RestFull).Put    |
		|---------|---------|---------|--------|-------|--------------------------------------------------------------|------------|
	*/
}

//初始化方法和结束回调方法
type LifeCycle struct {
}

//调用api之前调用的方法
func (l *LifeCycle) Init(r *ghttp.Request) {
	r.Response.Write("init")
}

//调用api之后调用的方法
func (l *LifeCycle) Shut(r *ghttp.Request) {
	r.Response.Write("shut")
}

func (l *LifeCycle) Call(r *ghttp.Request) {
	r.Response.Write("call")
}

func TestLifeCycle(t *testing.T) {
	//调用服务时，先调用Init()方法，后调用服务,call()方法，最后调用Shut()方法。
	s := g.Server()
	l := new(LifeCycle)
	s.BindObject("/user", l)
	s.SetPort(8081)
	s.Run()

}

type GroupRouter struct {
}

func (g *GroupRouter) List(r *ghttp.Request) {
	r.Response.Write("list")
}

func (g *GroupRouter) GetUser(r *ghttp.Request) {
	r.Response.Write("getUser")
}

func (g *GroupRouter) DeleteUser(r *ghttp.Request) {
	r.Response.Write("delete")
}

func (g *GroupRouter) UpdateUser(r *ghttp.Request) {
	r.Response.Write("updateUser")
}
func (g *GroupRouter) AddUser(r *ghttp.Request) {
	r.Response.Write("addUser")
}

//分组路由
func TestGroupRouter(t *testing.T) {

	s := g.Server()
	//路由前缀
	group := s.Group("/user")
	/*参数一：路由url，参数二：控制器对象，参数三：控制器对象里面的方法名。
	  绑定对应的路由url和方法作为暴露出去的服务
	*/
	gr := new(GroupRouter)
	group.ALL("/list", gr, "List")
	group.GET("/getUser", gr, "GetUser")
	group.POST("/addUser", gr, "AddUser")
	group.PUT("/updateUser", gr, "UpdateUser")
	group.DELETE("/deleteUser", gr, "DeleteUser")
	//参数一：路由url，参数二：匿名方法，对外暴露的服务，绑定在该路由
	group.GET("/getAnon", func(r *ghttp.Request) {
		r.Response.Write("getAnon")
	})
	s.SetPort(8081)
	s.Run()
	/*
		  SERVER  | DOMAIN  | ADDRESS | METHOD |      ROUTE       |                               HANDLER                               | MIDDLEWARE
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
		  default | default | :8081   | POST   | /user/addUser    | github.com/kim709394/go-demo/test/goframe.(*GroupRouter).AddUser    |
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
		  default | default | :8081   | DELETE | /user/deleteUser | github.com/kim709394/go-demo/test/goframe.(*GroupRouter).DeleteUser |
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
		  default | default | :8081   | GET    | /user/getAnon    | github.com/kim709394/go-demo/test/goframe.TestGroupRouter.func1     |
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
		  default | default | :8081   | GET    | /user/getUser    | github.com/kim709394/go-demo/test/goframe.(*GroupRouter).GetUser    |
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
		  default | default | :8081   | ALL    | /user/list       | github.com/kim709394/go-demo/test/goframe.(*GroupRouter).List       |
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
		  default | default | :8081   | PUT    | /user/updateUser | github.com/kim709394/go-demo/test/goframe.(*GroupRouter).UpdateUser |
		|---------|---------|---------|--------|------------------|---------------------------------------------------------------------|------------|
	*/
}
