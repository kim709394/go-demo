package hello

/*
@Author kim
@Description
@date 2020-10-19 14:30
*/
import (
	"fmt"
	"math"
	"strings"
)

//
func Sdk1() float64 {
	return math.Pi
}

/***********内置函数*******************/
func length() {
	//获取字符串或者数组的长度
	fmt.Println(len("123"))
	fmt.Println(len([]int{1, 2}))
}

/***********内置函数*******************/

//strings操作
func Strs() {
	//分割函数
	spli := strings.Split("1,2", ",")
	fmt.Println(spli)
	//是否包含
	strings.Contains("go是世界上最好的语言", "go")
	//是否有前缀
	strings.HasPrefix("123.png", "123")
	//是否有后缀
	strings.HasSuffix("123.jpg", "jpg")
	//返回字符下标
	strings.Index("123545", "5")
	//返回最后一次出现的字符的下标
	strings.LastIndex("123545", "5")
	//将字符加入一个字符串数组之间，例如：如下输出结果:a|b|c
	strings.Join([]string{"a", "b", "c"}, "|")

}
