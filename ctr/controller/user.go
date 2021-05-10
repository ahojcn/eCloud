package controller

import (
	"net/http"

	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	"github.com/gin-gonic/gin"
)

// CreateUser 用户注册
func CreateUser(c *gin.Context) {
	var (
		requestData entity.UserRegisterRequestData
	)
	g := newGin(c)
	err := c.ShouldBind(&requestData)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	// 查询数据库中是否存在
	_, has1 := model.UserOne(map[string]interface{}{"username": requestData.Username})
	_, has2 := model.UserOne(map[string]interface{}{"email": requestData.Email})
	if has1 || has2 {
		g.response(http.StatusBadRequest, "用户名||邮件已被注册", nil)
		return
	}
	// 存入数据库
	user := &model.User{
		Username: requestData.Username,
		Password: util.Md5Str(requestData.Password),
		Email:    requestData.Email,
		IsActive: 0,
	}
	err = model.UserAdd(user)
	if err != nil {
		g.response(http.StatusInternalServerError, "注册失败", err)
		return
	}

	g.response(http.StatusOK, "注册成功", user.User2UserInfo())
}

// GetUserInfoById 获取用户信息
func GetUserInfoById(c *gin.Context) {
	g := newGin(c)
	_, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	var rd entity.UserInfoByIdRequestData
	err = c.ShouldBindUri(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	user, has := model.UserOne(map[string]interface{}{"id": rd.Id})
	if !has {
		g.response(http.StatusNotFound, "未找到", rd.Id)
		return
	}

	g.response(http.StatusOK, "获取用户信息完成", user.User2UserInfo())
}

// GetUsersInfoByUsername 根据 username 模糊查询用户信息
func GetUsersInfoByUsername(c *gin.Context) {
	g := newGin(c)
	_, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", nil)
		return
	}

	var rd entity.GetUsersInfoByUsernameRequestData
	err = c.ShouldBindQuery(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	users, err := model.UsersInfoListByUsername(rd.Username)
	if err != nil {
		g.response(http.StatusInternalServerError, "服务器错误", err)
		return
	}

	g.response(http.StatusOK, "查询用户信息完成", users)
}

// UserActive 用户激活接口
func UserActive(c *gin.Context) {
	//g := newGin(c)
}

// todo 仅管理员
// UpdateUserInfoById 更新用户信息
func UpdateUserInfoById(c *gin.Context) {}

// todo 仅管理员
// DeleteUserInfoById 删除用户信息
func DeleteUserInfoById(c *gin.Context) {}
