package hello

/*
@Author kim
@Description   go语言基础语法
@date 2020-10-19 14:30
*/
import (
	"container/list"
	"fmt"
	"math"
	"sync"
	"unicode/utf8"
)

var GlobalVal int = 100 //全局变量,首字母大写相当于publc,小写相当于private
//函数定义：func(变量名 变量类型)(返回值类型){}
func Function(a int, b string) (int, string) {
	return a, b
}

//变量声明
func VariableInit() {
	/*变量声明后未使用则会编译报错*/
	//单个变量声明
	var value int
	//多个变量声明
	/*Go语言的基本类型有：
	bool
	string
	int、int8、int16、int32、int64
	uint、uint8、uint16、uint32、uint64、uintptr
	byte // uint8 的别名
	rune // int32 的别名 代表一个 Unicode 码
	float32、float64
	complex64、complex128*/
	var (
		val1 int
		val2 string
		val3 bool
		val4 float32
	)
	//简短式声明变量,
	/*简短模式（short variable declaration）有以下限制：
	*定义变量，同时显式初始化。
	*不能提供数据类型。
	*只能用在函数内部。*/
	val0 := 10
	//变量赋值
	value = 11
	val1 = 1
	val2 = "1"
	val3 = false
	val4 = 1.9
	//多个变量直接声明并赋初始值
	val5, val6 := "val5", 66
	//编译器根据右边值类型自动定义变量的类型
	var val7 = "自动编译"
	//多重赋值,简化变量交换算法
	value, val1 = val1, value

	/*匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，
	被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），
	但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。
	使用匿名变量时，只需要在变量声明的地方使用下画线替换即可*/
	val8, _ := Function(200, "300")

	fmt.Println(val0)
	fmt.Println(value + val1)
	fmt.Println(val2)
	fmt.Println(val3)
	fmt.Println(val4)
	fmt.Println(val5)
	fmt.Println(val6)
	fmt.Println(val7)
	fmt.Println(val8)

}

//各基本数据类型的初始值
func GetBasicInitValue() {
	var (
		bool1       bool
		string1     string
		int1        int
		int81       int8
		int161      int16
		int321      int32
		int641      int64
		uint1       uint
		uint81      uint8
		uint161     uint16
		uint321     uint32
		uint641     uint64
		uintptr1    uintptr
		byte1       byte
		rune1       rune // int32 的别名 代表一个 Unicode 码
		float321    float32
		float641    float64
		complex641  complex64
		complex1281 complex128
	)
	fmt.Println("bool:", bool1)
	fmt.Println("string:", string1)
	fmt.Println("int:", int1)
	fmt.Println("int8:", int81)
	fmt.Println("int16:", int161)
	fmt.Println("int32:", int321)
	fmt.Println("int64:", int641)
	fmt.Println("uint8:", uint81)
	fmt.Println("uint:", uint1)
	fmt.Println("uint16:", uint161)
	fmt.Println("uint32:", uint321)
	fmt.Println("uint64:", uint641)
	fmt.Println("uintptr:", uintptr1)
	fmt.Println("rune1:", rune1)
	fmt.Println("byte:", byte1)
	fmt.Println("float32:", float321)
	fmt.Println("float64:", float641)
	fmt.Println("complex64:", complex641)
	fmt.Println("complex128:", complex1281)

}

func Complex() {
	fmt.Println("--------------------复数-------------------------")
	//复数
	z := complex(1.2, 2.3)
	fmt.Println(real(z)) //复数实部
	fmt.Println(imag(z)) //复数虚部

}

//布尔值运算
func boolean() {
	fmt.Println("---------------------布尔值运算------------------------")
	a := false
	b := true
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(a != b)
	fmt.Println(a == b)
	fmt.Println(!a)

}

