package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gsession"
	"testing"
	"time"
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

//session，存储方式有三种：文件存储、内存存储、redis存储

//文件存储,默认存储方式
func TestFileSession(t *testing.T) {

	s := g.Server()
	//设置session过期时间为1分钟
	s.SetConfigWithMap(g.Map{
		"SessionMaxAge": time.Minute,
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/set", func(r *ghttp.Request) {
			//往session设置用户名
			r.Session.Set("user", "kim")
			r.Response.Write("set session over")
		})
		group.ALL("/get", func(r *ghttp.Request) {
			//往session里面获取用户名
			user := r.Session.GetString("user")
			r.Response.Write(user)
		})
		//删除session
		group.ALL("/del", func(r *ghttp.Request) {
			r.Session.Clear()
			r.Response.Write("ok")
		})
	})
	s.SetPort(8081)
	s.Run()

}

//内存存储
func TestCacheSession(t *testing.T) {

	s := g.Server()

	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,                 //设置session过期时间为1分钟
		"SessionStorage": gsession.NewStorageMemory(), //设置存储方式
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/set", func(r *ghttp.Request) {
			//往session设置用户名
			r.Session.Set("user", "kim")
			r.Response.Write("set session over")
		})
		group.ALL("/get", func(r *ghttp.Request) {
			//往session里面获取用户名
			user := r.Session.GetString("user")
			r.Response.Write(user)
		})
		//删除session
		group.ALL("/del", func(r *ghttp.Request) {
			r.Session.Clear()
			r.Response.Write("ok")
		})
	})
	s.SetPort(8081)
	s.Run()

}

//redis存储
func TestRedisSession(t *testing.T) {
	s := g.Server()

	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,                         //设置session过期时间为1分钟
		"SessionStorage": gsession.NewStorageRedis(g.Redis()), //设置存储方式
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/set", func(r *ghttp.Request) {
			//往session设置用户名
			r.Session.Set("user", "kim")
			r.Response.Write("set session over")
		})
		group.ALL("/get", func(r *ghttp.Request) {
			//往session里面获取用户名
			user := r.Session.GetString("user")
			r.Response.Write(user)
		})
		//删除session
		group.ALL("/del", func(r *ghttp.Request) {
			r.Session.Clear()
			r.Response.Write("ok")
		})
	})
	s.SetPort(8081)
	s.Run()
}
