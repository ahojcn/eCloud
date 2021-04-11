package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
	"strings"
	"time"
)

func main() {
	h, _ := model.HostOne(map[string]interface{}{"id": 7})
	res, _ := h.RunCmd("docker images", time.Second * 60)
	sa := strings.Split(res, "\n")[1:]
	sa = sa[:len(sa) - 1]
	for _, s := range sa {
		fmt.Println(s, "---")
	}
}
