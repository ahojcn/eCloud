package entity

type CreateServiceRequestData struct {
	TreeId  *int64  `json:"tree_id" binding:"required,gt=0" label:"服务树id"`
	FlowMap *string `json:"flow_map" binding:"required" label:"预案配置信息"`
}

type ServiceGetRequestData struct {
	TreeId  *int64  `form:"tree_id" binding:"required,gt=0" label:"服务树id"`
}
