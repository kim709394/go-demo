package main

/*
@Author kim
@Description
@date 2020-10-19 14:30
*/

import (
	"fmt"
	"hello"
	"time"
)

func main() {
	//输出hello world
	fmt.Println("hello1 world")
	fmt.Println(hello.Function(1, "b"))
	fmt.Println('b')
	fmt.Println(hello.GlobalVal, "kk")
	v1 := 1

	fmt.Println(v1)
	fmt.Printf("%.2f\n", hello.Sdk1())
	fmt.Printf("%f\n", hello.Sdk1())
	hello.Complex()
	hello.Str()
	hello.Float()
	hello.Ptr()
	hello.Loop()
	hello.Array()
	hello.MapAndCollections()
	hello.List()
	hello.Swh()
	hello.Brek()
	hello.Contu()
	fmt.Println(hello.F())
	hello.EnableParam()
	hello.Der()
	hello.InsStruct()
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("程序结束")
}
