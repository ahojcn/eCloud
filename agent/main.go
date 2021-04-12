package main

import (
	"github.com/ahojcn/ecloud/agent/ctr"
	"github.com/sevlyar/go-daemon"
	"log"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "agent.pid",
		PidFilePerm: 0644,
		LogFileName: "agent.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"agent"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	Start()
}

func Start() {
	ctr.ReportHostExtra()
	ctr.RunMetrics()
	select {
	}
}
