package ctr

import (
	"encoding/json"
	"github.com/ahojcn/ecloud/agent/util"
	"testing"
)

func TestReportHostInfo(t *testing.T) {
	cpuinfo := new(util.CpuInfoMonitor)
	cpuinfo.Get()
	b, _ := json.Marshal(cpuinfo)
	t.Log(string(b))
}
