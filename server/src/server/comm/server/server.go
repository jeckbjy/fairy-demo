package server

import (
	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy/filter"
	"github.com/jeckbjy/fairy/frame"
	"github.com/jeckbjy/fairy/identity"
	"github.com/jeckbjy/fairy/plugins/pbcodec"
	"github.com/jeckbjy/fairy/tcp"
)

type Server struct {
	tran fairy.Tran
}

func (srv *Server) Start() {
	tran := tcp.NewTran()
	tran.AddFilters(
		filter.NewLogging(),
		filter.NewFrame(frame.NewVarintLength()),
		filter.NewPacket(identity.NewServer(), pbcodec.New()),
		filter.NewExecutor())

	srv.tran = tran
	// srv.SubTran(tran)
}

func (srv *Server) Run() {
	// srv.Client.Start(":8080")
	// srv.tran.Listen(":8080", nil)
}

func (srv *Server) Stop() {

}
