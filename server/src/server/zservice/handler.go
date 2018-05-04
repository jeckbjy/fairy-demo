package zservice

import (
	"server/comm/msg"

	"github.com/jeckbjy/fairy"
)

func init() {
	addHandler(msg.MsgId_HandlerDictReq, onHandlerDict)
}

func addHandler(msgid int, cb fairy.HandlerCB) {
	fairy.RegisterHandler(msgid, cb)
}

func onHandlerDict(ctx *fairy.HandlerCtx) {
	// req := ctx.GetMessage().(*msg.HandlerDictReq)
	rsp := &msg.HandlerDictRsp{}
	rsp.ServerId = ""
	// rsp.Msgid
	ctx.Send(rsp)
}
