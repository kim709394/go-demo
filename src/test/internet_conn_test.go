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

//测试tcp服务端
func TestTcpServer(t *testing.T) {
	hello.TcpServer()
}

//测试tcp客户端
func TestTcpCustomer(t *testing.T) {
	hello.TcpCustomer()
}
