package nginx

import "fmt"

func (n *Nginx) Status() bool {
	res, err := n.Host.RunCmd("ps -C nginx --no-header", RunCmdTimeOut)
	if err != nil || res == "" {
		return false
	}
	return true
}

func (n *Nginx) Reload() error {
	return n.nginxOptions(nginxReloadCmd)
}

func (n *Nginx) ReOpen() error {
	return n.nginxOptions(nginxReopenCmd)
}

func (n *Nginx) Stop() error {
	return n.nginxOptions(nginxStopCmd)
}

func (n *Nginx) Quit() error {
	return n.nginxOptions(nginxQuitCmd)
}

func (n *Nginx) Start() error {
	return n.nginxOptions(nginxStartCmd)
}

func (n *Nginx) nginxOptions(ops string) error {
	if n.Host == nil {
		return fmt.Errorf("未指定主机,err:nginx.host==nil")
	}

	// 判断是否安装 nginx
	cmd := fmt.Sprintf("ls %s", BinPath)
	res, err := n.Host.RunCmd(cmd, RunCmdTimeOut)
	if err != nil {
		return fmt.Errorf("指定主机上获取没有部署router,err:%v,cmd:%s", err, cmd)
	}
	if res == "" {
		return fmt.Errorf("指定主机上获取没有部署router,err:%v,cmd:%s", err, cmd)
	}

	cmd = fmt.Sprintf("%s %s", BinPath, ops)
	res, err = n.Host.RunCmd(cmd, RunCmdTimeOut)
	if err != nil {
		return fmt.Errorf("执行nginx操作失败,err=%s,cmd=%s", err, cmd)
	}
	if res != "" {
		return fmt.Errorf("执行nginx操作失败,err=%s,cmd=%s", err, cmd)
	}
	return nil
}
