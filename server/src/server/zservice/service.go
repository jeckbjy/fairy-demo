package zservice

import (
	"server/comm/server"
)

func New() *Service {
	srv := &Service{}
	return srv
}

type Service struct {
	server.Server
}
