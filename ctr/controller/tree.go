package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTreeNode 创建服务树节点
func CreateTreeNode(c *gin.Context) {
	// todo 判断用户是否有同类型同名的节点
	// todo 判断用户是否有这个 parent_id 节点的新增权限（4）
	g := newGin(c)
	data := entity.CreateTreeNodeRequestData{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		g.response(http.StatusBadRequest, "参数错误", err)
		return
	}

	// 判断 parent_id 是否存在
	_, has := model.TreeOne(map[string]interface{}{"id": data.ParentId})
	if !has {
		g.response(http.StatusBadRequest, "父节点不存在", data.ParentId)
		return
	}

	// 添加节点
	// todo 在 user_tree 中添加记录
	tree := &model.Tree{
		Name:        data.Name,
		Description: data.Description,
		Type:        data.Type,
		ParentId:    data.ParentId,
	}
	err = model.TreeAdd(tree)
	if err != nil {
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
	}

	rdata, err := service.GetAllTreeNodeByUser(user)
	if err != nil {
		g.response(http.StatusInternalServerError, "服务器错误", err)
		return
	}

	g.response(http.StatusOK, "ok", rdata)
	return
}
