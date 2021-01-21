package handler

import (
	"encoding/json"
	g "github.com/GramYang/gylog"
	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/codec/json"
)

//此处的所有业务回调都是短连接
func Handler(ev cellnet.Event) {
	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted:
		g.Debugln("server accepted")
	case *JsonMsg:
		g.Debugln("server recv", msg)
		jsonHandler(ev,msg.Msg)
	case *cellnet.SessionClosed:
		g.Debugln("session closed: ", ev.Session().ID())
	}
}

func jsonHandler(ev cellnet.Event,msg string){
	resMap:=map[string]interface{}{}
	err:=json.Unmarshal([]byte(msg),&resMap)
	if err!=nil{
		g.Errorln("JsonMsg Unmarshal:",err.Error())
		return
	}
	if resMap["Action"]!=nil&&resMap["Action"]=="GetDeviceLogin"{
		if resMap["DeviceName"]!=nil&&resMap["DeviceName"]!=""{
			Login(ev,resMap["DeviceName"].(string))
		}else{
			g.Errorln("JsonMsg missing DeviceName")
		}
	}else{
		g.Errorln("JsonMsg missing Action")
	}
}