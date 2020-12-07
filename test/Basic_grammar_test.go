package test

import (
	"fmt"
	"go-demo/hello"
	"testing"
	"time"
)

/*
@Author kim
@Description  基础语法测试类
@date 2020-11-23 13:57
*/

//测试异常和捕捉
func TestException(t *testing.T) {
	hello.Exception()
}

//测试全局变量
func TestGlobalVal(t *testing.T) {
	fmt.Println(hello.GlobalVal, "kk")
}

//测试复数
func TestComplex(t *testing.T) {
	hello.Complex()
}

//测试字符串
func TestString(t *testing.T) {
	hello.String()
}

//测试浮点型数据类型
func TestFloat(t *testing.T) {
	hello.Float()
}

//测试指针
func TestPtr(t *testing.T) {
	hello.Ptr()
}

//测试循环
func TestLoop(t *testing.T) {
	hello.Loop()
}

//测试数组
func TestArray(t *testing.T) {
	hello.Array()
}

//测试映射
func TestMap(t *testing.T) {
	hello.Map()
}

//测试列表
func TestList(t *testing.T) {
	hello.List()
}

//测试选择结构
func TestSwh(t *testing.T) {
	hello.Swh()
}

//测试退出循环
func TestBreak(t *testing.T) {
	hello.Brek()
}

//测试循环继续下一轮
func TestContu(t *testing.T) {
	hello.Contu()
}

//测试函数返回值带变量名
func TestF(t *testing.T) {
	fmt.Println(hello.F())
}

//测试可变参数
func TestEnableParam(t *testing.T) {
	hello.EnableParam()
}

//测试延迟语句
func TestDer(t *testing.T) {
	hello.Der()
}

//测试实例化结构体
func TestInsStruct(t *testing.T) {
	hello.InsStruct()
}

//测试接口实现
func TestImpl(t *testing.T) {
	hello.Impl()
}

//测试类型断言
func TetsAssertion(t *testing.T) {
	hello.Assertion()
}

//测试获取当前时间
func TestCurrenTime(t *testing.T) {
	time.Sleep(time.Duration(10) * time.Second)

}

//测试获取键盘输入信息
func TestScan(t *testing.T) {
	hello.Scan()
}

//测试删除同步map的元素
func TestSyncMap(t *testing.T) {
	hello.DeleteSyncMap()
}
