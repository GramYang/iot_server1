package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"iot_server1/config"
	"iot_server1/handler"
	"iot_server1/log"
	sc "iot_server1/sqlx_client"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "iot_server1/proc/tcp"
)

//iot_server1是基于cellnet的短连接服务器
func main(){
	config.Setup()
	if config.Conf.LocalDebug{
		log.InitLog(1)
	}else{
		log.InitLog(0)
		log.Debug(true)
	}
	sc.SetUp()
	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Acceptor", "iot_server1", config.Conf.ServerPort, queue)
	//tcp.iotltv仅仅只是在ltv的基础上添加了字节数组的log，方便调试，其他不变
	proc.BindProcessorHandler(peerIns, "tcp.iotltv", handler.Handler)
	peerIns.Start()
	queue.StartLoop()
	queue.Wait()
}