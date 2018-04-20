package conf

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jeckbjy/fairy/log"
	yaml "gopkg.in/yaml.v2"
)

func init() {
	// first parse option
	mode := flag.String("m", "", "game mode")
	env := flag.String("e", "local", "env")
	port := flag.Int("p", 0, "port")
	debug := flag.Bool("d", false, "debug")
	version := flag.String("v", "", "version")
	cwd := flag.String("w", "", "work dir")

	flag.Parse()

	Mode = parseMode(*mode)
	Env = *env
	Port = *port
	Debug = *debug
	Version = *version
	WorkDir = *cwd

	if WorkDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("get exec path fail:%+v", err))
		}
		WorkDir = wd
	}

	if !strings.HasSuffix(WorkDir, "/") {
		WorkDir += "/"
	}

	// load cfg
	loadConfig("server.yaml")
}

func parseMode(s string) ModeType {
	switch s {
	case "center":
		return ModeCenter
	case "gate":
		return ModeGate
	case "hall":
		return ModeHall
	}

	return ModeNone
}

func loadConfig(name string) {
	file := GetPath(name)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Debug("load yaml conf fail,%+v", err)
		return
	}

	root := yaml.MapSlice{}
	if err := yaml.Unmarshal([]byte(data), &root); err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\n", root)
	}
}
