package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHost(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	// 检查参数
	rd := entity.CreateHostRequestData{}
	err = c.ShouldBindJSON(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	host := model.Host{
		UserId:      user.Id,
		Description: rd.Description,
		IP:          rd.IP,
		Username:    rd.Username,
		Password:    rd.Password,
		Port:        rd.Port,
	}
	err = model.HostAdd(&host)
	if err != nil {
		g.response(http.StatusInternalServerError, "添加主机信息失败", err)
		return
	}

	res, err := service.DeployAgent(host)
	if err != nil {
		g.response(http.StatusInternalServerError, "部署失败", append(res, err.Error()))
		return
	}

	g.response(http.StatusOK, "添加成功 && 部署 agent 成功", res)
}

func UpdateHost(c *gin.Context) {
	g := newGin(c)

	// 校验参数
	rd := entity.UpdateHostRequestData{}
	err := c.ShouldBind(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	h := &model.Host{
		UserId:      rd.UserId,
		Description: rd.Description,
		IP:          rd.IP,
		Username:    rd.Username,
		Password:    rd.Password,
		Port:        rd.Port,
		Extra:       rd.Extra,
	}
	err = model.HostUpdate(rd.HostId, h)
	if err != nil {
		g.response(http.StatusInternalServerError, "更新失败", err)
		return
	}

	g.response(http.StatusOK, "更新成功", nil)
}

func DeleteHost(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	// 校验参数
	rd := entity.DeleteHostRequestData{}
	err = c.ShouldBindUri(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	// 查找主机
	host, has := model.HostOne(map[string]interface{}{"id": rd.HostId})
	if !has {
		g.response(http.StatusNotFound, "未找到host", rd.HostId)
		return
	}

	// 判断这个用户是否是管理员
	if user.Id != host.UserId {
		g.response(http.StatusUnauthorized, "权限不足", nil)
		return
	}

	err = model.HostDelete(host)
	if err != nil {
		g.response(http.StatusInternalServerError, "删除失败", err)
		return
	}

	// 删除这个主机相关的权限
	hus, _ := model.HostUserList(map[string]interface{}{"host_id": rd.HostId})
	for _, hu := range hus {
		_ = model.HostUserDelete(&hu)
	}

	g.response(http.StatusOK, "删除成功", nil)
}

func CreateHostUser(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.CreateHostUserRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if rd.UserId == user.Id {
		g.response(http.StatusOK, "你已经有权限啦，无须添加", nil)
		return
	}

	_, h1 := model.UserOne(map[string]interface{}{"id": rd.UserId})
	host, h2 := model.HostOne(map[string]interface{}{"id": rd.HostId})
	if !h1 || !h2 {
		g.response(http.StatusBadRequest, "参数错误，用户||主机不存在", map[string]bool{
			"user": h1,
			"host": h2,
		})
		return
	}

	// 判断这个用户是否是管理员
	if user.Id != host.UserId {
		g.response(http.StatusUnauthorized, "权限不足", nil)
		return
	}

	if hu, has := model.HostUserOne(map[string]interface{}{"host_id": rd.HostId, "user_id": rd.UserId}); has {
		g.response(http.StatusOK, "已添加过", hu)
		return
	}

	hu := new(model.HostUser)
	hu.UserId = rd.UserId
	hu.HostId = rd.HostId
	if err = model.HostUserAdd(hu); err != nil {
		g.response(http.StatusInternalServerError, "添加失败", err)
		return
	}

	g.response(http.StatusOK, "添加成功", hu)
}

func DeleteHostUser(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.DeleteHostUserRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if rd.UserId == user.Id {
		g.response(http.StatusOK, "请先删除主机", nil)
		return
	}

	_, h1 := model.UserOne(map[string]interface{}{"id": rd.UserId})
	host, h2 := model.HostOne(map[string]interface{}{"id": rd.HostId})
	if !h1 || !h2 {
		g.response(http.StatusBadRequest, "参数错误，用户||主机不存在", map[string]bool{
			"user": h1,
			"host": h2,
		})
		return
	}

	// 判断这个用户是否是管理员
	if user.Id != host.UserId {
		g.response(http.StatusUnauthorized, "权限不足", nil)
		return
	}

	hu := new(model.HostUser)
	hu.HostId, hu.UserId = rd.HostId, rd.UserId
	if err = model.HostUserDelete(hu); err != nil {
		g.response(http.StatusInternalServerError, "删除失败", err)
		return
	}

	g.response(http.StatusOK, "删除成功", nil)
}
