package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/service/nginx"
	"net/http"
	"strings"

	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
)

func MarkHostAsRouter(user *model.User, rd *entity.MarkHostAsRouterRequestData) (int, []string, error) {
	ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": *rd.NsId})
	if !has || ut.Rights != model.PermAdmin {
		return http.StatusUnauthorized, nil, fmt.Errorf("无权限")
	}

	if ut.UserTree2UserTreeInfo().Tree.Type != model.TreeTypeNamespace {
		return http.StatusBadRequest, nil, fmt.Errorf("节点类型必须是Namespace")
	}

	hu, has := model.HostUserOne(map[string]interface{}{"user_id": user.Id, "host_id": *rd.HostId})
	if !has {
		return http.StatusUnauthorized, nil, fmt.Errorf("无权限")
	}

	if _, has = model.RouterOne(map[string]interface{}{"ns_id": *rd.NsId, "host_id": *rd.HostId}); has {
		return http.StatusBadRequest, nil, fmt.Errorf("已存在")
	}

	r := &model.Router{NsId: *rd.NsId, HostId: *rd.HostId}
	if err := model.RouterAdd(r); err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("服务器错误:err=%v", err)
	}

	host := hu.GetHost()
	model.HostUpdateRouter(true)

	res, err := DeployRouter(host)
	if err != nil {
		_ = model.RouterDelete(r)
		model.HostUpdateRouter(false)
		return http.StatusInternalServerError, nil, fmt.Errorf("部署router失败:err=%v", err)
	}

	r.Log = strings.Join(res, "\n")
	err = model.RouterUpdate(r.Id, r)

	return http.StatusOK, res, err
}

func RouterList(user *model.User, rd *entity.RouterListRequestData) (int, []model.RouterInfo, error) {
	res := make([]model.RouterInfo, 0)

	if rd.Id != nil {
		r, has := model.RouterOne(map[string]interface{}{"id": *rd.Id})
		if !has {
			return http.StatusNotFound, nil, fmt.Errorf("没有对应的接入层信息")
		}
		res = append(res, *r.GetRouterInfo())
		return http.StatusOK, res, nil
	}

	if rd.NsId != nil {
		rl, err := model.RouterList(map[string]interface{}{"ns_id": *rd.NsId})
		if err != nil {
			return http.StatusNotFound, nil, fmt.Errorf("没有对应的接入层信息")
		}
		for _, r := range rl {
			res = append(res, *r.GetRouterInfo())
		}
		return http.StatusOK, res, nil
	}

	if rd.HostId != nil {
		rl, err := model.RouterList(map[string]interface{}{"host_id": *rd.HostId})
		if err != nil {
			return http.StatusNotFound, nil, fmt.Errorf("没有对应的接入层信息")
		}
		for _, r := range rl {
			res = append(res, *r.GetRouterInfo())
		}
		return http.StatusOK, res, nil
	}

	rl, err := model.RouterList(map[string]interface{}{})
	if err != nil {
		return http.StatusInternalServerError, res, fmt.Errorf("查询接入层出错,err:%v", err)
	}

	for _, r := range rl {
		res = append(res, *r.GetRouterInfo())
	}

	return http.StatusOK, res, nil
}

func RouterStatus(user *model.User, rd *entity.RouterStatusRequestData) (int, *entity.RouterStatusResponseData, error) {
	if rd.Id == nil {
		return http.StatusBadRequest, nil, fmt.Errorf("参数错误，id不能为空")
	}
	r, has := model.RouterOne(map[string]interface{}{"id": *rd.Id})
	if !has {
		return http.StatusBadRequest, nil, fmt.Errorf("不存在的router id")
	}

	h, err := r.GetHostInfo()
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("获取Router关联的主机信息失败")
	}

	res := new(entity.RouterStatusResponseData)

	res.NginxStatus, _ = h.RunCmd("ps x | grep nginx", 0)
	res.LogstashStatus, _ = h.RunCmd("ps x | grep agent", 0)

	res.LogstashConfig, _ = h.RunCmd("cat /root/.eCloud/logstash/logstash.conf", 0)

	res.NginxConfig = make(map[string]interface{})
	nginxConfigsFileName, _ := h.RunCmd("ls /root/.eCloud/nginx/conf/conf.d/", 0)
	nginxConfigFiles := strings.Split(nginxConfigsFileName, "\n")
	for i := 0; i < len(nginxConfigFiles)-1; i++ {
		fileContent, _ := h.RunCmd(fmt.Sprintf("cat /root/.eCloud/nginx/conf/conf.d/%v", nginxConfigFiles[i]), 0)
		res.NginxConfig[nginxConfigFiles[i]] = fileContent
	}

	return http.StatusOK, res, nil
}

func NginxConfig(user *model.User, rd *entity.NginxConfigRequestData) (int, []string, error) {
	if rd == nil || rd.NsId == nil {
		return http.StatusBadRequest, nil, fmt.Errorf("参数错误")
	}

	r, has := model.RouterOne(map[string]interface{}{"ns_id": *rd.NsId})
	if !has {
		return http.StatusBadRequest, nil, fmt.Errorf("未查询到nsid,可能未部署router")
	}

	h, err := r.GetHostInfo()
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("服务器开小差了，err:%v", err)
	}

	ngx := nginx.New(h)
	if rd.FileName != nil {
		res, err := ngx.ConfContent(*rd.FileName)
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("服务器开小差了，err:%v", err)
		}
		return http.StatusOK, []string{res}, nil
	} else {
		fileNames, err := ngx.ConfList()
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("服务器开小差了，err:%v", err)
		}
		return http.StatusOK, fileNames, nil
	}
}
