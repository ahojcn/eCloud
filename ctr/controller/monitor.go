package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MonitorWriteMetrics(c *gin.Context) {
	g := newGin(c)

	rd := new(entity.MonitorWriteMetricsRequestData)
	if err := c.ShouldBindJSON(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	go service.MonitorMetricsWrite(rd.HostId, rd.Metrics, rd.Data)

	g.response(http.StatusOK, "ok", nil)
}

func MonitorQueryMetrics(c *gin.Context) {
	g := newGin(c)

	rd := new(entity.MonitorQueryMetricsRequestData)
	if err := c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	res, err := service.MonitorMetricsQuery(rd.HostId, rd.Metrics, rd.Cols)
	if err != nil {
		g.response(http.StatusInternalServerError, "查询出错", err)
		return
	}

	g.response(http.StatusOK, "ok", res)
}