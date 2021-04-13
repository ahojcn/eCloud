package util

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/docker"
)

type PsInfo interface {
	Get() error
}

type CpuInfo struct {
	Physical int            `json:"physical"`
	Logical  int            `json:"logical"`
	Info     []cpu.InfoStat `json:"info"`
}

func (c *CpuInfo) Get() error {
	var err error
	c.Physical, err = cpu.Counts(false)
	if err != nil {
		c.Physical = -1
	}
	c.Logical, err = cpu.Counts(true)
	if err != nil {
		c.Logical = -1
	}
	c.Info, err = cpu.Info()
	return err
}

type CpuInfoMonitor struct {
	PrePercent []float64       `json:"pre_percent"` // 每个 cpu 的使用率
	Percent    float64         `json:"percent"`     // cpu 总使用率
	PreTimes   []cpu.TimesStat `json:"pre_times"`   // 每个 cpu 的各类使用时间
	Times      cpu.TimesStat   `json:"times"`       // cpu 总使用时间
}

func (c *CpuInfoMonitor) Get() error {
	var err error
	c.PrePercent, err = cpu.Percent(0, true)
	p, err := cpu.Percent(0, false)
	if err == nil {
		c.Percent = p[0]
	}
	c.PreTimes, err = cpu.Times(true)
	t, err := cpu.Times(false)
	if err == nil {
		c.Percent = p[0]
	}
	c.Times = t[0]
	return err
}

type DiskInfo struct {
	Partitions []disk.PartitionStat `json:"partitions"`
}

func (d *DiskInfo) Get() error {
	var err error
	d.Partitions, err = disk.Partitions(false)
	return err
}

type DiskInfoMonitor struct {
	Usage *disk.UsageStat `json:"usage"`
}

func (d *DiskInfoMonitor) Get() error {
	var err error
	diskInfo := new(DiskInfo)
	err = diskInfo.Get()
	us, _ := disk.Usage("/")
	d.Usage = us

	return err
}

type DockerInfo struct {
	Has bool `json:"has"`
}

func (d *DockerInfo) Get() error {
	_, err := docker.GetDockerStat()
	if err != nil {
		d.Has = false
	} else {
		d.Has = true
	}
	return nil
}

type DockerInfoMonitor struct {
	CgroupDockerStat []docker.CgroupDockerStat `json:"cgroup_docker_stat"`
	CgroupMemDocker  []*docker.CgroupMemStat   `json:"cgroup_mem_docker"`
	CgroupCPUDocker  []*docker.CgroupCPUStat   `json:"cgroup_cpu_docker"`
}

func (d *DockerInfoMonitor) Get() error {
	var err error
	d.CgroupDockerStat, err = docker.GetDockerStat()
	if err != nil {
		return err
	}

	for _, item := range d.CgroupDockerStat {
		memStat, _ := docker.CgroupMemDocker(item.ContainerID)
		cpuStat, _ := docker.CgroupCPUDocker(item.ContainerID)
		d.CgroupMemDocker = append(d.CgroupMemDocker, memStat)
		d.CgroupCPUDocker = append(d.CgroupCPUDocker, cpuStat)
	}

	return nil
}

type HostInfo struct {
	Info *host.InfoStat `json:"info"`
}

func (h *HostInfo) Get() error {
	info, err := host.Info()
	h.Info = info
	return err
}

type HostInfoMonitor struct {
	Temperatures []host.TemperatureStat `json:"temperatures"`
}

func (h *HostInfoMonitor) Get() error {
	temp, err := host.SensorsTemperatures()
	h.Temperatures = temp
	return err
}

type LoadInfoMonitor struct {
	Avg  *load.AvgStat  `json:"avg"`
	Misc *load.MiscStat `json:"misc"`
}

func (l *LoadInfoMonitor) Get() error {
	var err error
	l.Avg, err = load.Avg()
	l.Misc, err = load.Misc()
	return err
}

type MemInfoMonitor struct {
	SwapMemoryStat    *mem.SwapMemoryStat    `json:"swap_memory_stat"`
	VirtualMemoryStat *mem.VirtualMemoryStat `json:"virtual_memory_stat"`
}

func (m *MemInfoMonitor) Get() error {
	var err error
	m.SwapMemoryStat, err = mem.SwapMemory()
	m.VirtualMemoryStat, err = mem.VirtualMemory()
	return err
}

type NetInfoMonitor struct {
	TcpConn    []net.ConnectionStat `json:"tcp_conn"`
	UdpConn    []net.ConnectionStat `json:"udp_conn"`
	IOCounters net.IOCountersStat   `json:"io_counters"`
}

func (n *NetInfoMonitor) Get() error {
	var err error
	n.TcpConn, err = net.Connections("tcp")
	n.UdpConn, err = net.Connections("udp")
	var ioCounters []net.IOCountersStat
	ioCounters, err = net.IOCounters(false)
	n.IOCounters = ioCounters[0]
	return err
}

type ProcessMonitor struct {
	Process []*process.Process `json:"process"`
}

func (p *ProcessMonitor) Get() error {
	var err error
	p.Process, err = process.Processes()
	return err
}
