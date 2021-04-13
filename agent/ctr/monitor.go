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
	spec := "@every 10s"
	c.AddFunc(spec, CpuMetrics)
	c.AddFunc(spec, MemMetrics)
	c.AddFunc(spec, DiskMetrics)
	c.AddFunc(spec, LoadMetrics)
	c.AddFunc(spec, NetMetrics)
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

type memMetrics struct {
	SwapTotal        uint64  `json:"swap_total"`
	SwapUsed         uint64  `json:"swap_used"`
	SwapFree         uint64  `json:"swap_free"`
	SwapPercent      float64 `json:"swap_percent"`
	VirtualTotal     uint64  `json:"virtual_total"`
	VirtualAvailable uint64  `json:"virtual_available"`
	VirtualUsed      uint64  `json:"virtual_used"`
	VirtualPercent   float64 `json:"virtual_percent"`
	VirtualFree      uint64  `json:"virtual_free"`
	VirtualActive    uint64  `json:"virtual_active"`
	VirtualInactive  uint64  `json:"virtual_inactive"`
	VirtualWired     uint64  `json:"virtual_wired"`
	VirtualBuffers   uint64  `json:"virtual_buffers"`
	VirtualCached    uint64  `json:"virtual_cached"`
}

type diskMetrics struct {
	Total         uint64  `json:"total"`
	Free          uint64  `json:"free"`
	Used          uint64  `json:"used"`
	Percent       float64 `json:"percent"`
	InodesTotal   uint64  `json:"inodes_total"`
	InodesUsed    uint64  `json:"inodes_used"`
	InodesFree    uint64  `json:"inodes_free"`
	InodesPercent float64 `json:"inodes_percent"`
}

type loadMetrics struct {
	Load1        float64 `json:"load_1"`
	Load5        float64 `json:"load_5"`
	Load15       float64 `json:"load_15"`
	ProcessCount int `json:"process_count"`
}

type netMetrics struct {
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
	ErrIn       uint64 `json:"err_in"`
	ErrOut      uint64 `json:"err_out"`
	DropIn      uint64 `json:"drop_in"`
	DropOut     uint64 `json:"drop_out"`
	FifoIn      uint64 `json:"fifo_in"`
	FifoOut     uint64 `json:"fifo_out"`
}

func CpuMetrics() {
	cpu := new(util.CpuInfoMonitor)
	cpu.Get()
	MonitorMetricsWrite("cpu", cpuMetrics{
		User:    cpu.Times.User,
		System:  cpu.Times.System,
		Idle:    cpu.Times.Idle,
		Percent: cpu.Percent,
	})
}

func MemMetrics() {
	mem := new(util.MemInfoMonitor)
	mem.Get()
	MonitorMetricsWrite("mem", memMetrics{
		SwapTotal:        mem.SwapMemoryStat.Total,
		SwapUsed:         mem.SwapMemoryStat.Used,
		SwapFree:         mem.SwapMemoryStat.Free,
		SwapPercent:      mem.SwapMemoryStat.UsedPercent,
		VirtualTotal:     mem.VirtualMemoryStat.Total,
		VirtualAvailable: mem.VirtualMemoryStat.Available,
		VirtualUsed:      mem.VirtualMemoryStat.Used,
		VirtualPercent:   mem.VirtualMemoryStat.UsedPercent,
		VirtualFree:      mem.VirtualMemoryStat.Free,
		VirtualActive:    mem.VirtualMemoryStat.Available,
		VirtualInactive:  mem.VirtualMemoryStat.Inactive,
		VirtualWired:     mem.VirtualMemoryStat.Wired,
		VirtualBuffers:   mem.VirtualMemoryStat.Buffers,
		VirtualCached:    mem.VirtualMemoryStat.Cached,
	})
}

func DiskMetrics() {
	d := new(util.DiskInfoMonitor)
	d.Get()
	MonitorMetricsWrite("disk", diskMetrics{
		Total:         d.Usage.Total,
		Free:          d.Usage.Free,
		Used:          d.Usage.Used,
		Percent:       d.Usage.UsedPercent,
		InodesTotal:   d.Usage.InodesTotal,
		InodesUsed:    d.Usage.InodesUsed,
		InodesFree:    d.Usage.InodesFree,
		InodesPercent: d.Usage.InodesUsedPercent,
	})
}

func LoadMetrics() {
	l := new(util.LoadInfoMonitor)
	p := new(util.ProcessMonitor)
	l.Get()
	p.Get()
	MonitorMetricsWrite("load", loadMetrics{
		Load1:  l.Avg.Load1,
		Load5:  l.Avg.Load5,
		Load15: l.Avg.Load15,
		ProcessCount: len(p.Process),
	})
}

func NetMetrics() {
	n := new(util.NetInfoMonitor)
	n.Get()
	MonitorMetricsWrite("net", netMetrics{
		BytesSent:   n.IOCounters.BytesSent,
		BytesRecv:   n.IOCounters.BytesRecv,
		PacketsSent: n.IOCounters.PacketsSent,
		PacketsRecv: n.IOCounters.PacketsRecv,
		ErrIn:       n.IOCounters.Errin,
		ErrOut:      n.IOCounters.Errout,
		DropIn:      n.IOCounters.Dropin,
		DropOut:     n.IOCounters.Dropout,
		FifoIn:      n.IOCounters.Fifoin,
		FifoOut:     n.IOCounters.Fifoout,
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
