package controller

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/gin-contrib/sessions"
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
	sess := sessions.Default(c)
	username := sess.Get("username")
	if username == nil {
		g.response(http.StatusUnauthorized, "未登录", nil)
		return
	}
	// 获取用户信息
	user, has := model.UserOne(map[string]interface{}{"username": username})
	if !has {
		g.response(http.StatusUnauthorized, "未登录", nil)
		return
	}
	uts, err := model.UserTreeList(map[string]interface{}{"user_id": user.Id})
	if err != nil {
		g.response(http.StatusInternalServerError, "服务器错误", err)
		return
	}

	rdata := []*entity.GetTreeNodesResponseData{}
	for _, ut := range uts {
		uti := ut.UserTree2UserTreeInfo()
		if uti.ParentId == 0 {
			rdata = append(rdata, &entity.GetTreeNodesResponseData{
				UserTreeInfo: uti,
				Children:     buildTree(uti, user),
			})
		}
	}

	g.response(http.StatusOK, "ok", rdata)
	return
}

func buildTree(uti *model.UserTreeInfo, user *model.User) []*entity.GetTreeNodesResponseData {
	ts, err := model.TreeList(map[string]interface{}{"parent_id": uti.Id})
	if err != nil {
		return []*entity.GetTreeNodesResponseData{}
	}
	rdata := []*entity.GetTreeNodesResponseData{}
	for _, t := range ts {
		ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": t.Id})
		if !has {
			continue
		}
		uti := ut.UserTree2UserTreeInfo()

		rdata = append(rdata, &entity.GetTreeNodesResponseData{
			UserTreeInfo: uti,
			Children:     buildTree(uti, user),
		})
	}
	return rdata
}
