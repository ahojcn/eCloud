package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func CPUTest() {
	fmt.Println(cpu.Counts(true))
	fmt.Println(cpu.Counts(false))
	fmt.Println(cpu.Info())

	fmt.Println(cpu.Percent(0, true))
	fmt.Println(cpu.Percent(0, false))
	fmt.Println(cpu.Times(true))
	fmt.Println(cpu.Times(false))
}

func MEMTest() {
	fmt.Println(mem.SwapMemory())
	fmt.Println(mem.VirtualMemory())
}

func DISKTest() {
	fmt.Println(disk.Partitions(false))
	fmt.Println(disk.Partitions(true))

	fmt.Println(disk.Usage("/"))
	fmt.Println(disk.IOCounters("/System/Volumes/Data/home"))
}

func DockerTest() {
	ids, err := docker.GetDockerIDList()
	fmt.Println(ids, err)
	ds, err := docker.GetDockerStat()
	fmt.Println(ds, err)

	fmt.Println(ids[0])
	fmt.Println(docker.CgroupMemDocker(ids[0]))
	fmt.Println(docker.CgroupCPUDocker(ids[0]))
}

func HostTest() {
	fmt.Println(host.Info())

	fmt.Println(host.SensorsTemperatures())
}

func LoadTest() {
	fmt.Println(load.Avg())
	fmt.Println(load.Misc())
}

func NetTest() {
	tcpConn, _ := net.Connections("tcp")
	fmt.Println(len(tcpConn))
	udpConn, _ := net.Connections("udp")
	fmt.Println(len(udpConn))
	allConn, _ := net.Connections("*")
	fmt.Println(len(allConn))

	fmt.Println(net.IOCounters(true))
	fmt.Println(net.IOCounters(false))
}

func ProcessTest() {
	fmt.Println(process.Processes())
}

func main() {
	//CPUTest()
	//DISKTest()
	//DockerTest()
	//HostTest()
	//LoadTest()
	//MEMTest()
	//NetTest()
	//ProcessTest()
}
