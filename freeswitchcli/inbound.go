package freeswitchcli

import (
	"flag"
	"fmt"
	"github.com/0x19/goesl"
)

/*
@Author kim
@Description       freeswitch goland esl 内连
@date 2021-1-5 9:07
*/

var (
	//设置ip
	fshost = flag.String("fshost", "localhost", "Freeswitch hostname. Default: localhost")
	//设置端口号 默认是8021
	fsport = flag.Uint("fsport", 8021, "Freeswitch port. Default: 8021")
	//设置密码  默认是 ClueCon
	password = flag.String("pass", "ClueCon", "Freeswitch password. Default: ClueCon")
	//设置连接超时时间
	timeout = flag.Int("timeout", 10, "Freeswitch conneciton timeout in seconds. Default: 10")

	cli *goesl.Client
)

//连接freeswitch服务器
func InboundClient() error {

	client, err := goesl.NewClient(*fshost, *fsport, *password, *timeout)

	if err != nil {
		fmt.Println("连接异常:", err)
		return err
	}
	go client.Handle()
	cli = client
	return nil
}

//订阅事件
func SendEvents() {
	//InboundClient()
	err := cli.Send("events json CHANNEL_CREATE channel_hangup_complete CHANNEL_ANWSER custom portaudio::ringing")
	if err != nil {
		fmt.Println("send err:", err)
	}
	/*b := make([]byte, 1024*1024)
	//读取客户端发过来的消息
	read, err2 := cli.SocketConnection.Conn.Read(b)
	acceptMsg := string(b[:read])
	if err2 != nil {
		fmt.Println("read err:", err2)
	}
	fmt.Println(acceptMsg)*/
	//Receive()
}

//执行命令
func ExecuteCommand(cmd string) {
	//InboundClient()

	cli.Api(cmd)
	//b := make([]byte, 1024*1024)
	//读取客户端发过来的消息
	/*read, err2 := cli.Read(b)
	acceptMsg := string(b[:read])
	if err2 != nil {
		fmt.Println("read err:", err2)
	}
	fmt.Println("返回消息:", acceptMsg)*/

	//Receive()
}

//接收控制台消息
func Receive() {
	for {
		message, err3 := cli.ReadMessage()
		if err3 != nil {
			fmt.Println(err3)
		}

		fmt.Println("接收消息:", message.GetHeader("Event-Name"))
	}
}
