package test

import (
	"hello"
	"testing"
)

/*
@Author kim
@Description  网络编程测试
@date 2020-11-30 14:02
*/

//测试get服务接口
func TestGet(t *testing.T) {
	hello.Get()
}
