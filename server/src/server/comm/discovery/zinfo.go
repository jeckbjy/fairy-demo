package discovery

// Info 服务器注册信息
type Info struct {
	ID         uint64 // 唯一ID
	Type       uint   // 服务器类型枚举
	Services   uint   // 提供的服务类型
	InnerAddr  string // 内网地址
	OuterAddr  string // 外网地址
	Port       uint   // 端口
	Subscribes uint   // 订阅的服务器类型,按位索引
}

func (info *Info) isSubscribe(serverType uint) bool {
	return (info.Subscribes & (1 << serverType)) != 0
}

// InfoMap id->info
type InfoMap map[uint64]*Info
