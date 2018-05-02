package zcenter

import (
	"server/comm/server"
)

func New() *Center {
	srv := &Center{}
	return srv
}

type Center struct {
	server.Server
}

func (srv *Center) Start() {
	// srv.server = soa.NewServer()
	// srv.server.Start(conf.Conf.Mode)
}

func (srv *Center) Run() {

}

func (srv *Center) Stop() {

}
