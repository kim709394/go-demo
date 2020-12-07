package test

import (
	"go-demo/hello"
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
