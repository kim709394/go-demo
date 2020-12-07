package hello

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
@Author kim
@Description   网络编程
@date 2020-11-24 11:15
*/

//http编程
//get方法服务接口
func Get() {

	http.HandleFunc("/my/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.RawQuery)
		fmt.Println(writer)
		writer.Write([]byte("myGoHttp"))
	})
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println("http listen err :", err)
	}
}

//http get 请求客户端
func GetClient() {
	resp, err := http.Get("http://127.0.0.1:8888/my/go?param=hello")
	if err != nil {
		fmt.Println("get client err：", err)
		return
	}
	fmt.Println("resp:", resp)
	b := make([]byte, 1024*100)
	var body string
	for {
		read, err := resp.Body.Read(b)
		body += string(b[:read])
		if err != nil {
			if err == io.EOF || read == 0 {
				break
			}
			fmt.Println("read err:", err)
		}
	}
	fmt.Println("response body :", body)
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

//传输文件服务端
func TcpFileServer() {

	server, err := net.Listen("tcp", "127.0.0.1:8888")
	defer server.Close()
	if err != nil {
		fmt.Println("server listen err :", err)
		return
	}

	for {
		accept, err := server.Accept()
		if err != nil {
			fmt.Println("server accept err :", err)
			continue
		}
		go func(conn net.Conn) {
			defer accept.Close()
			//准备写入文件
			create, err3 := os.Create("../file/tcpFile.txt")
			defer create.Close()
			if err3 != nil {
				fmt.Println("file create err:", err3)
				return
			}
			//读取文件
			b := make([]byte, 1024*1024)
			read, err2 := conn.Read(b)
			if err2 != nil {
				fmt.Println("read err:", err2)
			}
			create.Write(b[:read])

			conn.Write([]byte("ok"))
			fmt.Println("服务端返回信息完毕")
		}(accept)
	}
}

//传输文件客户端
func TcpFileCustomer() {

	dial, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("tcp customer connection err:", err)
		return
	}
	defer dial.Close()
	open, err := os.Open("../file/writ.txt")
	defer open.Close()
	if err != nil {
		fmt.Println("file open err:", err)
		return
	}
	b := make([]byte, 1024*1024)
	for {
		read, err := open.Read(b)
		if err != nil {
			if read <= 0 || err == io.EOF {
				break
			}
			fmt.Println("file read err:", err)
			return
		}
		dial.Write(b[:read])
	}
	response := readByConn(dial)
	fmt.Println("服务端返回信息:", response)
}

//聊天室用户对象
type user struct {
	Ip       string
	UserName string
	Conn     net.Conn
}

//tcp聊天室服务端
func TcpChatServer() {
	server, err := net.Listen("tcp", "127.0.0.1:8888")
	defer server.Close()
	if err != nil {
		fmt.Println("server listen err :", err)
		return
	}

	for {
		accept, err := server.Accept()
		if err != nil {
			fmt.Println("server accept err :", err)
			continue
		}
		go func(conn net.Conn) {
			var u user
			defer conn.Close()
			defer func() {
				//广播下线通知
				ch <- u.UserName + "下线了!"
				//删除该用户
				users.Delete(u.UserName)
			}()
			for {
				b := make([]byte, 1024*1024)
				//读取客户端发过来的消息
				read, err2 := conn.Read(b)
				if err2 != nil {
					fmt.Println("read err:", err2)
					return
				}
				msg := string(b[:read])
				//如果是上线发送的第一条,就记录用户信息，并且存进用户map
				if strings.Contains(msg, "user") {
					u = user{conn.RemoteAddr().String(), msg, conn}
					users.Store(u.UserName, u)
					ch <- u.UserName + "上线了"
				} else {
					//日常对话发送到协程共享通道
					ch <- u.UserName + ":" + msg
				}
			}
		}(accept)
		go func() {
			for {
				//主协程从通道读取用户通话信息，通知所有用户
				s := <-ch
				//给所有用户发送聊天消息
				users.Range(func(key, value interface{}) bool {
					u, f := value.(user)
					if f {
						u.Conn.Write([]byte(s))
					}
					return true
				})
			}
		}()
	}
}

//线程安全存取所有用户的map
var users sync.Map
var ch = make(chan string)

//tcp聊天室客户端
func TcpChatCustomer(userName string) {
	dial, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("customer connection err :", err)
		return
	}
	defer dial.Close()

	//先进行身份认证
	_, err1 := dial.Write([]byte(userName))
	if err1 != nil {
		fmt.Println("user auth err:", err1)
		return
	}
	go func() {
		//读取服务器传过来的信息
		for {
			b := make([]byte, 1024*1024)
			_, err2 := dial.Read(b)
			if err2 != nil {
				fmt.Println("read err:", err2)
			}
			fmt.Println(string(b))
		}
	}()
	for {
		//获取用户键盘的输入
		var scan string
		fmt.Println("请输入:")
		fmt.Scan(&scan)
		//将输入信息发送给聊天服务器
		dial.Write([]byte(scan))
	}

}

//ip地址转换
func IpConvert() {

	s := "192.168.1.174"
	ip := net.ParseIP(s)
	fmt.Println(ip)
}

//udp编程
//udp服务端
func UdpServer() {
	//开启udp监听
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.ParseIP("127.0.0.1"), Port: 8000, Zone: "",
	})
	if err != nil {
		fmt.Println("udp server start err:", err)
	}
	defer udp.Close()
	//开启一个协程接收数据
	var customer *net.UDPAddr
	go func(conn *net.UDPConn) {
		b := make([]byte, 1024*100)
		//死循环阻塞读取客户端发送的数据
		for {
			n, addr, err2 := conn.ReadFromUDP(b)
			if err2 != nil {
				fmt.Println("客户端IP:", addr.IP.String(), "read err:", err2)
				continue
			}
			if customer == nil {
				customer = addr
			}

			fmt.Println("客户端IP:", addr.IP.String(), "发过来的消息:", string(b[:n]))
		}
	}(udp)

	//定时每s向客户端发送消息
	ticker := time.NewTicker(1 * time.Second)
	for {
		t := <-ticker.C
		if customer != nil {
			_, err3 := udp.WriteToUDP([]byte("发送给客户端:"+strconv.Itoa(t.Second())), customer)
			if err3 != nil {
				fmt.Println("writ err :", err3)
				continue
			}
		}
	}
}

//udp客户端
func UdpCustomer() {
	//连接服务端
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8000, Zone: ""})
	defer udp.Close()
	if err != nil {
		fmt.Println("customer connection err:", err)
		return
	}

	//给服务器发送消息
	udp.Write([]byte("客户端发送的消息"))
	//循环接收服务端的消息
	b := make([]byte, 1024*100)
	for {
		n, addr, err1 := udp.ReadFromUDP(b)
		if err1 != nil {
			fmt.Println("read err:", err1)
		}
		fmt.Println("收到服务端:"+addr.IP.String()+":"+"消息:", string(b[:n]))
	}

}
