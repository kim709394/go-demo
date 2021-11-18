package hello

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/kim709394/go-demo/hello"
	"testing"
	"time"
)

/*
@Author kim
@Description
@date 2020-12-1 8:57
*/

//测试字符串拼接
func TestStr(t *testing.T) {
	hello.Strs()
}

//测试json序列化
func TestJsonSerialize(t *testing.T) {

	hello.JsonSerialize()

}

//测试序列化map
func TestJsonMap(t *testing.T) {
	hello.JsonSerializeMap()
}

//测试反序列化结构体
func TestUnSerialize(t *testing.T) {
	hello.JsonUnSerialize()
}

//测试反序列化到map
func TestUnSerializeMap(t *testing.T) {
	hello.JsonUnSerializeMap()
}

//获取本机ip地址
func TestGetIp(t *testing.T) {
	hello.GetIp()
}

//测试字符串占位符拼接
func TestSprintf(t *testing.T) {
	hello.Sprintf()
}

//测试panic
func TestPanic(t *testing.T) {
	fmt.Println("主协程开始")
	timer := time.NewTimer(5 * time.Second)
	for i := 0; i < 10; i++ {
		go func(j int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			panic("故意panic" + gconv.String(j))
		}(i)

	}

	<-timer.C
	fmt.Println("主协程结束")

}
