package entity

type MonitorWriteMetricsRequestData struct {
	HostId  int64                  `json:"host_id" binding:"required,gt=0" label:"主机id"`
	Metrics string                 `json:"metrics" binding:"required" label:"指标"`
	Data    map[string]interface{} `json:"data" binding:"required" label:"数据"`
}

type MonitorQueryMetricsRequestData struct {
	HostId   *int64   `form:"host_id" binding:"required,gt=0" label:"主机id"`
	Metrics  *string  `form:"metrics" binding:"required" label:"指标"`
	Cols     []string `form:"cols" binding:"required" label:"列"`
	FromTime *string  `form:"from_time" binding:"required" label:"开始时间"`
	ToTime   *string  `form:"to_time" binding:"required" label:"开始时间"`
}

type RouterWriteMetricsRequestData map[string]string

type RouterMonitorMetricsQueryRequestData struct {
	Un       *string `form:"un" binding:"required" label:"服务树的unique name"`
	Uri      *string `form:"uri" binding:"required" label:"uri"`
	Docker   *string `form:"docker" binding:"required" label:"docker"`
	Metrics  *string `form:"metrics" binding:"required" label:"metrics"`
	FromTime *string `form:"from_time" binding:"required" label:"开始时间"`
	ToTime   *string `form:"to_time" binding:"required" label:"开始时间"`
	Overview *bool   `form:"overview" binding:"required" label:"是否获取预览数据"`
}
