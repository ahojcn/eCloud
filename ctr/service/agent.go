package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	"time"
)

func DeployAgent(host model.Host) ([]string, error) {
	runCmdTimeout := 10 * time.Second
	var res []string
	ip := util.Config.Section("system").Key("ip").String()
	port := util.Config.Section("system").Key("listen_port").String()
	deployPath := util.Config.Section("agent").Key("deploy_path").String()
	deployShell := util.Config.Section("agent").Key("deploy_shell").String()
	cmd := fmt.Sprintf("mkdir -p %s"+
		"&& cd %s"+
		"&& curl -fsSL %s/deploy.sh --output deploy.sh"+
		"&& chmod +x deploy.sh"+
		"&& ./deploy.sh %d %s %s %s"+
		"&& echo ok",
		deployPath, deployPath, deployShell, host.Id, ip, port, deployShell)
	res1, err := host.RunCmd(cmd, runCmdTimeout)
	res = append(res, cmd)
	if err != nil {
		return res, err
	}
	cmd = "cd /root/.eCloud && cat agent.pid"
	res2, err := host.RunCmd(cmd, runCmdTimeout)
	res = append(res, cmd)
	if err != nil {
		return res, err
	}
	res = append(res, res1)
	res = append(res, res2)
	return res, err
}
