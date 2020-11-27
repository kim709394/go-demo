package hello

import "fmt"

/*
@Author kim
@Description   结构体定义，类似于java的类
@date 2020-10-23 14:38
*/

/*Go语言可以通过自定义的方式形成新的类型，结构体就是这些类型中的一种复合类型，结构体是由零个或多个任意类型的值聚合成的实体，每个值都可以称为结构体的成员。

结构体成员也可以称为“字段”，这些字段有以下特性：
字段拥有自己的类型和值；
字段名必须唯一；
字段的类型也可以是结构体，甚至是字段所在结构体的类型。*/

type MyStruct struct {
	column1 int
	column2 string
	column3 bool
}

//结构体的实例化
func InsStruct() {
	fmt.Println("-------------------------------------结构体的实例化-----------------------------------------")
	//当做数据类型进行初始化
	var my MyStruct
	fmt.Println(my)
	//调出成员变量
	fmt.Println(my.column1)
	//实例化结构体作为指针类型,*MyStruct
	my1 := new(MyStruct)
	fmt.Println(my1.column2)
	//在Go语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
	my2 := &MyStruct{}
	fmt.Println(my2)
	//使用键值对形式进行成员变量初始化
	my3 := &MyStruct{
		column1: 123,
		column2: "123",
		column3: true,
	}
	fmt.Println(my3.column3)
	//使用多个值列表进行初始化
	/*使用这种格式初始化时，需要注意：
	必须初始化结构体的所有字段。
	每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	键值对与值列表的初始化形式不能混用。*/
	my4 := MyStruct{
		321,
		"321",
		false,
	}
	fmt.Println(my4.column2)
	//匿名结构体的实例化和初始化
	my5 := struct {
		co1 int
		co2 string
		co3 bool
	}{
		123,
		"123",
		false,
	}
	fmt.Println(my5)

	//结构体派生，相当于java的父子类继承，但是只有属性的派生，却没有java的多态
	cat := BlackCat{}
	cat.color = "yellow"
	cat.age = 11

}

//结构体：猫
type Cat struct {
	color string
	age   int
}

//结构体:黑猫
type BlackCat struct {
	//将猫结构体内嵌进来，用于派生，类似java的继承猫类，可以服用猫结构体的全部属性
	Cat
}

//结构体内嵌
func InnerAndOuter() {

	outer := &Outer{}
	outer.x = 1
	outer.b = "1"
	outer.x = 2
	outer.y = 3
	outer.z = "3"
	outer.int = 10
	outer.string = "10"
	fmt.Println(outer.Inner)
	//结构体初始化
	outer1 := Outer{10, "b", 20, "string", Inner{1, 2, "3"}}
	fmt.Println(outer1.Inner.x)
	fmt.Println(outer1.Inner)
	fmt.Println(outer1.x)
	//初始化内嵌结构体
	car := &Car{
		Engine{
			power: 10,
			style: "BMW",
		},
		Wheel{
			size: 100,
		},
	}
	fmt.Println(car)
	//初始化内嵌匿名结构体
	car0 := Engine{
		power: 100,
		style: "BENZ",
		//初始化内部匿名结构体
		Start: struct {
			power int
		}{
			power: 20,
		},
	}
	fmt.Println(car0)

}

/*结构内嵌特性
Go语言的结构体内嵌有如下特性。
1) 内嵌的结构体可以直接访问其成员变量
嵌入结构体的成员，可以通过外部结构体的实例直接访问。如果结构体有多层嵌入结构体，结构体实例访问任意一级的嵌入结构体成员时都只用给出字段名，而无须像传统结构体字段一样，通过一层层的结构体字段访问到最终的字段。例如，ins.a.b.c的访问可以简化为ins.c。
2) 内嵌结构体的字段名是它的类型名
内嵌结构体字段仍然可以使用详细的字段进行一层层访问，内嵌结构体的字段名就是它的类型名，代码如下：
var c Color
c.BasicColor.R = 1
c.BasicColor.G = 1
c.BasicColor.B = 0
一个结构体只能嵌入一个同类型的成员，无须担心结构体重名和错误赋值的情况，编译器在发现可能的赋值歧义时会报错。*/
type Inner struct {
	x int
	y int
	z string
}

type Outer struct {
	x   int
	b   string
	int //成员变量可以没有名字，此时类型名字即是成员名字,但是不可以有两个同样类型没有名字的成员变量
	string
	Inner //内嵌结构体，派生，此时可以继承Inner结构体的全部成员变量
}

type Wheel struct {
	size int
}

type Engine struct {
	power int
	style string
	//类似java的内部类，内部直接定义匿名结构体
	Start struct {
		power int
	}
}
type Car struct {
	Engine
	Wheel
}

//结构体标签
type Dolphin struct {
	name string `key1:"val1 key2:val2"` //标签：为字段添加补充信息
}
