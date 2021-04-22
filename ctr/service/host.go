package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"net/http"
	"time"
)

// GetHostInfoList 获取用户有权限的主机列表
func GetHostInfoList(user *model.User, rd *entity.GetHostInfoRequestData) (res []*model.HostInfo, err error) {
	if rd == nil || user == nil {
		return nil, fmt.Errorf("获取失败:参数错误或未登录")
	}

	hul := make([]model.HostUser, 0)

	if rd.Id != nil {
		hul, err = model.HostUserList(map[string]interface{}{"user_id": user.Id, "host_id": rd.Id})
		if err != nil {
			return nil, err
		}
	} else {
		hul, err = model.HostUserList(map[string]interface{}{"user_id": user.Id})
		if err != nil {
			return nil, err
		}
	}

	for _, hu := range hul {
		h := hu.GetHost()
		if h != nil {
			res = append(res, h.GetHostInfo())
		}
	}

	return res, nil
}

// CreateHost 添加主机
func CreateHost(user *model.User, rd *entity.CreateHostRequestData) error {
	host := model.Host{
		UserId:      user.Id,
		Description: rd.Description,
		IP:          rd.IP,
		Username:    rd.Username,
		Password:    rd.Password,
		Port:        rd.Port,
	}
	err := model.HostAdd(&host)
	if err != nil {
		return fmt.Errorf("添加主机信息失败, err:%v", err)
	}

	res, err := DeployAgent(host)
	if err != nil {
		_ = model.HostDelete(&host)
		return fmt.Errorf("部署失败, err:%v", append(res, err.Error()))
	}

	if err = model.HostUserAdd(&model.HostUser{UserId: user.Id, HostId: host.Id}); err != nil {
		_ = model.HostDelete(&host)
		return fmt.Errorf("添加主机信息失败, err:%v", err)
	}

	return nil
}

// DeleteHostUser 删除用户权限
func DeleteHostUser(user *model.User, rd *entity.DeleteHostUserRequestData) (int, error) {
	if *rd.UserId == user.Id {
		return http.StatusUnauthorized, fmt.Errorf("不能删除自己的权限")
	}

	adminUser, h1 := model.UserOne(map[string]interface{}{"id": rd.UserId})
	host, h2 := model.HostOne(map[string]interface{}{"id": rd.HostId})
	if !h1 || !h2 {
		return http.StatusBadRequest, fmt.Errorf("参数错误，用户||主机不存在")
	}

	// 判断这个用户是否是管理员
	if user.Id != host.UserId {
		return http.StatusUnauthorized, fmt.Errorf("权限不足, 请联系管理员 %v", adminUser.Username)
	}

	hu := new(model.HostUser)
	hu.HostId, hu.UserId = *rd.HostId, *rd.UserId
	if err := model.HostUserDelete(hu); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("删除失败, err:%v", err)
	}

	return http.StatusOK, nil
}

// DeleteHost 删除主机
func DeleteHost(user *model.User, rd *entity.DeleteHostRequestData) (int, error) {
	// 查找主机
	host, has := model.HostOne(map[string]interface{}{"id": rd.HostId})
	if !has {
		return http.StatusNotFound, fmt.Errorf("未找到host")
	}

	// 判断这个用户是否是管理员
	if user.Id != host.UserId {
		return http.StatusUnauthorized, fmt.Errorf("权限不足，仅管理员可删除，请联系管理 %v", host.GetHostInfo().CreateUser.Username)
	}

	// kill 这个主机的 agent
	res, err := host.RunCmd("cd /root/.eCloud && kill -9 `cat agent.pid`", time.Second*60)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("停止主机agent失败, 请手动重试, err:%v, res:%v", err, res)
	}

	if err = model.HostDelete(host); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("删除失败, err:%v", err)
	}

	// 删除这个主机相关的权限
	hus, _ := model.HostUserList(map[string]interface{}{"host_id": rd.HostId})
	for _, hu := range hus {
		_ = model.HostUserDelete(&hu)
	}

	return http.StatusOK, nil
}

// RunCommand 执行命令并返回结果
func RunCommand(user *model.User, rd *entity.RunCommandRequestData) (int, string, error) {
	if rd.HostId == nil || rd.Cmd == nil {
		return http.StatusBadRequest, "", fmt.Errorf("参数错误")
	}

	hu, has := model.HostUserOne(map[string]interface{}{"host_id": *rd.HostId, "user_id": user.Id})
	if !has {
		return http.StatusUnauthorized, "", fmt.Errorf("没有权限")
	}

	result, err := hu.GetHost().RunCmd(*rd.Cmd, 60*time.Second)
	if err != nil {
		return http.StatusInternalServerError, "", fmt.Errorf("服务器错误, err:%v", err)
	}

	return http.StatusOK, result, nil
}