//字符串运算
/* \n：换行符
\r：回车符
\t：tab 键
\u 或 \U：Unicode 字符
\\：反斜杠自身 */
func String() {
	fmt.Println("------------------字符串运算---------------------------")
	s := "GO语言是世界上最好的编程语言"
	s2 := "吗"
	fmt.Println(s)
	fmt.Println(s + s2)
	fmt.Println("换行符:" + "\n")
	fmt.Println("回车符:" + "\r")
	fmt.Println("tab键:" + "\t")
	fmt.Println("反斜杠:" + "\\")
	//定义多行字符串
	s3 := `第一行
第二行
第三行
第四行`
	fmt.Println(s3)
	fmt.Println(len(s2))                    //len()函数获取字节长度
	fmt.Println(utf8.RuneCountInString(s2)) //获取字符串有多少个字符

}

//浮点型数据类型,有float64和float32两种，通常使用float64
func Float() {
	fmt.Println("----------------浮点型数据类型,有float64和float32两种，通常使用float64-----------------------------")
	f := 1.666
	fmt.Println(f)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(int(f)) //浮点型强制转换为整型

}

//指针
func Ptr() {
	fmt.Println("----------------指针-----------------------------")
	v := "go语言"
	ptr := &v //用&符号表示取变量的指针，也就是内存地址
	fmt.Println(ptr)

	fmt.Println(*ptr) //用*获取指针指向的地址的变量的值
	//指针赋值
	*ptr = "c语言" //将指针指向的变量的值进行赋值，则变量的值就被修改了
	fmt.Println(v)
	//用new()函数创建指针
	ptr1 := new(string)
	*ptr1 = "go语言学习"

	fmt.Println(*ptr1)
	//指针类型,类型前加*号，**string表示指针类型*string的指针
	var pt *string
	p := "java"
	pt = &p
	*pt = "rust"
	fmt.Println(p)
}

//常量使用
//只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。由于编译时的限制，定义常量的表达式必须为能被编译器求值的常量表达式
func constant() {

	//声明单个常量
	const c0 = 1.15
	//显式说明常量数据类型
	const c1 int = 64
	//声明多个常量
	const (
		c2 = 25
		c3 = "go语言"
	)

	//iota常量生成器,iota初始化为0，以下每个常量递增一，周天是0，周一是1，以此类推
	const (
		Sunday    = iota //0
		Monday           //1
		Tuesday          //2
		Wednesday        //3
		Thursday         //4
		Friday           //5
		Saturday         //6
		_                //"7"  跳过该值
		Next             //8
	)
}

//类型别名和类型定义
func typeAlias() {

	//给类型起别名
	type myInt = int
	//定义自定义类型
	type myFloat float64
}

/**
go的关键字:
break	default 	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var
*/

//函数名首字符大写表示可以被其他地方访问，类似java的public。否则只能本地访问，类似java的private

//流程控制
func flowContro() {
	fmt.Println("-----------------流程控制----------------------------")
	//选择结构，和java大致一样，只是条件不需要加小括号
	i := 11
	if i < 10 {
		//执行业务逻辑
	}
	if i > 10 {

	} else if i < 9 {

	} else {

	}
	if k := 3 * 4; k != 0 {

	}
}

//循环,只支持for，和java差不多，就是少了小括号
func Loop() {
	fmt.Println("-----------------循环,只支持for，和java差不多，就是少了小括号----------------------------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//无限循环
	/*for{
		//循环业务代码
	}*/

	//break语句
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break
			} else {
				continue
			}
		}
	}

	//增强for，类似java的foreach，for range，k是索引，v是数值
	for k, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(k)
		fmt.Println(v)
	}

}

