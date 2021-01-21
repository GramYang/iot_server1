package handler

import (
	"encoding/json"
	g "github.com/GramYang/gylog"
	"github.com/davyxu/cellnet"
	sc "iot_server1/sqlx_client"
)

//一个deviceName必须只标定一个设备
func Login(ev cellnet.Event, deviceName string){
	resMap1:=map[string]interface{}{}
	res,err:=sc.LoginByDeviceName(deviceName)
	g.Debugln("mysql result",res)
	if err!=nil{
		g.Errorln(err)
		resMap1["Success"]=false
		data,_:=json.Marshal(resMap1)
		ev.Session().Send(&JsonMsg{Msg: string(data)})
		ev.Session().Close()
		return
	}
	resMap1["Success"]=true
	data,_:=json.Marshal(res)
	resMap2:=map[string]interface{}{}
	resMap2["Login"]=string(data)
	data1,_:=json.Marshal(resMap2)
	resMap1["Props"]=string(data1)
	data2,_:=json.Marshal(resMap1)
	g.Debugln("send msg",string(data2))
	ev.Session().Send(&JsonMsg{Msg:string(data2)})
	ev.Session().Close()
}