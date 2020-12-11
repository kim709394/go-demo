package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

/*
@Author kim
@Description   cookie session
@date 2020-12-11 18:01
*/

//cookie
func TestCookie(t *testing.T) {

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		//设置一个cookie
		r.Cookie.Set("coo", "myCoo")
		//获取cookie信息
		coo := r.Cookie.Get("coo")
		r.Cookie.Set("coo1", "myCoo")
		r.Response.Write(coo)
	})
	s.SetPort(8081)
	s.Run()
}

//session
func TestSession(t *testing.T) {

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {

	})

}
