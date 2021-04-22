package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
	"time"
)

func main() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	res, err := h.RunCmd("lsof -i:22", time.Second * 60)
	fmt.Println(res, "res")
	fmt.Println(err, "err")
}
