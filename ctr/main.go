package main

import (
	"github.com/ahojcn/ecloud/ctr/router"
	"github.com/ahojcn/ecloud/ctr/util"
)

func main() {
	r := router.SetupRouter()

	port := util.Config.Section("system").Key("listen_port").String()
	_ = r.Run(":" + port)
}
