package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServiceCreate(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.CreateServiceRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, content, err := service.CreateService(user, rd)
	if err != nil {
		g.response(status, "创建预案配置失败", err)
		return
	}

	g.response(http.StatusOK, "创建源配置完成", content)
}

func ServiceGet(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.ServiceGetRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, si, err := service.GetService(user, rd)
	if err != nil {
		g.response(status, "获取预案配置失败", err)
		return
	}

	g.response(http.StatusOK, "获取预案配置完成", si)
}

func ServiceDelete(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.ServiceGetRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, err := service.DeleteService(user, rd)
	if err != nil {
		g.response(status, "删除预案配置失败", err)
		return
	}

	g.response(http.StatusOK, "已删除", nil)
}