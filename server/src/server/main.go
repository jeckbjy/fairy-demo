package main

import (
	"server/center"

	"github.com/jeckbjy/fairy"
)

func main() {
	switch cfg.Mode {
	case cfg.ModeCenter:
		center.Start()
		break
	case cfg.ModeGate:
		break
	case cfg.ModeHall:
		break
	}

	fairy.WaitExit()
}
