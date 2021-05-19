package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClusterRetrieve(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.ClusterRetrieveRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.ClusterRetrieve(user, rd)
	if err != nil {
		g.response(status, "获取集群配置失败", err)
		return
	}

	g.response(http.StatusOK, "获取集群配置成功", res)
}

func ClusterCreate(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.ClusterCreateRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.ClusterCreate(user, rd)
	if err != nil {
		g.response(status, "创建集群配置失败", err)
		return
	}

	g.response(http.StatusOK, "创建集群配置成功", res)
}

func ClusterList(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.ClusterListRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.ClusterList(user, rd)
	if err != nil {
		g.response(status, "获取集群列表失败", err)
		return
	}

	g.response(http.StatusOK, "获取集群列表成功", res)
}
