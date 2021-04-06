package entity

import "github.com/ahojcn/ecloud/ctr/model"

type CreateTreeNodeRequestData struct {
	Name        string `json:"name" binding:"required,alphanum,min=5,max=128" label:"节点名"`
	Description string `json:"description" binding:"required,min=5,max=1024" label:"节点描述"`
	Type        int    `json:"type" binding:"required,gt=0,lt=4" label:"节点类型"`
	ParentId    int64  `json:"parent_id" binding:"required,gt=1"`
}

type GetTreeNodesResponseData struct {
	*model.UserTreeInfo
	Children []*GetTreeNodesResponseData `json:"children"`
}
