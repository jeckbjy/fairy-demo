package discovery

import (
	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy/codec"
	"github.com/jeckbjy/fairy/filter"
	"github.com/jeckbjy/fairy/frame"
	"github.com/jeckbjy/fairy/identity"
	"github.com/jeckbjy/fairy/tcp"
	"github.com/jeckbjy/fairy/util"
)

// InfoGroup 按组存储服务器
type InfoGroup struct {
	Type  uint    // 服务器类型
	Index uint    // 轮训当前索引
	Infos []*Info // 服务器信息
}

// GroupMap server_type->infos
type GroupMap map[uint]*InfoGroup

// Subscribe 自己订阅服务器信息
type Subscribe struct {
	Infos  InfoMap  // 按ID映射
	Groups GroupMap // 按类型分组
}

// Add 添加一条信息
func (sub *Subscribe) Add(info *Info) bool {
	// 已经存在了
	if _, ok := sub.Infos[info.ID]; ok {
		return false
	}

	group := sub.Groups[info.Type]
	if group == nil {
		group = &InfoGroup{Type: info.Type, Index: 0}
	}

	group.Infos = append(group.Infos, info)

	sub.Infos[info.ID] = info

	return true
}

// Remove 删除一条信息
func (sub *Subscribe) Remove(id uint64) bool {
	info, ok := sub.Infos[id]
	if !ok {
		return false
	}

	delete(sub.Infos, id)

	// 通过类型删除
	group := sub.Groups[info.Type]
	if group == nil {
		return false
	}

	for idx, temp := range group.Infos {
		if info == temp {
			group.Infos = append(group.Infos[:idx], group.Infos[idx+1:]...)
			break
		}
	}

	if len(group.Infos) == 0 {
		delete(sub.Groups, info.Type)
	}

	return true
}

// Discovery 注册自身信息，并获取自己关注的服务器信息
type Discovery struct {
	id      uint64     // 唯一ID
	tran    fairy.Tran //
	conn    fairy.Conn //
	info    *Info      // 自身信息
	moniter uint       // 关注的
	sub     Subscribe  // 订阅的信息
}

// Connect 连接注册中心,并注册自身信息
func (dis *Discovery) Connect(host string) error {
	// 初始化信息
	if dis.id == 0 {
		id, err := util.NextID()
		if err != nil {
			return err
		}

		dis.id = id
	}

	// 自动初始化内网IP
	if dis.info.InnerAddr == "" {
		addr, err := util.GetIPv4()
		if err != nil {
			return err
		}
		dis.info.InnerAddr = addr
	}

	// 注册回调

	// 建立连接
	tran := tcp.NewTran()
	tran.AddFilters(
		filter.NewFrame(frame.NewVarintLength()),
		filter.NewPacket(identity.NewString(), codec.NewGob()),
		filter.NewExecutor())

	tran.Connect(host, 0)
	tran.Start()

	dis.tran = tran
	return nil
}

func (dis *Discovery) Stop() {
	if dis.tran != nil {
		dis.tran.Stop()
	}
}

// 连接成功,注册自身信息
func (dis *Discovery) onConnect(conn fairy.Conn) {

}

// 注册返回
func (dis *Discovery) onRegister(ctx *fairy.HandlerCtx) {
	rsp := ctx.GetMessage().(*zSDRegisterRsp)
	for _, info := range rsp.Infos {
		// add info
		dis.sub.Add(info)
		// auto connect?
	}
}

// 删除节点
func (dis *Discovery) onRemove(ctx *fairy.HandlerCtx) {
	rsp := ctx.GetMessage().(*zSDRemoveRsp)
	dis.sub.Remove(rsp.ID)
}
