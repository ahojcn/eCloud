package util

import (
	"os"
	"strconv"
)

var Config *config

type config struct {
	Ctr ctrConfig `json:"ctr"`
}

type ctrConfig struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	HostId int64  `json:"host_id"`
}

func init() {
	id, _ := strconv.Atoi(os.Getenv("ECLOUD_AGENT_HOSTID"))
	ctrConf := ctrConfig{
		IP:     os.Getenv("ECLOUD_CTR_IP"),
		Port:   os.Getenv("ECLOUD_CTR_PORT"),
		HostId: int64(id),
	}
	Config = new(config)
	Config.Ctr = ctrConf
}
