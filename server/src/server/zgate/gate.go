package zgate

import "server/comm/server"

func New() *Gate {
	srv := &Gate{}
	return srv
}

type Gate struct {
	server.Server
}
