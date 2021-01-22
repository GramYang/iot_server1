package handler

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"reflect"
)

type JsonMsg struct {
	Msg map[string]interface{}
}

// 将消息注册到系统
func init() {
	//测试消息JsonEcho
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*JsonMsg)(nil)).Elem(),
		ID:    1,
	})
}
