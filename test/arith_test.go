package test

import (
	"fmt"
	"go-demo/hello"
	"testing"
)

/*
@Author kim
@Description
@date 2020-11-11 11:20
*/
var slice = []int{1, 8, 5, 4, 3, 2, 6, 7, 0, 4, 5, 6, 6, 10, 13, 69, 11}

//插入排序算法
func TestInsertSort(t *testing.T) {
	hello.InsertSort(slice, true)
	fmt.Println(slice)
	hello.InsertSort(slice, false)
	fmt.Println(slice)

}

//希尔排序算法
func TestShellSort(t *testing.T) {
	hello.ShellSort(slice, true)
	fmt.Println(slice)
	hello.ShellSort(slice, false)
	fmt.Println(slice)
}

//快速排序
func TestQuickSort(t *testing.T) {

	hello.QuickSort(slice, true)
	fmt.Println(slice)
	hello.QuickSort(slice, false)
	fmt.Println(slice)
}
