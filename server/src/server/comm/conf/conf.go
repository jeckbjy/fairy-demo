package conf

const (
	ModeCenter  = "center"
	ModeGate    = "gate"
	ModeHall    = "hall"
	ModeService = "service"
)

const (
	EnvProd = "prod"
)

var Conf struct {
	Mode    string
	Env     string
	Port    int
	Debug   bool
	Version string
	WorkDir string
}

var Mode string
var Env string
var Port int
var Debug bool
var Version string
var WorkDir string

func GetPath(relativePath string) string {
	return WorkDir + relativePath
}
