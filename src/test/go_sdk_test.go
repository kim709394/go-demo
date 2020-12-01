package test

import (
	"hello"
	"testing"
)

/*
@Author kim
@Description
@date 2020-12-1 8:57
*/

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
