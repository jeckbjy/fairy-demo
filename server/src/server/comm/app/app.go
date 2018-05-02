package app

// IApp 程序入口
type IApp interface {
	Start()
	Run()
	Stop()
}
