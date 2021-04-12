package ctr

import (
	"encoding/json"
	"fmt"
	"github.com/ahojcn/ecloud/agent/util"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"time"
)

type HostExtra struct {
	CpuInfo    util.CpuInfo    `json:"cpu_info"`
	DiskInfo   util.DiskInfo   `json:"disk_info"`
	DockerInfo util.DockerInfo `json:"docker_info"`
	HostInfo   util.HostInfo   `json:"host_info"`
}

func (h *HostExtra) Get() {
	_ = h.HostInfo.Get()
	_ = h.DockerInfo.Get()
	_ = h.DiskInfo.Get()
	_ = h.CpuInfo.Get()
}

func (h *HostExtra) String() string {
	bs, _ := json.Marshal(h)
	return string(bs)
}

func ReportHostExtra() {
	targetUrl := fmt.Sprintf("http://%s:%s/host", util.Config.Ctr.IP, util.Config.Ctr.Port)
	hostExtra := new(HostExtra)
	hostExtra.Get()
	type data struct {
		HostId int64  `json:"host_id"`
		Extra  string `json:"extra"`
	}
	d := data{
		HostId: util.Config.Ctr.HostId,
		Extra:  hostExtra.String(),
	}
	_, _, err := gorequest.New().Put(targetUrl).Retry(3, 3*time.Second, 500).
		Send(d).EndBytes()
	if err != nil {
		log.Errorln("report host extra failed! err :", err)
		return
	}
}
