package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
	"time"
)

func main() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	res, _ := h.RunCmd("cd /root/.eCloud && kill -9 `cat agent.pid`", time.Second * 60)
	fmt.Println(res)
}
