package hello

/*
@Author kim
@Description
@date 2020-10-19 14:30
*/
import (
	"encoding/json"
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

type myJson struct {
	Name   string   `json:"name"` //``里面的相当于转义字符串""双引号的字符串类型，添加结构体变迁指定序列化后的json的key值
	Age    int      `json:"age"`  // `json:"-"` 这样则该字段不会被序列化和反序列化
	Lang   []string `json:"lang"`
	IsTrue bool     `json:"isTrue"`
	Price  float64  `json:"price"`
}

//json序列化
func JsonSerialize() {
	//初始化结构体，要序列化成json的结构体的成员变量首字母必须大写，如果要序列化成小写，要添加结构体标签指定
	j := myJson{Name: "myJson", Age: 18, Lang: []string{"go", "c", "java", "c", "c艹"}, IsTrue: true, Price: 999.999}
	marshal, err := json.Marshal(j)
	if err != nil {
		//处理错误
		fmt.Println("err :", err)
		return
	} else {
		//将转化后的字节码转换为字符串
		fmt.Println(string(marshal))
	}
}

//json序列化map
func JsonSerializeMap() {
	myMap := make(map[string]interface{}, 2)
	myMap["k1"] = "v1"
	myMap["k2"] = []string{"go", "c", "java", "c艹"}
	myMap["k3"] = 444.33
	myMap["k4"] = false
	marshal, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("err:", err)
		return
	} else {
		fmt.Println(string(marshal))
	}
}

//json反序列化成结构体
func JsonUnSerialize() {
	jsonStr := `{"name":"myJson","age":18,"lang":["go","c","java","c","c艹"],"isTrue":true,"price":999.999}`
	my := new(myJson)
	err := json.Unmarshal([]byte(jsonStr), my)
	if err != nil {
		fmt.Println("err:", err)
		return
	} else {
		fmt.Println(my)
	}
}

//json反序列化到map
func JsonUnSerializeMap() {
	jsonStr := `{"name":"myJson","age":18,"lang":["go","c","java","c","c艹"],"isTrue":true,"price":999.999}`
	myMap := make(map[string]interface{}, 2)
	//转换成map对象,要传指针类型
	err := json.Unmarshal([]byte(jsonStr), &myMap)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//获取map的每个键值对，键为字符串，值是空接口万能类型，因此需要类型断言进行数据类型转换
	s, f := myMap["name"].(string)
	if f {
		fmt.Println("name", s)
	}
	age, af := myMap["age"].(float64) //数字通常反序列化成float64类型
	if af {
		fmt.Println("age:", age)
	}
	//切片通常会反序列化成空接口切片

	lang, lf := myMap["lang"].([]interface{})
	if lf {
		for _, value := range lang {
			fmt.Println(value)
		}
	}
	isTrue, isF := myMap["isTrue"].(bool)
	if isF {
		fmt.Println("isTrue:", isTrue)
	}
	price, pf := myMap["price"].(float64)
	if pf {
		fmt.Println(price)
	}

}
