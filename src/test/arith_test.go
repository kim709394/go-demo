package test

import (
	"fmt"
	"hello"
	"testing"
)

/*
@Author kim
@Description
@date 2020-11-11 11:20
*/

//插入排序算法
func TestInsertSort(t *testing.T) {
	slice := []int{1, 6, 3, 4, 2, 5, 9, 0, 8}
	hello.InsertSort(slice, true)
	fmt.Println(slice)
	hello.InsertSort(slice, false)
	fmt.Println(slice)

}

//希尔排序算法
func TestShellSort(t *testing.T) {
	slice := []int{1, 8, 5, 4, 3, 2, 6, 7, 0, 4, 5, 6, 6, 10, 13, 69, 11}
	hello.ShellSort(slice, true)
	fmt.Println(slice)
	hello.ShellSort(slice, false)
	fmt.Println(slice)
}
