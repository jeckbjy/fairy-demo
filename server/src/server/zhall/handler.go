package zhall

import (
	"server/comm/msg"

	"github.com/jeckbjy/fairy"
)

var gServiceHandlerMap map[uint]string

// 注册回调函数
func init() {
	fairy.RegisterUncaughtHandler(onUncaught)
	addHandler(msg.MsgId_HandlerDictRsp, onHandlerDic)
	addHandler(msg.MsgId_AuthReq, onAuth)
	addHandler(msg.MsgId_LoginReq, onLogin)
}

func addHandler(msgid int, cb fairy.HandlerCB) {
	fairy.RegisterHandler(msgid, cb)
}

func onUncaught(ctx *fairy.HandlerCtx) {
	// auto transfer to service
	msgid := ctx.GetId()
	if _, ok := gServiceHandlerMap[msgid]; ok {
		// GetServer().Send()
	} else {
		// logging
		// log.Debug("cannot find handler:%+v", msgid)
	}
}

// 获取service关心的handler
func onHandlerDic(ctx *fairy.HandlerCtx) {
	rsp := ctx.GetMessage().(*msg.HandlerDictRsp)
	for _, msgid := range rsp.Msgid {
		if _, ok := gServiceHandlerMap[uint(msgid)]; !ok {
			gServiceHandlerMap[uint(msgid)] = rsp.ServerId
		}
	}
}

func onAuth(ctx *fairy.HandlerCtx) {
	req := ctx.GetMessage().(*msg.AuthReq)
	rsp := &msg.AuthRsp{}

	if req.Account != "test" {
		rsp.Code = msg.Error_ErrBadAccount
		ctx.Send(rsp)
		return
	}

	if req.Password != "test" {
		rsp.Code = msg.Error_ErrBadPassword
		ctx.Send(rsp)
		return
	}

	rsp.Code = msg.Error_ErrOk
	ctx.Send(rsp)
}

func onLogin(ctx *fairy.HandlerCtx) {
	// rsp := &msg.LoginRsp{}
	// ctx.Send(rsp)
}
