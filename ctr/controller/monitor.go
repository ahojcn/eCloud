package controller

import (
	"net/http"
	"strconv"

	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
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

	res, err := service.MonitorMetricsQuery(*rd.HostId, *rd.Metrics, rd.Cols, *rd.FromTime, *rd.ToTime)
	if err != nil {
		g.response(http.StatusInternalServerError, "查询出错", err)
		return
	}

	g.response(http.StatusOK, "ok", res)
}

func RouterMonitorMetricsGet(c *gin.Context) {
	g := newGin(c)

	res, err := service.RouterMonitorMetricsGet()
	if err != nil {
		g.response(http.StatusInternalServerError, "查询出错", err)
		return
	}

	g.response(http.StatusOK, "ok", res)
}

func RouterMonitorMetricsWrite(c *gin.Context) {
	g := newGin(c)
	rd := entity.RouterWriteMetricsRequestData{}
	if err := c.ShouldBindJSON(&rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
	}

	data := map[string]interface{}{}
	data["status"], _ = strconv.Atoi(rd["status"])
	data["request_time"], _ = strconv.ParseFloat(rd["request_time"], 64)
	data["upstream_response_time"], _ = strconv.ParseFloat(rd["upstream_response_time"], 64)
	go service.RouterMonitorMetricsWrite(rd["un"], rd["uri"], data)

	g.response(http.StatusOK, "ok", nil)
}

func RouterMonitorMetricsQuery(c *gin.Context) {
	g := newGin(c)

	rd := new(entity.RouterMonitorMetricsQueryRequestData)
	if err := c.ShouldBindQuery(rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if *rd.Overview == true {
		res1, err := service.RouterMonitorMetricsQueryOverview(rd)
		if err != nil {
			g.response(http.StatusInternalServerError, "查询出错", err)
			return
		}
		g.response(http.StatusOK, "ok", res1)
		return
	}

	res2, err := service.RouterMonitorMetricsQuery(rd)
	if err != nil {
		g.response(http.StatusInternalServerError, "查询出错", err)
		return
	}
	g.response(http.StatusOK, "ok", res2)
}
