package hello

import (
	"fmt"
)

/*
@Author kim
@Description   接口定义和使用
@date 2020-10-27 11:41
*/

//接口定义，go语言的接口实现都是隐式的，不需要显式写implememts关键字，只要符合条件就是接口的实现
type MyInterface interface {
	Print(s string) string
}

//接口实现结构体，相当于定义接口实现类
type str struct {
}

type str2 struct{}

func (*str2) Print(s string) string {
	fmt.Print("实现结构体str2:", s)
	return s
}

//(*str)表示该方法属于str结构体的函数，类似java的类中的方法，实现接口中的方法
func (*str) Print(s string) string {
	fmt.Println("实现结构体str:", s)
	return s
}

//结构体内嵌，内嵌的结构体实现了接口的方法，外层结构体可以“继承”下来
type Men interface {
	Eat()

	Run()
}

type girl struct {
}

type women struct {
	girl
}

func (*girl) Eat() {
	fmt.Println("吃饭")
}

func (*women) Run() {
	fmt.Println("跑步")
}

//接口实现
func Impl() {
	f := new(str)
	var myInter MyInterface
	f.Print("ss")
	myInter = f
	myInter.Print("go语言接口定义")

	wom := new(women)
	var men Men = wom
	men.Eat()
	men.Run()

}

//接口类型转换
func InterConvert() {
	var myIn interface{}
	myIn = new(str)
	//将对象强制转换为Men接口，将失败
	p, f1 := myIn.(*MyInterface)
	fmt.Println(p, f1)
	//接口强制转换为其实现类***
	var myIn2 MyInterface

	myIn2 = new(str)
	s, f2 := myIn2.(*str)
	fmt.Println(s, f2)
}
