package hello

import (
	"errors"
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
type Str struct {
}

type Str2 struct{}

func (*Str2) Print(s string) string {
	fmt.Print("实现结构体str2:", s)
	return s
}

//(*str)表示该方法属于str结构体的函数，类似java的类中的方法，实现接口中的方法
func (*Str) Print(s string) string {
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

//可以将结构体实例传参进来
func (g *girl) Eat() {
	fmt.Println("吃饭")
}

func (*women) Run() {
	fmt.Println("跑步")
}

//接口实现
func Impl() {
	f := new(Str)
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
	myIn = new(Str)
	//将对象强制转换为Men接口，将失败
	p, f1 := myIn.(*MyInterface)
	fmt.Println(p, f1)
	//接口强制转换为其实现类***
	var myIn2 MyInterface

	myIn2 = new(Str)
	s, f2 := myIn2.(*Str)
	fmt.Println(s, f2)
}

//空接口
func NullInterface() {
	//声明一个空接口，类似于java的Object，会比其他类型的变量导入内存慢
	var a interface{}
	a = 1
	fmt.Println(a)
	//类型断言进行转换
	var b interface{}
	b = 1
	c := b.(int)
	fmt.Println(c)
	//空接口值比较,先比较类型，再比较值，类型不同则为false，类型相同再对值进行比较
	var d interface{} = 2
	var f interface{} = "hi"
	fmt.Println(d == f)
	var g interface{} = 3
	var h interface{} = 3
	fmt.Println(g == h)
	var i interface{} = 5
	var j interface{} = 6
	fmt.Println(i == j)

	/*类型的可比较性
	类  型	说  明
	map	宕机错误，不可比较
	切片（[]T）	宕机错误，不可比较
	通道（channel）	可比较，必须由同一个 make 生成，也就是同一个通道才会是 true，否则为 false
	数组（[容量]T）	可比较，编译期知道两个数组是否一致
	结构体	可比较，可以逐个比较结构体的值
	函数	可比较*/
}

//使用switch判断空接口的类型
func SwitInter(obj interface{}) {

	switch obj.(type) {
	case Str:
		fmt.Println("str类型")
	case MyInterface:
		fmt.Println("MyInterface类型")
	case Str2:
		fmt.Println("str2类型")
	default:
		fmt.Println("无类型")
	}

}

//错误接口
func Err(err interface{}) error {
	if err == nil {
		//实例化一个异常
		error := errors.New("抛出异常")
		return error
	} else {
		//返回自定义异常
		return &myErr{404, "404异常"}
	}

}

//自定义ERROR结构体类型
type myErr struct {
	ErrorCode int
	Msg       string
}

//重写Error方法，实现error接口
func (e *myErr) Error() string {
	return fmt.Sprint("错误码:", e.ErrorCode, "错误信息:", e.Msg)
}
