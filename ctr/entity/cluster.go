package entity

type ClusterRetrieveRequestData struct {
	TreeID *int64 `form:"tree_id" binding:"required,gt=0" label:"服务树Cluster类型节点的id"`
}

type ClusterCreateRequestData struct {
	TreeID        *int64 `json:"tree_id" binding:"required,gt=0" label:"服务树Cluster类型节点的id"`
	ClusterNum    *int   `json:"cluster_num" binding:"required,gt=1" label:"集群容器数量"`
	ContainerPort *int   `json:"container_port" binding:"required,gt=0" label:"容器端口"`
}

type ClusterListRequestData struct {
	TreeID *int64 `form:"tree_id" binding:"required,gt=0" label:"服务树Service类型节点的id"`
}
