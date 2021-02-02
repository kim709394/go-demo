package pojo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

/*
@Author kim
@Description
@date 2021-1-20 17:57
*/

type Phone struct {
	Id        int64       //主键
	IpAddress string      //ip地址
	Name      string      //名字
	PhoneNum  string      //电话号码
	Password  string      //密码
	CreatedAt *gtime.Time //时间
	DeletedAt *gtime.Time //删除时间
	UpdatedAt *gtime.Time //修改时间
	GId       int64       //组id
}

type Group struct {
	Id        int64       //主键
	Name      string      //名字
	CreatedAt *gtime.Time //时间
	DeletedAt *gtime.Time //删除时间
	UpdatedAt *gtime.Time //修改时间
}

func (g *Group) Hello(param string) string {

	return param

}

//统一返回值
type ResultVO struct {
	Code int
	Msg  string
	Data *g.Var `swaggerignore:"true"` //swagger忽略该字段
}

type Person struct {
	Id   int
	Name string
}

type PersonController struct {
}

//@Summary 人接口
//@Tags 人服务
//@Accept json
//@Produce  json
//@Param  entity body pojo.Person true "人员参数"
//@Success 201 {object} pojo.Person "返回结果"
//@Success 201 {integer} int64 "返回结果"
//@Failure 5001 {object} pojo.ResultVO
//@Header 200 {string} Token "token"
//@Router /swagger/test/post [POST]
func (pcr *PersonController) Call(r *ghttp.Request) {
	person := new(Person)
	r.GetStruct(person)
	r.Response.WriteJsonExit(ResultVO{201, "成功", g.NewVar(person)})
}