//数组与切片
func Array() {
	fmt.Println("-----------------数组与切片----------------------------")
	//初始化数组，数组必须要有长度，否则就是切片
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(array))
	fmt.Println(array[0]) //取得第一个元素
	//多维数组
	multiArray := [2][2]int{{1, 2}, {3, 4}} //二维数组
	multiArray3 := [3][4][5]int{}           //三维数组
	fmt.Println(multiArray[0][0])
	fmt.Println(multiArray3)
	//声明数组
	var arr [10]int
	var multiArr [5][6]int
	var multiArr1 [1][2][3][4][5][6][7]string //声明多维数组
	fmt.Println(arr)
	fmt.Println(multiArr)
	fmt.Println(multiArr1)
	//切片
	//从原数组里面生成切片，类似java的字符串的substring方法
	sliceArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice0 := sliceArr[0:2] //从数组中切取索引0到2的元素形成新的切片，包前不包后
	fmt.Println(slice0)
	slice1 := sliceArr[3:]  //从索引3到最后，包含最后的元素
	slice2 := sliceArr[:]   //不指定开始和结束索引，相当于以完整原数组生成一个切片
	slice3 := sliceArr[0:0] //初始和结束索引都为0，生成的切片是空切片
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	//切片，遍历方式和数组一样可以用for range来遍历
	for _, v := range slice1 {
		fmt.Println(v)
	}
	//判断切片或者数组是否为空
	fmt.Println(slice0 == nil)
	//使用make()函数声明切片
	slice4 := make([]int, 10, 20) //第二个参数是指定切片大小，第三个参数是预分配的元素数量
	slice4[0] = 1
	fmt.Println(slice4[0])

	//append()函数在切片追加元素
	var slice5 []int
	slice5 = append(slice5, 1, 2) //往切片后面追加元素
	fmt.Println(slice5)
	slice5 = append(slice5, []int{3, 4}...) //往切片后面追加一个切片
	fmt.Println(slice5)
	/*在使用 append() 函数为切片动态添加元素时，如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。

	切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……*/
	//在切片前面追加元素
	slice5 = append([]int{0}, slice5...)
	//使用copy函数赋值切片,
	slice6 := []int{1, 2, 3, 4}
	slice7 := []int{5, 6, 7}
	//copy(slice6,slice7)    //把slice7复制到slice6，只能把前三个元素复制到slice6的前三个元素，数量少的复制到数量多的
	fmt.Println(slice6)
	//把slice6复制到slice7，只把slice6的前三个元素复制到slice7的前三个元素，如果两个切片元素数量一致则复制元素相同，返回被复制的元素数量
	fmt.Println(copy(slice7, slice6))
	fmt.Println(len(slice7)) //获取切片的元素个数
	/*删除切片的元素，利用切片中切取一部分元素作为新切片的方式来进行删除，
	  go没有直接的删除切片元素的方法
	*/

}

//map
func Map() {
	fmt.Println("------------------------map与集合-------------------------")
	//声明map,与java的Map类似，键值对，无序，容量会自动扩充，也可以初始指定容量大小，map是无序的
	var mapList0 map[string]int
	mapList0 = map[string]int{"one": 1, "two": 2} //给map赋值
	mapList1 := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	//使用make()函数来初始化map,指定初始容量为10
	mapList2 := make(map[string]int, 10)
	fmt.Println(len(mapList0))
	//通过key值获取value值
	fmt.Println(mapList0["one"])
	//根据key值对键值对进行赋值
	mapList0["two"] = 22
	fmt.Println(mapList1)
	fmt.Println(mapList2)
	//用切片作为map的值
	mapList3 := make(map[string][8]int, 10)
	fmt.Println(mapList3)
	//判断map是否包含某个键
	k, f := mapList0["k"]
	if f {
		fmt.Println("k=", k)
	} else {
		fmt.Println("该key值:k不存在")
	}
	//用for range遍历map,遍历键值对，注意是无序的，遍历顺序不会是填充顺序
	for k, v := range mapList1 {
		fmt.Println(k)
		fmt.Println(v)
	}

	//只遍历键值
	for k := range mapList1 {
		fmt.Println(k)
	}

	//map删除键值对元素
	delete(mapList1, "k1")

	//只遍历值
	for _, v := range mapList1 {
		fmt.Println(v)
	}

	//清空map键值对集合，重新创建一个map集合即可

	//map是线程不安全的，sync包下的map是线程安全的
	var syncMap sync.Map //初始化线程安全的map
	/*sync.Map 有以下特性：
	无须初始化，直接声明即可。
	sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
	使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，
	返回 true，终止迭代遍历时，返回 false。*/

	//保存键值对
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	//根据键取值
	value1, ok := syncMap.Load("key1")
	fmt.Println(value1)
	fmt.Println(ok)
	//根据键值删除键值对
	syncMap.Delete("key1")
	//遍历,继续遍历时返回true，终止遍历时返回false
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}

