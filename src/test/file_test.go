package test

import (
	"hello"
	"testing"
)

/*
@Author kim
@Description
@date 2020-11-27 9:57
*/

//测试控制台输出
func TestPrintln(t *testing.T) {
	hello.DeviceInput()
}

//测试写文本文件
func TestWriteTxt(t *testing.T) {
	hello.CreateFile()
}

//测试读文本文件
func TestReadTxt(t *testing.T) {
	hello.ReadFile()
}

//测试分行读取文件
func TestReadFileLine(t *testing.T) {
	hello.ReadTxtLine()
}

//测试copy文件
func TestCopyFile(t *testing.T) {
	hello.CopyFile()
}

//测试获取文件信息
func TestFileStat(t *testing.T) {
	hello.FileState()
}
