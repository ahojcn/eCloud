package entity

import "github.com/ahojcn/ecloud/ctr/model"

type CreateICodeRequestData struct {
	Name *string `json:"name" binding:"required,min=1,max=32" label:"开发机名称"`
}

type DeleteICodeRequestData struct {
	Id *int `form:"id" binding:"required,gt=0" label:"开发机id"`
}

type GetICodeListRequestData struct {
	Id *int `form:"id" binding:"omitempty,gt=0" label:"开发机id"`
}

type GetICodeListResponseData struct {
	model.ICode
	UserInfo *model.UserInfo `json:"user_info"`
	HostInfo *model.HostInfo `json:"host_info"`
}
