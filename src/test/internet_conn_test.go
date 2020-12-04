package test

import (
	"fmt"
	"hello"
	"net"
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

//测试get http客户端
func TestGetClient(t *testing.T) {
	hello.GetClient()
}

//测试tcp服务端
func TestTcpServer(t *testing.T) {
	hello.TcpServer()
}

//测试tcp客户端
func TestTcpCustomer(t *testing.T) {
	hello.TcpCustomer()
}

//测试tcp文件传输服务端
func TestTcpFileServer(t *testing.T) {
	hello.TcpFileServer()
}

//测试tcp文件传输客户端
func TestTcpFileCustomer(t *testing.T) {
	hello.TcpFileCustomer()
}

//测试tcp聊天室服务端
func TestTcpChatServer(t *testing.T) {
	hello.TcpChatServer()
}

//测试转换ip
func TestIpConvert(t *testing.T) {
	hello.IpConvert()
	var a *net.UDPAddr
	fmt.Println(a)
}

//测试udp服务端
func TestUdpServer(t *testing.T) {
	hello.UdpServer()
}

//测试udp客户端
func TestUdpCustomer(t *testing.T) {
	hello.UdpCustomer()
}
