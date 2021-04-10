package main

import (
	"github.com/sevlyar/go-daemon"
	"github.com/sirupsen/logrus"
	"time"

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
	for {
		logrus.Info(time.Now())
		time.Sleep(1 * time.Second)
	}
}
