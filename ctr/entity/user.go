package entity

type UserRegisterRequestData struct {
	Username  string `json:"username" binding:"required,alphanum,min=5,max=32" label:"用户名"`
	Password  string `json:"password" binding:"required,min=5,max=18" label:"密码"`
	RPassword string `json:"r_password" binding:"required,eqfield=Password" label:"重复密码"`
	Email     string `json:"email" binding:"required,email" label:"邮箱"`
}

type UserInfoByIdRequestData struct {
	Id int64 `uri:"id"`
}