func DeleteSyncMap() {
	var syncMap sync.Map
	syncMap.Store("k1", "v1")
	syncMap.Store("k2", "v2")
	fmt.Println(syncMap)
	syncMap.Delete("k1")
	fmt.Println(syncMap)
}

//List列表(集合)
func List() {
	fmt.Println("------------------------List列表(集合)-------------------------")
	//初始化列表
	l := list.New()
	var l0 list.List
	//往集合后面添加元素,返回这个元素
	element := l.PushBack("p1")
	//往集合前面添加元素
	l0.PushFront("p2")

	//往某个元素前面插入元素
	l.InsertBefore("p0", element)
	//往某个元素后面插入元素
	l.InsertAfter("p3", element)
	//删除元素
	l.Remove(element)

	//l.PushBackList(&l0)
	//l.PushFrontList(&l0)
	fmt.Println(l)
	//遍历列表
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}

	//nil 是 map、slice、pointer、channel、func、interface 的初始值，引用类型

}

//switch语句
func Swh() {
	fmt.Println("------------------------switch语句-------------------------")
	//不需要写break跳出switch，只允许有一个default,变量不需要为整型也不需要为常量
	a := "go"
	switch a {
	case "go":
		fmt.Println(a)
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("default")
	}

	//一分支多值
	switch a {
	case "java", "c", "c++":
		fmt.Println("other")
	default:
		fmt.Println("default")
	}
	//分支表达式
	b := 10
	switch {
	case b >= 10:
		fmt.Println("b大于等于10")
	}

	//即使匹配后仍然往下执行，fallthough关键字
	c := 20
	switch {
	case c < 30 && c > 10:
		fmt.Println("c小于30并且大于10")
		fallthrough
	case c == 20:
		fmt.Println("c等于20")
	}

}

//goto
func Got() {
	fmt.Println("------------------------goto-------------------------")
	//goto退出任意循环圈，进入到任意一个循环圈或者任意的其他位置
breakBefore:
	fmt.Println("退出到循环前面")
	for x := 0; x < 10; x++ {

		for y := 0; y < 5; y++ {
			if y == 3 {
				goto breakBefore
			} else {
				goto breakAfter
			}
		}

	}
breakAfter:
	fmt.Println("退出到循环后面")
}

//break
func Brek() {
	fmt.Println("------------------------break-------------------------")
	//退出当前循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	//退出所有循环
codeBlock:
	for i := 0; i < 10; i++ {
		fmt.Println("外层循环")
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break codeBlock
			}
		}
	}
}

//continue
func Contu() {
	fmt.Println("------------------------continue-------------------------")
	//退出当轮循环
	for i := 0; i < 10; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	//退出外层当轮循环
outerLoop:
	for i := 0; i < 10; i++ {

		for y := 0; y < 5; y++ {
			fmt.Println(i, y)
			if y == 2 {
				continue outerLoop
			}
		}

	}

}

//类型断言
func Assertion() {
	fmt.Println("------------------------类型断言-------------------------")
	//判断变量的类型，得先声明一个interface{}变量
	var ass interface{}
	ass = 10
	//判断变量是否是int类型，返回变量的值和一个布尔值，如果类型匹配返回true，否则返回false
	value, ok := ass.(int)
	fmt.Println(value, ok)
	var ass2 interface{}
	ass2 = "str"
	value2 := ass2.(string)
	fmt.Println(value2)

}

//异常捕捉
func Exception() {
	//recover捕捉异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//panic函数抛出异常
	panic("error")
}

func Scan() {
	var scan string
	fmt.Println("请输入内容:")
	_, err := fmt.Scanln(&scan)
	if err != nil {
		fmt.Println("报错:", err)
		return
	}
	fmt.Println(scan)
}

//自定义异常
type MyError struct {
	Code int
	Msg  string
}

func (myErr *MyError) Error() string {
	return string(myErr.Code) + myErr.Msg
}
