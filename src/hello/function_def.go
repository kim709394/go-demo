package hello

/*
@Author kim
@Description  函数的定义和规则
@date 2020-10-21 15:42
*/
import (
	"fmt"
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
