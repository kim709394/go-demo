package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

/**
 *@Author kim
 *@Description
 *@Date 2021/5/6 15:04
 **/

func TestWs(t *testing.T) {

	s := g.Server()
	s.BindHandler("/ws", func(r *ghttp.Request) {
		//获取websocket对象
		ws, err := r.WebSocket()
		if err != nil {
			g.Dump(err)
		}
		//向客户端发送消息
		ws.WriteMessage(1, gconv.Bytes("哈哈"))
		//监听获取客户端发过来的消息
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				g.Dump(err)
			}
			g.Dump("消息类型:", msgType)
			g.Dump("消息:", string(msg))

		}

	})
	s.SetPort(8089)
	s.Run()
}
