package goframe

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"strconv"
	"strings"
	"testing"
)

//字符串截取
func TestSubStr(t *testing.T) {
	//获得后缀名
	s := "123.jpg"
	//获取最后一次出现"."的下标
	index := strings.LastIndex(s, ".")
	g.Dump(index)
	//截取从输入下标为起始到最后的字符串，最后一个参数如果没有就是截取到最后，如果有就是截取长度，包前也包后
	str := gstr.SubStr(s, strings.LastIndex(s, ".")+1)
	g.Dump(str)
	s = "1"
	str = gstr.SubStr(s, 0, len(s)-1)
	g.Dump(str)

}

//整数转化为浮点型数
func TestIntToFloat(t *testing.T) {

	f := gconv.Float64(1024*1024+1) / 1024 / 1024
	g.Dump(f)
	//四舍五入保留两位小数
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 9.8345), 64)
	g.Dump(value)
	//四舍五入保留三位小数
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", 9.8345), 64)
	g.Dump(value)
}

//字符串分割
func TestSplit(t *testing.T) {
	s := "1,2,3,4,5,6,7"
	array := gstr.Split(s, ",")
	g.Dump(array)
}
