package project_name

import (
	"github.com/author_name/project_name/configs"
)

type Arguments struct {
	DelayOpt  int
	OutputOpt string
	EnvOpt    string
	ProxyOpt  string
}

func Run(args *Arguments) {
	configs.InitHttpClient(args.ProxyOpt)
}
