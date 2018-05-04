package discovery

// 注册自身信息和关注的信息
type zSDRegisterReq struct {
	Info *Info
}

type zSDRegisterRsp struct {
	Infos []*Info
}

// 通知删除
type zSDRemoveRsp struct {
	ID uint64
}
