package fscli

import (
	"github.com/kim709394/go-demo/freeswitchcli"
	"testing"
)

/*
@Author kim
@Description   内连测试
@date 2021-1-5 9:42
*/

//初始化连接
func TestInbound(t *testing.T) {
	freeswitchcli.InboundClient()
}

//订阅事件监听
func TestSendEvents(t *testing.T) {
	freeswitchcli.SendEvents()
}

//执行命令
func TestExecuteCommand(t *testing.T) {
	//cmd := "sofia status profile internal reg"
	cmd := "show modules mod_portaudio1"
	freeswitchcli.ExecuteCommand(cmd)
}

//接收控制台消息
func TestReceive(t *testing.T) {
	freeswitchcli.InboundClient()

	freeswitchcli.Receive()

}
