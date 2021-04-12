package controller

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录
func Login(c *gin.Context) {
	g := newGin(c)
	sess := sessions.Default(c)

	// 判断是否已经登录
	if username := sess.Get("username"); username != nil {
		if user, has := model.UserOne(map[string]interface{}{"username": username}); !has {
			sess.Clear()
			g.response(http.StatusUnauthorized, "用户信息有误，请重新登录", nil)
			return
		} else {
			sess.Set("username", username)
			sess.Options(sessions.Options{MaxAge: 60 * 60 * 24})
			_ = sess.Save()
			g.response(http.StatusOK, "已登录", user.User2UserInfo())
			return
		}
	}

	var data entity.LoginRequestData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	// 获取用户信息，校验用户名密码
	user, has := model.UserOne(map[string]interface{}{"username": data.Username})
	if !has {
		g.response(http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}
	if user.Password != util.Md5Str(data.Password) {
		g.response(http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 登录成功，将 username 设置到 session 中
	sess.Set("username", user.Username)
	sess.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	_ = sess.Save()
	g.response(http.StatusOK, fmt.Sprintf("欢迎回来，%s", user.Username), user.User2UserInfo())
}

// Logout 退出
func Logout(c *gin.Context) {
	g := newGin(c)
	sess := sessions.Default(c)
	sess.Clear()
	_ = sess.Save()
	g.response(http.StatusOK, "再见", nil)
}

func IsLogin(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	g.response(http.StatusOK, "已登录", user.User2UserInfo())
}
