package entity

import "github.com/ahojcn/ecloud/ctr/model"

type CreateTreeNodeRequestData struct {
	Name        string `json:"name" binding:"required,alphanum,min=5,max=128" label:"节点名"`
	Description string `json:"description" binding:"required,min=5,max=1024" label:"节点描述"`
	Type        *int   `json:"type" binding:"required,gt=-1,lt=5" label:"节点类型"`
	ParentId    *int64 `json:"parent_id" binding:"required,gt=-1"`
}

type TreeNodeInfo struct {
	*model.UserTreeInfo
	Children []*TreeNodeInfo `json:"children"`
}

type TreeNodeDetail struct {
	UserTreeInfo *model.UserTreeInfo   `json:"user_tree_info"`
	Users        []model.UserTreeInfo `json:"users"`
	Children     []*TreeNodeInfo       `json:"children"`
}

type GetTreeNodeRequestData struct {
	Id   *int64  `form:"id" binding:"omitempty,gt=0" label:"服务树节点id"`
	Name *string `form:"name" binding:"omitempty" label:"服务树名"`
}

type CreateUserTreeRequestData struct {
	UserId *int64 `json:"user_id" binding:"required,gt=0" label:"用户id"`
	TreeId *int64 `json:"tree_id" binding:"required,gt=0" label:"节点id"`
	Rights *int   `json:"rights" binding:"required,gt=-1,lt=7" label:"权限"`
}
