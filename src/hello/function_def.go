package hello

/*
@Author kim
@Description  函数的定义和规则
@date 2020-10-21 15:42
*/
import (
	"fmt"
	"os"
	"sync"
)

//返回值带有变量名的函数，可以在函数体直接对返回值赋值，然后直接返回
func F() (a, b int) {
	fmt.Println("------------------------返回值带有变量名的函数，可以在函数体直接对返回值赋值，然后直接返回-------------------------")
	a, b = 1, 2
	return
}

func fx() {
	fmt.Println("函数被调用")
}

//函数变量，函数也可以作为一种类型，进行赋值和传递
func funcVariable() {
	fmt.Println("------------------------函数变量，函数也可以作为一种类型，进行赋值和传递-------------------------")
	//函数类型变量的声明
	var f func() (a, b int)
	//将函数赋值给函数变量,函数变量的函数参数和返回值要和赋值的函数一致，将函数名赋值等同于赋值整个函数
	f = F
	//调用函数变量
	f()
	var x func()
	x = fx
	x()
}

//匿名函数
func FooFunc() {
	fmt.Println("------------------------匿名函数-------------------------")
	//将匿名函数赋值给变量
	f := func(a int) {
		fmt.Println(a)
	}
	//调用函数
	f(1)

	//在定义是调用匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(10, 20)

	//回调函数
	func(l []int, f func(k, v int)) {

		for k0, v0 := range l {
			f(k0, v0)
		}

	}([]int{1, 2, 3, 4}, func(k, v int) {
		fmt.Println(k, v)
	})

}

//可变参数
func EnableParam() {
	fmt.Println("------------------------可变参数-------------------------")
	//固定类型的可变参数
	func(p ...int) {
		for k, v := range p {
			fmt.Println(k, v)
		}
	}(1, 2, 3)

	//任意类型的可变参数,interface{}代表任意类型，类似java的Object
	func(params ...interface{}) {
		for k, v := range params {
			fmt.Println(k, v)
		}
	}(1, "string", "int")

	//用switch判断获取不同的参数类型
	f := func(params ...interface{}) {
		for _, v := range params {

			//对值进行类型断言
			switch v.(type) {
			case bool:
				fmt.Println("布尔值")
			case int:
				fmt.Println("整型")
			case string:
				fmt.Println("字符串")
			}
		}
	}
	f(1, false, "foo")

	//可变参数传递
	func(paramList ...interface{}) {

		f(paramList)
	}(1, 2, 3)
}

//defer延迟语句,被defer修饰的代码将被延迟执行，最先修饰的最后执行，最后执行的最先修饰
func Der() {

	fmt.Println("------------------------defer延迟语句-------------------------")

	//类似于java的finally，在整个函数执行完最后才执行
	fmt.Println("defer begin")

	defer fmt.Println("首先被修饰，最后执行")

	defer fmt.Println("第二个被修饰，倒数第二执行")

	defer fmt.Println("最后被修饰，首先执行")

	fmt.Println("defer end")

	//延迟并发解锁
	var lock sync.Mutex //初始化互斥线程锁
	func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("并发操作")

	}()

	//用于文件最后释放资源

	file, err := os.Open("")
	defer file.Close()
	fmt.Println(file)
	fmt.Println(err)
}

//init函数.初始化函数，go文件初始化的时候调用，可以有多个，执行顺序从上往下
func init() {
	fmt.Println("初始化function_def的go文件")

}
