package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service"
)

func main() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	res, err := service.DeployNginx(h)
	fmt.Println(res, err)
}
