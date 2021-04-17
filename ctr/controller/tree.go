package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTreeNode 创建服务树节点
func CreateTreeNode(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	data := entity.CreateTreeNodeRequestData{}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if err = service.CreateTreeWithUser(user, data); err != nil {
		g.response(http.StatusInternalServerError, "创建失败", err)
		return
	}

	g.response(http.StatusOK, "创建成功", nil)
}

// GetTreeNodes 获取有权限的树节点
func GetTreeNodes(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	rd := entity.GetTreeNodeRequestData{}
	if err = c.ShouldBindQuery(&rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	if rd.Id != nil {
		treeNodeDetail, err := service.GetTreeNodesDetailById(*rd.Id, user)
		if err != nil {
			g.response(http.StatusInternalServerError, "服务器错误", err)
			return
		}
		g.response(http.StatusOK, "ok", treeNodeDetail)
		return
	} else if rd.Name != nil {
		treeList, err := service.GetTreeNodeInfoByName(*rd.Name, user)
		if err != nil {
			g.response(http.StatusInternalServerError, "服务器错误", err)
			return
		}
		g.response(http.StatusOK, "ok", treeList)
		return
	}

	rdata, err := service.GetAllTreeNodeByUser(user)
	if err != nil {
		g.response(http.StatusInternalServerError, "服务器错误", err)
		return
	}

	g.response(http.StatusOK, "ok", rdata)
	return
}

// CreateUserTree 给用户添加权限
func CreateUserTree(c *gin.Context) {
	g := newGin(c)
	user, err := g.loginRequired()
	if err != nil {
		g.response(http.StatusUnauthorized, "未登录", err)
		return
	}

	// 校验参数
	rd := entity.CreateUserTreeRequestData{}
	if err = c.ShouldBindJSON(&rd); err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	// 添加节点权限
	if err = service.AddUserTree(user, *rd.UserId, *rd.TreeId, *rd.Rights); err != nil {
		g.response(http.StatusInternalServerError, "服务器错误", err)
		return
	}

	g.response(http.StatusOK, "ok", nil)
}
