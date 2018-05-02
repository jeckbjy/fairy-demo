package zhall

import (
	"server/comm/msg"

	"github.com/jeckbjy/fairy"
)

// 注册回调函数
func init() {
	fairy.RegisterUncaughtHandler(onUncaught)
	addHandler(msg.MsgId_AuthReq, onAuth)
	addHandler(msg.MsgId_LoginReq, onLogin)
}

func addHandler(msgid int, cb fairy.HandlerCB) {
	fairy.RegisterHandler(msgid, cb)
}

func onUncaught(ctx *fairy.HandlerCtx) {
	// auto transfer to service
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
