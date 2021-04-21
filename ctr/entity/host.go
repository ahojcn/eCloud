package entity

type GetHostInfoRequestData struct {
	Id *int64 `form:"id" binding:"omitempty,gt=0" label:"主机id"`
}

type CreateHostRequestData struct {
	IP          string `json:"ip" binding:"required,ipv4" label:"ipv4地址"`
	Username    string `json:"username" binding:"required,alphanum,min=1,max=128" label:"用户名"`
	Password    string `json:"password" binding:"required,min=1,max=128" label:"密码"`
	Port        int    `json:"port" binding:"required,gt=0,lt=65534" label:"ssh端口"`
	Description string `json:"description" binding:"required,min=1,max=1023" label:"描述"`
}

type UpdateHostRequestData struct {
	HostId      int64  `json:"host_id" binding:"required,gt=0" label:"主机id"`
	UserId      int64  `json:"user_id" binding:"omitempty,gt=0" label:"用户id"`
	Description string `json:"description" binding:"omitempty,min=1,max=1023" label:"主机描述"`
	IP          string `json:"ip" binding:"omitempty,ipv4" label:"ip地址"`
	Username    string `json:"username" binding:"omitempty,min=1,max=127" label:"用户名"`
	Password    string `json:"password" binding:"omitempty,min=1,max=127" label:"密码"`
	Port        int    `json:"port" binding:"omitempty,gt=0,lt=65534" label:"端口"`
	Extra       string `json:"extra" binding:"omitempty" label:"额外信息"`
}

type DeleteHostRequestData struct {
	HostId int64 `uri:"id" binding:"required,gt=0" label:"主机id"`
}

type CreateHostUserRequestData struct {
	HostId int64 `json:"host_id" binding:"required,gt=0" label:"主机id"`
	UserId int64 `json:"user_id" binding:"required,gt=0" label:"用户id"`
}

type DeleteHostUserRequestData struct {
	HostId *int64 `form:"host_id" binding:"required,gt=0" label:"主机id"`
	UserId *int64 `form:"user_id" binding:"required,gt=0" label:"用户id"`
}

type RunCommandRequestData struct {
	HostId *int64  `json:"host_id" binding:"required,gt=0" label:"主机id"`
	Cmd    *string `json:"cmd" binding:"required,min=0" label:"命令"`
}
