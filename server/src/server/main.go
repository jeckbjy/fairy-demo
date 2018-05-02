package main

import (
	"server/comm/app"
	"server/comm/conf"
	"server/zcenter"
	"server/zgate"
	"server/zhall"
	"server/zservice"

	"github.com/jeckbjy/fairy"
)

func newServer() app.IApp {
	switch conf.Mode {
	case conf.ModeCenter:
		return zcenter.New()
	case conf.ModeGate:
		return zgate.New()
	case conf.ModeHall:
		return zhall.New()
	case conf.ModeService:
		return zservice.New()
	default:
		// logging.Debugf()
		return nil
	}
}

func main() {
	srv := newServer()

	srv.Start()
	srv.Run()

	fairy.WaitExit()
	srv.Stop()
}
