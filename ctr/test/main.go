package main

import (
	"fmt"
	"strings"

	"github.com/ahojcn/ecloud/ctr/model"
)

func main() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	res, err := h.RunCmd("ls /root/.eCloud/nginx/conf/", 0)
	resArr := strings.Split(res, "\n")
	fmt.Println(resArr[len(resArr)-2], err)
}
