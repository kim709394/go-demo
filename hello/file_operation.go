package hello

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
@Author kim
@Description   文件操作
@date 2020-11-27 9:01
*/

//设备文件：【屏幕输出，控制台显示，以及键盘输入
func DeviceInput() {
	os.Stdout.Close()
	fmt.Println("无法打印，因为控制台流已关闭")
}

//创建文件
func CreateFile() {
	//创建一个文件，从无到有，如果本身存在则会覆盖
	myFile, err := os.Create("../file/writ.txt") //当前go文件所在路径的相对路径
	//myFile,err := os.Create("d:/writ.txt")   //文件的绝对路径
	if err != nil {
		fmt.Println("err:", err)
		return
	} else {
		defer myFile.Close()
		myFile.WriteString("go写文件")
	}
}

//读文件
func ReadFile() {

	//打开文件
	myFile, err := os.Open("../file/writ.txt")
	if err != nil {
		fmt.Println("file open err:", err)
		return
	} else {
		defer myFile.Close()
		b := make([]byte, 1024*1024)
		//读取一定数量的字节码存放在字节码数组中，最多读取并存放字节码数组的长度数量的字节码，返回读取的字节码数量
		read, err := myFile.Read(b)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(string(b[:read]))
	}
}

//分行读取文件内容
func ReadTxtLine() {

	//打开文件
	myFile, err := os.Open("../file/writ.txt")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer myFile.Close()
		//将文件放入缓冲区,新建缓冲流
		reader := bufio.NewReader(myFile)
		for {
			//读取文件当遇到'\n'换行符的时候就结束，包含读取的反斜杠，返回读取到的内容
			bytes, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("文件读取结束")
					break
				}
				fmt.Println("err:", err)
			}
			fmt.Print(string(bytes))
		}
	}
}

//文件拷贝
func CopyFile() {

	//读取源文件
	addBill, err1 := os.Open("../file/addWorkBill.png")
	//创建目标文件
	copy, err2 := os.Create("../file/copyAddWorkBill.png")
	//文件先开后闭原则
	defer addBill.Close()
	defer copy.Close()
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	if err2 != nil {
		fmt.Println("err2:", err2)
		return
	}
	//缓冲字节码切片
	b := make([]byte, 1024*1024)
	//拷贝文件
	for {
		//读取一定数量字节码存到字节码切片中，返回读取的字节码数量
		read, err3 := addBill.Read(b)
		if err3 != nil {
			fmt.Println("err3:", err3)
			//文件读取结束时，err是EOF，则说明拷贝结束，退出拷贝
			if err3 == io.EOF {
				fmt.Println("文件读取结束")
				break
			}
		}
		//将源文件读取的字节码写入目标文件中
		copy.Write(b[:read])
	}
}

//查看文件信息
func FileState() {
	stat, err := os.Stat("../file/writ.txt")
	if err != nil {
		fmt.Println("file stat err:", err)
		return
	}
	fmt.Println(stat)
}
