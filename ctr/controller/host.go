package controller

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHost(c *gin.Context) {
	g := newGin(c)
	var err error
	//user, err := g.loginRequired()
	//if err != nil {
	//	g.response(http.StatusUnauthorized, "未登录", err)
	//	return
	//}
	username := "ahojcn"
	user, _ := model.UserOne(map[string]interface{}{"username": username})

	// 检查参数
	rd := entity.CreateHostRequestData{}
	err = c.ShouldBindJSON(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	host := model.Host{
		UserId:      user.Id,
		Description: rd.Description,
		IP:          rd.IP,
		Username:    rd.Username,
		Password:    rd.Password,
		Port:        rd.Port,
	}
	err = model.HostAdd(&host)
	if err != nil {
		g.response(http.StatusInternalServerError, "添加主机信息失败", err)
		return
	}

	res, err := service.DeployAgent(host)
	if err != nil {
		g.response(http.StatusInternalServerError, "部署失败", append(res, err.Error()))
		return
	}

	g.response(http.StatusOK, "添加成功 && 部署 agent 成功", res)
}

func UpdateHost(c *gin.Context) {
	g := newGin(c)

	// 校验参数
	rd := entity.UpdateHostRequestData{}
	err := c.ShouldBind(&rd)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	// todo here
	h := &model.Host{
		Extra:       rd.Extra,
	}
	err = model.HostUpdate(rd.HostId, h)
	fmt.Println(err, rd.HostId, rd.Extra, rd.Description)
	if err != nil {
		g.response(http.StatusInternalServerError, "更新失败", err)
		return
	}

	g.response(http.StatusOK, "更新成功", nil)
}
