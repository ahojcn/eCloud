package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"net/http"
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

	// todo 添加 host_user 信息
	if err = model.HostUserAdd(&model.HostUser{UserId: user.Id, HostId: host.Id}); err != nil {
		_ = model.HostDelete(&host)
		return fmt.Errorf("添加主机信息失败, err:%v", err)
	}

	res, err := DeployAgent(host)
	if err != nil {
		_ = model.HostDelete(&host)
		return fmt.Errorf("部署失败, err:%v", append(res, err.Error()))
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
