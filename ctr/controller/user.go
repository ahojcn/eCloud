package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
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
		Username:   requestData.Username,
		Password:   util.Md5Str(requestData.Password),
		Email:      requestData.Email,
		IsActive:   0,
	}
	err = model.UserAdd(user)
	if err != nil {
		g.response(http.StatusInternalServerError, "注册失败", err)
		return
	}

	g.response(http.StatusOK, "ok", user.User2UserInfo())
}

// 获取用户信息
func GetUserInfoById(c *gin.Context) {
	g := newGin(c)

	var rd entity.UserInfoByIdRequestData
	err := c.ShouldBindUri(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	user, has := model.UserOne(map[string]interface{}{"id": rd.Id})
	if !has {
		g.response(http.StatusNotFound, "未找到", rd.Id)
		return
	}

	g.response(http.StatusOK, "ok", user.User2UserInfo())
}

// todo 仅管理员
// 更新用户信息
func UpdateUserInfoById(c *gin.Context) {}

// todo 仅管理员
// 删除用户信息
func DeleteUserInfoById(c *gin.Context) {}
