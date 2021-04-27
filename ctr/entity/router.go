package entity

type MarkHostAsRouterRequestData struct {
	NsId   *int64 `json:"ns_id" binding:"required,gt=0" label:"ns节点id"`
	HostId *int64 `json:"host_id" binding:"required,gt=0" label:"主机id"`
}

type RouterListRequestData struct {
	Id     *int64 `form:"id" binding:"omitempty,gt=0" label:"id"`
	NsId   *int64 `form:"ns_id" binding:"omitempty,gt=0" label:"ns节点id"`
	HostId *int64 `form:"host_id" binding:"omitempty,gt=0" label:"主机id"`
}

type RouterStatusRequestData struct {
	Id *int64 `form:"id" binding:"required,gt=0" label:"id"`
}

type RouterStatusResponseData struct {
	NginxConfig    map[string]interface{} `json:"nginx_config"`
	NginxStatus    string                 `json:"nginx_status"`
	LogstashConfig string                 `json:"logstash_config"`
	LogstashStatus string                 `json:"logstash_status"`
}
