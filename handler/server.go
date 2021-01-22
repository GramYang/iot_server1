package handler

import (
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
		jsonHandler(ev, msg.Msg)
	case *cellnet.SessionClosed:
		g.Debugln("session closed: ", ev.Session().ID())
	}
}

func jsonHandler(ev cellnet.Event, msg map[string]interface{}) {
	if msg["Action"] != nil && msg["Action"] == "GetDeviceLogin" {
		if msg["DeviceName"] != nil && msg["DeviceName"] != "" {
			Login(ev, msg["DeviceName"].(string))
		} else {
			g.Errorln("JsonMsg missing DeviceName")
		}
	} else {
		g.Errorln("JsonMsg missing Action")
	}
}
