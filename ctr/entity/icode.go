package entity

type CreateICodeRequestData struct {
	Name *string `json:"name" binding:"required,min=1,max=32" label:"开发机名称"`
}

type DeleteICodeRequestData struct {
	Id *string `form:"id" binding:"required,gt=0" label:"开发机id"`
}
