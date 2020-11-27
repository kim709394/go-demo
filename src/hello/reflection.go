package hello

import (
	"fmt"
	"reflect"
)

/*
@Author kim
@Description   反射
@date 2020-11-27 9:57
*/

type myStu struct{}

//获取变量的类型信息
func TypeOf() {

	a := 10
	typeOf := reflect.TypeOf(a)
	fmt.Println("变量a的类型名:", typeOf.Name(), ",种类名:", typeOf.Kind())
	//指针类型，map，slice，sturct等类型的类型信息
	my := &myStu{}
	myTypeOf := reflect.TypeOf(my)
	fmt.Println("my name:", myTypeOf.Name(), "my kind:", myTypeOf.Kind())
	elem := myTypeOf.Elem()
	fmt.Println("element name:", elem.Name(), "element kind:", elem.Kind())
}

type dog struct {
	Name  string
	Age   int
	color byte
}

//通过反射获取结构体内部信息
func GetStructMsg() {
	myDog := dog{"旺财", 2, 'y'}
	typeOf := reflect.TypeOf(myDog)
	//遍历结构所有成员变量
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fmt.Println(field)
	}
	//通过字段名获取字段
	fmt.Println(typeOf.FieldByName("color"))
	//获取结构体全路径信息
	fmt.Println(typeOf.PkgPath())
}

//获取结构体标签信息
func GetStructTag() {

	dol := Dolphin{"myDo"}
	typeOf := reflect.TypeOf(dol)
	name, _ := typeOf.FieldByName("name")
	fmt.Println("标签信息:", name.Tag)
}
