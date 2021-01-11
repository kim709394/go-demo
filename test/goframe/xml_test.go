package goframe

import (
	"fmt"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"testing"
)

/*
@Author kim
@Description    处理xml文件
@date 2021-1-6 16:05
*/

//解析xml文件
func TestParseXml(t *testing.T) {
	parseXml("C:\\Program Files\\FreeSWITCH\\conf\\dialplan\\default.xml")
}

func parseXml(path string) map[string]interface{} {
	bytes := gfile.GetBytes(path)
	decode, err := gxml.Decode(bytes)
	if err != nil {
		fmt.Println("xml解析错误:", err)
	}
	fmt.Println(decode)
	return decode
}

//修改xml文件
func TestUpdateXml(t *testing.T) {
	xml := parseXml("../../file/default.xml")
	m, ok := xml["include"].(map[string]interface{})
	if ok {
		fmt.Println(m["context"].(map[string]interface{})["-name"])
		m["context"].(map[string]interface{})["-name"] = "default123"
	}

	encode, err := gxml.Encode(xml)
	if err != nil {
		fmt.Println("xml输出异常:", err)
	}
	fmt.Println(string(encode))
	err2 := gfile.PutBytes("../../file/default.xml", encode)
	if err2 != nil {
		glog.Info("输出xml文件报错:", err2)
	}
}
