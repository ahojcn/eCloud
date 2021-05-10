package nginx

import (
	"fmt"
	"strings"
)

// ConfList 获取 nginx 根配置
func (n *Nginx) ConfRoot() (string, error) {
	cmd := fmt.Sprintf("cat %s", RootConfPath)
	res, err := n.Host.RunCmd(cmd, RunCmdTimeOut)
	if err != nil {
		return "", fmt.Errorf("获取nginx根配置失败,err:%v,cmd:%s", err, cmd)
	}
	if res == "" {
		return "", fmt.Errorf("获取nginx根配置失败,err:%v,cmd:%s", err, cmd)
	}
	return res, nil
}

// ConfList 获取 nginx 配置文件列表
// 只返回配置文件的名字数组
func (n *Nginx) ConfList() ([]string, error) {
	cmd := fmt.Sprintf("ls %s", ConfDirPath)
	res, err := n.Host.RunCmd(cmd, RunCmdTimeOut)
	if err != nil {
		return nil, fmt.Errorf("获取nginx配置列表失败,err:%v,cmd:%s", err, cmd)
	}
	res = strings.Trim(res, "\n")
	if res == "" {
		return []string{}, nil
	}
	return strings.Split(res, "\n"), nil
}

// ConfContent 返回某个配置文件的内容
func (n *Nginx) ConfContent(fileName string) (string, error) {
	filePath := ConfDirPath + fileName
	// 判断文件是否存在
	cmd := fmt.Sprintf("ls %s", filePath)
	res, err := n.Host.RunCmd(cmd, RunCmdTimeOut)
	if err != nil || res == "" {
		return "", fmt.Errorf("获取nginx配置失败,文件可能不存在,err:%v,cmd:%s", err, cmd)
	}

	cmd = fmt.Sprintf("cat %s", filePath)
	res, err = n.Host.RunCmd(cmd, RunCmdTimeOut)
	if err != nil {
		return "", fmt.Errorf("获取nginx配置失败,err:%v,cmd:%s", err, cmd)
	}
	return res, nil
}
