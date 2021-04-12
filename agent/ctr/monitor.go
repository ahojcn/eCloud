package ctr

import (
	"fmt"
	"github.com/ahojcn/ecloud/agent/util"
	"github.com/parnurzeal/gorequest"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"time"
)

func RunMetrics() {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		go CpuMetrics()
	})
	c.Start()
}

type reqType struct {
	HostId  int64       `json:"host_id"`
	Metrics string      `json:"metrics"`
	Data    interface{} `json:"data"`
}

type cpuMetrics struct {
	User    float64 `json:"user"`
	System  float64 `json:"system"`
	Idle    float64 `json:"idle"`
	Percent float64 `json:"percent"`
}

func CpuMetrics() {
	log.Infoln("cpu metrics to ctr")
	cpu := new(util.CpuInfoMonitor)
	cpu.Get()
	MonitorMetricsWrite("cpu", cpuMetrics{
		User:    cpu.Times.User,
		System:  cpu.Times.System,
		Idle:    cpu.Times.Idle,
		Percent: cpu.Percent,
	})
}

func MonitorMetricsWrite(metrics string, data interface{}) {
	targetUrl := fmt.Sprintf("http://%s:%v/metrics", util.Config.Ctr.IP, util.Config.Ctr.Port)

	d := reqType{
		HostId:  util.Config.Ctr.HostId,
		Metrics: metrics,
		Data:    data,
	}

	_, _, err := gorequest.New().Post(targetUrl).Send(d).Retry(3, 3*time.Second, 500).EndBytes()
	if err != nil {
		log.Errorln("metric failed", metrics, data)
	}
}
