package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
)

func main() {
	h := new(model.Host)
	h.Description = "测试主机"
	//fmt.Println(model.HostAdd(h))

	orm := model.GetMaster()
	fmt.Println(orm.ID(1).Update(h))
}
