package entity

type LoginRequestData struct {
	Username string `json:"username" binding:"required,alphanum,min=5,max=32" label:"用户名"`
	Password string `json:"password" binding:"required,min=5,max=18" label:"密码"`
}
