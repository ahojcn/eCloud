package controller

import (
	"net/http"

	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
)

func MarkHostAsRouter(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.MarkHostAsRouterRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.MarkHostAsRouter(user, rd)
	if err != nil {
		g.response(status, "创建router失败", err)
		return
	}

	g.response(http.StatusOK, "创建router成功", res)
}

func RouterList(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.RouterListRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.RouterList(user, rd)
	if err != nil {
		g.response(status, "获取接入层信息失败", err)
		return
	}

	g.response(http.StatusOK, "获取接入层信息成功", res)
}

func RouterStatus(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.RouterStatusRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.RouterStatus(user, rd)
	if err != nil {
		g.response(status, "获取接入层状态失败", err)
		return
	}

	g.response(http.StatusOK, "获取接入层状态成功", res)
}

func NginxConfig(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.NginxConfigRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.NginxConfig(user, rd)
	if err != nil {
		g.response(status, "获取nginx配置失败", err)
		return
	}

	g.response(http.StatusOK, "获取nginx配置完成", res)
}
