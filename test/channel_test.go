package test

import (
	"go-demo/hello"
	"testing"
)

/*
@Author kim
@Description
@date 2020-11-25 12:53
*/

//通道收发数据
func TestCn(t *testing.T) {
	hello.SendAndReceive()
}

//循环接收通道数据
func TestLoopRecv(t *testing.T) {
	hello.LoopRecv()
}

//单向通道
func TestOneWayCn(t *testing.T) {
	hello.OneWayCn()
}

//关闭通道
func TestCloseCn(t *testing.T) {
	hello.CloseCn()
}

//带缓冲通道
func TestCnCh(t *testing.T) {
	hello.ChanWithCache()
}

//select语句
func TestSele(t *testing.T) {
	hello.Sele()
}

//timer延时函数
func TestTimer(t *testing.T) {
	hello.Timer()
}

//ticker循环延时函数
func TestTicker(t *testing.T) {
	hello.Ticker()
}
