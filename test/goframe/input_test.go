package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"testing"
)

/*
@Author kim
@Description   请求参数
@date 2020-12-11 8:59
*/

//参数获取方法
func TestGetParam(t *testing.T) {

	s := g.Server()
	s.BindHandler("/get/param", func(r *ghttp.Request) {
		res := r.Response
		//不分提交方式获取参数,参数优先级：Router < Query < Body < Form < Custom
		res.Write("不分提交方式获取参数:", r.Get("p"))
		//等同
		res.Write("等同:", r.GetRequest("p"))
		//获取get请求方式的拼接参数
		res.Write("获取get请求方式的拼接参数", r.GetQuery("p"))
		//获取表单方式传过来的参数
		res.Write("获取表单方式传过来的参数:", r.GetForm("form"))
		//获取原始数据,该数据是客户端写入到body中的原始数据，与HTTP Method无关，例如客户端提交JSON/XML数据格式时可以通过该方法获取原始的提交数据
		res.Write("获取原始数据", string(r.GetBody()))
		//获取自动转换的参数类型
		res.Write("获取自动转换的参数类型,string:", r.GetString("str"))
		res.Write("获取自动转换的参数类型,int:", r.GetInt64("int"))
		res.Write("获取自动转换的参数类型:float64", r.GetFloat64("flo"))
		//getVar方式进行参数类型转换
		res.Write("getVar方式进行参数类型转换:", r.GetVar("gv").Int8())
		//获取数组参数,同名参数只能获取最后一个参数，前面的会被覆盖
		res.Write("获取数组参数:", r.GetFormArray("arr"))
	})
	s.BindHandler("/json", func(r *ghttp.Request) {
		res := r.Response
		//将json字符串参数自动转化为匹配的结构体
		//json数组
		/*param :=make([]Param,1)
		err1 := r.Parse(&param)
		if err1 != nil {
			fmt.Println(err1)
		}*/
		//单个json对象
		var single *Param
		err2 := r.Parse(&single)
		//服务端数据校验获取错误信息
		if err2 != nil {
			if v, ok := err2.(*gvalid.Error); ok {
				res.WriteJsonExit(ResultVO{5001, v.FirstString(), ""})
			}
			//向客户端写完json字符串后结束方法
			res.WriteJsonExit(ResultVO{5001, err2.Error(), ""})
		}
		//res.Write("将json字符串参数自动转化为匹配的结构体",param)
		res.Write("单个json对象:", single)
	})
	s.SetPort(8081)
	s.Run()

}

/*自动将json转化为结构体参数,要求属性是公开的，首字母大写，默认忽略大小写和" - "、 " _ " 进行匹配
  如果要修改匹配json字符串的属性名，可以用tag标签，标签key可以是p//param/params
  如果无法匹配，则忽略该属性
  服务端数据校验:tag标签添加即可，标签key值是v
*/
type Param struct {
	Id    int     `p:"id"`
	Name  string  `p:"name" v:"required|length:6,30#必须输入名字|名字长度要在6到30位之间"`
	Age   int     `p:"age"`
	Score float64 `p:"score"`
}

//自定义变量与上下文变量
func TestCustmer(t *testing.T) {

	s := g.Server()
	s.BindHandler("/customer", func(r *ghttp.Request) {
		//设置自定义参数
		r.SetParam("customer", "kim")
		//获取自定义参数
		customer := r.GetParamVar("customer").String()
		//上下文变量，推荐

		r.SetCtxVar("ctx", "kim of ctx")
		ctx := r.GetCtxVar("ctx")
		context := r.GetCtx()
		r.Response.WriteExit(customer, ctx, context.Value("ctx"))
	})
	s.SetPort(8081)
	s.Run()

}
