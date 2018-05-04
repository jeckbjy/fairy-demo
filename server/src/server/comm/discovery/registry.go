package discovery

import (
	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy/codec"
	"github.com/jeckbjy/fairy/filter"
	"github.com/jeckbjy/fairy/frame"
	"github.com/jeckbjy/fairy/identity"
	"github.com/jeckbjy/fairy/tcp"
)

// Registry server-side
type Registry struct {
	tran  fairy.Tran
	infos InfoMap
}

// Listen 监听服务
func (reg *Registry) Listen(host string) {
	// 注册回调函数
	fairy.RegisterHandler(&zSDRegisterReq{}, reg.onRegister)

	// 创建tran
	tran := tcp.NewTran()
	tran.AddFilters(
		filter.NewFrame(frame.NewVarintLength()),
		filter.NewPacket(identity.NewString(), codec.NewGob()),
		filter.NewExecutor(),
		filter.NewClose(reg.onClose))

	tran.Listen(host, 0)
	tran.Start()

	reg.tran = tran
}

func (reg *Registry) Stop() {
	if reg.tran != nil {
		reg.tran.Stop()
	}
}

// 节点中断,删除并通知
func (reg *Registry) onClose(conn fairy.Conn) {
	uid := conn.GetUid()
	if uid == 0 {
		return
	}

	if info, ok := reg.infos[uid]; ok {
		// remove
		delete(reg.infos, uid)
		// notify
		rsp := &zSDRemoveRsp{}
		rsp.ID = info.ID
		for _, si := range reg.infos {
			if si.isSubscribe(info.Type) {
				// send msg
			}
		}
	}

}

// 有新的客户端注册
func (reg *Registry) onRegister(ctx *fairy.HandlerCtx) {
	req := ctx.GetMessage().(*zSDRegisterReq)
	info := req.Info

	ctx.SetUid(info.ID)

	reg.infos[info.ID] = info

	// 通知自己关心的信息
	if info.Subscribes != 0 {
		rsp := &zSDRegisterRsp{}
		for _, sinfo := range reg.infos {
			if sinfo.ID == info.ID {
				continue
			}

			if info.isSubscribe(sinfo.Type) {
				rsp.Infos = append(rsp.Infos, sinfo)
			}
		}

		ctx.Send(rsp)
	}

	// 通知他人关心自己的信息
	rsp := &zSDRegisterRsp{}
	rsp.Infos = append(rsp.Infos, info)
	for _, sinfo := range reg.infos {
		if sinfo.ID == info.ID {
			continue
		}

		if sinfo.isSubscribe(info.Type) {
			// send
		}
	}
}
