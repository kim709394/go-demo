package hello

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/kim709394/go-demo/hello"
	"testing"
)

/*
@Author kim
@Description
@date 2020-11-27 9:57
*/

//获取变量的类型信息
func TestTypeOf(t *testing.T) {
	hello.TypeOf()
}

//获取结构体内部信息
func TestGetStruct(t *testing.T) {
	hello.GetStructMsg()
}

//获取结构体标签
func TestGetTag(t *testing.T) {
	hello.GetStructTag()
}

//排除空字段值
func TestIgnoreNull(t *testing.T) {
	myDog := hello.Dog{"旺财", 2, 'y', gtime.Now(), hello.MyStu{1, 2}}
	hello.IgnoreStructNull(myDog)
}

//基本类型初始值
func TestBasicInitValue(t *testing.T) {
	hello.GetBasicInitValue()
}
