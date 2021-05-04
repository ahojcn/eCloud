package service

import (
	"fmt"

	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
)

func DeployAgent(host *model.Host) ([]string, error) {
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
	res1, err := host.RunCmd(cmd, 0)
	res = append(res, cmd)
	if err != nil {
		return res, err
	}
	cmd = "cd /root/.eCloud && cat agent.pid"
	res2, err := host.RunCmd(cmd, 0)
	res = append(res, cmd)
	if err != nil {
		return res, err
	}
	res = append(res, res1)
	res = append(res, res2)
	return res, err
}

// DeployRouter 会在目标机器上部署 nginx、logstash
func DeployRouter(host *model.Host) ([]string, error) {
	var res []string
	ip := util.Config.Section("system").Key("ip").String()
	port := util.Config.Section("system").Key("listen_port").String()
	deployPath := util.Config.Section("agent").Key("deploy_path").String()
	deployShell := util.Config.Section("agent").Key("deploy_shell").String()
	cmd := fmt.Sprintf("cd %s"+
		"&& curl -fsSL %s/router.sh --output router.sh"+
		"&& chmod +x router.sh"+
		"&& ./router.sh %s http://%s:%v"+
		"&& echo OK",
		deployPath, deployShell, deployShell, ip, port)
	res1, err := host.RunCmd(cmd, 0)
	res = append(res, res1)
	if err != nil {
		return res, err
	}
	cmd = fmt.Sprintf("cd %s && cat logstash.pid", deployPath)
	res2, err := host.RunCmd(cmd, 0)
	res = append(res, res2)
	if err != nil {
		return res, err
	}
	return res, err
}
