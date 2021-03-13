package hello

import (
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"reflect"
)

/*
@Author kim
@Description   反射
@date 2020-11-27 9:57
*/

type MyStu struct {
	Sex   int
	Grade int
}

//获取变量的类型信息
func TypeOf() {

	a := 10
	typeOf := reflect.TypeOf(a)
	fmt.Println("变量a的类型名:", typeOf.Name(), ",种类名:", typeOf.Kind())
	//指针类型，map，slice，sturct等类型的类型信息
	my := &MyStu{}
	myTypeOf := reflect.TypeOf(my)
	fmt.Println("my name:", myTypeOf.Name(), "my kind:", myTypeOf.Kind())
	elem := myTypeOf.Elem()
	fmt.Println("element name:", elem.Name(), "element kind:", elem.Kind())
}

type Dog struct {
	Name      string
	Age       int
	Color     byte
	CreatedAt *gtime.Time
	MyStu
}

//通过反射获取结构体内部信息
func GetStructMsg() {
	myDog := Dog{"旺财", 2, 'y', gtime.Now(), MyStu{1, 2}}
	typeOf := reflect.TypeOf(myDog)
	valueOf := reflect.ValueOf(myDog)
	//遍历结构所有成员变量
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		value := valueOf.Field(i)
		fmt.Println("字段类型名:", field.Type.Name())
		fmt.Println("字段名", field.Name)
		fmt.Println("字段值:", value.Interface())
	}
	//通过字段名获取字段
	value, _ := typeOf.FieldByName("color")
	fmt.Println(value.Name)

	//获取结构体全路径信息
	fmt.Println(typeOf.PkgPath())
}

//通过反射排除结构体中空值的属性
func IgnoreStructNull(i interface{}) map[string]interface{} {
	fmt.Println(i)
	typeOf := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)
	m := make(map[string]interface{})
	resolve(typeOf, valueOf, m)
	return m
}

func resolve(t reflect.Type, v reflect.Value, m map[string]interface{}) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		switch field.Type.Kind() {
		case reflect.Struct:
			resolve(field.Type, value, m)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			m[field.Name] = value.Uint()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[field.Name] = value.Int()
		case reflect.Bool:
			m[field.Name] = value.Bool()
		case reflect.Float32, reflect.Float64:
			m[field.Name] = value.Float()
		case reflect.Complex64, reflect.Complex128:
			m[field.Name] = value.Complex()
		case reflect.Chan, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Func, reflect.Slice:
			if value.IsNil() {
				continue
			}
			if field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.Uint8 {
				m[field.Name] = value.Bytes()
			} else {
				m[field.Name] = value.Pointer()
			}
		case reflect.Interface:
			m[field.Name] = value.InterfaceData()
		case reflect.String:
			m[field.Name] = value.String()
		default:
			m[field.Name] = value.Interface()
		}
	}
}

//获取结构体标签信息
func GetStructTag() {

	dol := Dolphin{"myDo"}
	typeOf := reflect.TypeOf(dol)
	name, _ := typeOf.FieldByName("name")
	fmt.Println("标签信息:", name.Tag)
}
