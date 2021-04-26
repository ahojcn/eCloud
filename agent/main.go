package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/agent/ctr"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"os/exec"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: os.Args[1] + ".pid",
		PidFilePerm: 0644,
		LogFileName: os.Args[1] + ".log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        os.Args,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	if os.Args[1] == "agent" {
		Start()
	} else if os.Args[1] == "logstash" {
		StartLogstash()
	}
}

func Start() {
	ctr.ReportHostExtra()
	ctr.RunMetrics()
	select {}
}

func StartLogstash() {
	c := exec.Command("/root/.eCloud/logstash/logstash-7.12.0/bin/logstash", "-f", "/root/.eCloud/logstash/logstash.conf")
	fmt.Println(c.Run())
	select {}
}
