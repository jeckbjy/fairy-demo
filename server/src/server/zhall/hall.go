package zhall

import (
	"server/comm/app"
	"server/comm/server"
)

func New() app.IApp {
	srv := &Hall{}
	return srv
}

type Hall struct {
	server.Server
}

func (srv *Hall) Start() {

}

func (srv *Hall) Stop() {

}
