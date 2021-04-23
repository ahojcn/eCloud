package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateICode(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.CreateICodeRequestData)
	if err = c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, err := service.CreateICode(user, rd)
	if err != nil {
		g.response(status, "创建失败", err)
		return
	}

	g.response(http.StatusOK, "创建开发机成功", nil)
}

func GetICodeList(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.GetICodeListRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	status, iCodes, err := service.GetICodeList(user, rd)
	if err != nil {
		g.response(status, "获取开发机列表失败", err)
		return
	}

	g.response(http.StatusOK, "获取开发机列表完成", iCodes)
}

func DeleteICode(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := new(entity.DeleteICodeRequestData)
	if err = c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if status, err := service.DeleteICode(user, rd); err != nil {
		g.response(status, "删除失败", err)
		return
	}

	g.response(http.StatusOK, "删除开发机成功", nil)
}
