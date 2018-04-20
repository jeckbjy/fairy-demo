package main

import (
	"server/center"
	"server/comm/conf"

	"github.com/jeckbjy/fairy"
)

func main() {
	switch conf.Mode {
	case conf.ModeCenter:
		center.Start()
		break
	case conf.ModeGate:
		break
	case conf.ModeHall:
		break
	}

	fairy.WaitExit()
}
