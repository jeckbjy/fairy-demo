package conf

type ModeType int

const (
	ModeNone ModeType = iota
	ModeCenter
	ModeGate
	ModeHall
)

const (
	EnvProd = "prod"
)

var Mode ModeType
var Env string
var Port int
var Debug bool
var Version string
var WorkDir string

func GetPath(relativePath string) string {
	return WorkDir + relativePath
}
