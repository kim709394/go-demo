package hello

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

/*
@Author kim
@Description   网络编程
@date 2020-11-24 11:15
*/

//get方法服务接口
func Get() {

	http.HandleFunc("/my/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.RawQuery)
		fmt.Println(writer)
		writer.Write([]byte("myGoHttp"))
	})
	http.ListenAndServe("127.0.0.1:8888", nil)

}

//tcp、socket编程,转换大小写
//tcp服务端编程
func TcpServer() {
	//开启tcp协议服务端监听
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("tcp conn err :", err)
		return
	}

	for {
		//开启无限循环监听
		accept, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("tcp accept err:", err1)
			continue
		}
		go func(conn net.Conn) {
			defer accept.Close()

			acceptMsg := readByConn(accept)
			fmt.Println("从客户端发送过来的消息:", acceptMsg)
			//给客户端返回消息
			_, err3 := accept.Write([]byte(strings.ToUpper(acceptMsg)))
			if err3 != nil {
				fmt.Println("write err3:", err3)
				return
			}
		}(accept)
	}
}

//tcp客户端
func TcpCustomer() {
	//连接服务端
	dial, err := net.Dial("tcp", "127.0.0.1:8888")
	defer dial.Close()
	if err != nil {
		fmt.Println("conn server err:", err)
		return
	}
	//往服务端写数据
	_, err1 := dial.Write([]byte("my lowwer code"))
	if err1 != nil {
		fmt.Println("customer write err", err1)
		return
	}
	//接收服务端响应过来的数据
	response := readByConn(dial)
	fmt.Println("服务端响应:", response)

}

//读取数据
func readByConn(conn net.Conn) string {
	b := make([]byte, 1024*1024)
	//读取客户端发过来的消息
	read, err2 := conn.Read(b)
	acceptMsg := string(b[:read])
	if err2 != nil {
		fmt.Println("read err:", err2)
	}
	return acceptMsg
}
