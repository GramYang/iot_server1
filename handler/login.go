package handler

import (
	g "github.com/GramYang/gylog"
	"github.com/davyxu/cellnet"
	sc "iot_server1/sqlx_client"
)

//一个deviceName必须只标定一个设备
func Login(ev cellnet.Event, deviceName string) {
	resMap1 := map[string]interface{}{}
	res, err := sc.LoginByDeviceName(deviceName)
	if err != nil {
		g.Errorln(err)
		resMap1["Success"] = false
		ev.Session().Send(&JsonMsg{Msg: resMap1})
		ev.Session().Close()
		return
	}
	resMap1["Success"] = true
	resMap2 := map[string]interface{}{}
	resMap2["Login"] = res
	resMap1["Props"] = resMap2
	ev.Session().Send(&JsonMsg{Msg: resMap1})
	ev.Session().Close()
}
