package main

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"iot_server1/handler"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

//const peerAddress = "106.54.87.204:8086"

const peerAddress = "127.0.0.1:8086"

func main() {
	done := make(chan struct{})
	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Connector", "client", peerAddress, queue)
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		//连接成功后发送一个消息
		case *cellnet.SessionConnected:
			fmt.Println("client connected")
			loginReq(ev)
		//收到响应后就关闭
		case *handler.JsonMsg:
			fmt.Printf("client recv %+v\n", msg)
			loginAck(ev,msg.Msg)
			done <- struct{}{}
		case *cellnet.SessionClosed:
			fmt.Println("client closed")
		}
	})
	peerIns.Start()
	queue.StartLoop()
	<-done
}

func loginReq(ev cellnet.Event){
	msg:=map[string]interface{}{}
	msg["Action"]="GetDeviceLogin"
	msg["DeviceName"]="测试119"
	data,_:=json.Marshal(msg)
	m:=handler.JsonMsg{Msg:string(data)}
	ev.Session().Send(m)
}

func loginAck(ev cellnet.Event,msg string){
	resMap:=map[string]interface{}{}
	_=json.Unmarshal([]byte(msg),&resMap)
	resMap1:=map[string]interface{}{}
	_=json.Unmarshal([]byte(resMap["Props"].(string)),&resMap1)
	resMap2:=map[string]interface{}{}
	_=json.Unmarshal([]byte(resMap1["Login"].(string)),&resMap2)
	fmt.Println("parse json to map",resMap2)
}