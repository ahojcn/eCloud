package controller

import (
	"net/http"

	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
)

// GetHostInfo 获取主机列表接口
func GetHostInfo(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.GetHostInfoRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	hil, err := service.GetHostInfoList(user, rd)
	if err != nil {
		g.response(http.StatusInternalServerError, "获取失败", err)
		return
	}

	g.response(http.StatusOK, "获取主机信息完成", hil)
}

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

	if err = service.CreateHost(user, &rd); err != nil {
		g.response(http.StatusInternalServerError, "添加主机失败", err)
		return
	}

	g.response(http.StatusOK, "添加成功 && 部署 agent 成功", nil)
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

	if status, err := service.DeleteHost(user, &rd); err != nil {
		g.response(status, "删除失败", err)
		return
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
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if status, err := service.DeleteHostUser(user, rd); err != nil {
		g.response(status, err.Error(), nil)
		return
	}

	g.response(http.StatusOK, "删除成功", nil)
}

func RunCommand(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.RunCommandRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, result, err := service.RunCommand(user, rd)
	if err != nil {
		g.response(status, "执行失败", err)
		return
	}

	g.response(http.StatusOK, "执行成功", result)
}
