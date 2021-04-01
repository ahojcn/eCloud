package util

import (
	"github.com/go-ini/ini"
	"os"
)

var RootPath string
var Config *ini.File

func init() {
	var err error

	err = os.Setenv("ECLOUD_ROOT_PATH", "/Users/ahojcn/go/src/github.com/ahojcn/ecloud/ctr")
	if err != nil {
		panic(err)
	}

	RootPath = os.Getenv("ECLOUD_ROOT_PATH")
	if RootPath == "" {
		panic("root path is '', please set ECLOUD_ROOT_PATH!")
		return
	}

	Config, err = ini.Load(RootPath + "/conf/config.ini")
	if err != nil {
		panic(err)
	}
}
