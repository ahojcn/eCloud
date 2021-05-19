package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PipeLineList(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.PipeLineListRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.PipeLineList(user, rd)
	if err != nil {
		g.response(status, "获取流水线列表失败", err)
		return
	}

	g.response(http.StatusOK, "获取流水线列表完成", res)
}

func PipeLineCreate(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.PipeLineCreateRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.PipeLineCreate(user, rd)
	if err != nil {
		g.response(status, "创建流水线失败", err)
		return
	}

	g.response(http.StatusOK, "创建流水线成功", res)
}

func PipeLineRun(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.PipeLineRunRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, res, err := service.PipeLineRun(user, rd)
	if err != nil {
		g.response(status, "运行流水线失败", err)
		return
	}

	g.response(http.StatusOK, "运行流水线成功", res)
}
