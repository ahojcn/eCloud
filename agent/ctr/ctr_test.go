package ctr

import (
	"fmt"
	"strings"
	"testing"
)

func TestReportHostInfo(t *testing.T) {
	cols := []string{"user", "system", "idle"}
	ss := []string{}
	for _, s := range cols {
		ss = append(ss, fmt.Sprintf("\"%s\"", s))
	}
	fmt.Println(fmt.Sprintf("x %s x", strings.Join(ss, ",")))
	//cpuinfo := new(util.ProcessMonitor)
	//cpuinfo.Get()
	//b, _ := json.Marshal(cpuinfo)
	//t.Log(string(b))
}
