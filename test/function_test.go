package test

import (
	_ "go-demo/hello" //初始化导入hello包，仅仅只是调用此包下的所有初始化函数，无法用该包调用其他函数
	"testing"
)

/*
@Author kim
@Description
@date 2020-11-20 15:30
*/

//测试初始化函数
func TestInit(t *testing.T) {

}
