package entity

type PipeLineListRequestData struct {
	TreeID *int64 `form:"tree_id" binding:"required,gt=0" label:"服务树Service类型节点的id"`
}

type PipeLineCreateRequestData struct {
	TreeID          *int64 `json:"tree_id" binding:"required,gt=0" label:"服务树Service类型节点的id"`
	ClusterID       *int64 `json:"cluster_id" binding:"required,gt=0" label:"cluster表id"`
	ContainerImage  string `json:"container_image" binding:"required" label:"镜像地址"`
	AliveMethod     string `json:"alive_method" binding:"required,min=3,max=10" label:"存活测试方法"`
	AliveURI        string `json:"alive_uri" binding:"required,min=1,max=1024" label:"存活测试uri地址"`
	AliveReqQuery   string `json:"alive_req_query" binding:"required" label:"存活测试url参数"`
	AliveReqBody    string `json:"alive_req_body" binding:"required" label:"存活测试body数据"`
	AliveReqHeader  string `json:"alive_req_header" binding:"required" label:"存活测试请求头"`
	AliveRespStatus int    `json:"alive_resp_status" binding:"required" label:"存活测试响应状态码"`
	AliveRespBody   string `json:"alive_resp_body" binding:"required" label:"存活测试响应体"`
}

type PipeLineRunRequestData struct {
	Id *int64 `json:"id" binding:"required,gt=0" label:"流水线Id"`
}

type PipeLineStatusRequestData struct {
	Id *int64 `form:"id" binding:"required,gt=0" label:"流水线Id"`
}

type PipeLiseStatusResponseData struct {
	Current int      `json:"current"`
	Steps   []string `json:"steps"`
	Content []string `json:"content"`
	Logs    string   `json:"logs"`
}
