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
	fshost   = flag.String("fshost", "localhost", "Freeswitch hostname. Default: localhost")
	fsport   = flag.Uint("fsport", 8021, "Freeswitch port. Default: 8021")
	password = flag.String("pass", "ClueCon", "Freeswitch password. Default: ClueCon")
	timeout  = flag.Int("timeout", 10, "Freeswitch conneciton timeout in seconds. Default: 10")
)

func InboundClient() (cli *goesl.Client, err error) {

	client, err := goesl.NewClient(*fshost, *fsport, *password, *timeout)

	if err != nil {
		fmt.Println("连接异常:", err)
		return nil, err
	}
	go client.Handle()
	return client, nil
}

func SendEvents() {
	client, err := InboundClient()
	if err != nil {
		fmt.Println(err)
	}
	err2 := client.Send("events json CHANNEL_CREATE CHANNEL_ANWSER HEARTBEAT")
	if err != nil {
		fmt.Println("send err:", err2)
	}
	Receive(client)
}

func ExecuteCommand() {
	client, err := InboundClient()
	if err != nil {
		fmt.Println(err)
	}
	client.Api("sofia status profile internal reg")
	Receive(client)
}

func Receive(client *goesl.Client) {
	for {
		message, err3 := client.ReadMessage()
		if err3 != nil {
			fmt.Println(err3)
		}
		fmt.Println(string(message.Body))
	}
}
