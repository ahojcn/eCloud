package service

import (
	"fmt"
	"time"

	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
)

// CreateTreeWithUser 用户创建树
func CreateTreeWithUser(user *model.User, data entity.CreateTreeNodeRequestData) error {
	// 判断用户是否有这个 parent_id 节点的新增权限（4）
	if *data.ParentId != 0 {
		ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": data.ParentId})
		if !has {
			return fmt.Errorf("没有权限")
		}
		if ut.Rights < 4 || ut.Rights > 6 {
			return fmt.Errorf("没有权限")
		}
	}

	// 判断用户是否有同类型同名的节点
	tl, err := model.TreeList(map[string]interface{}{"parent_id": data.ParentId, "name": data.Name})
	if err != nil {
		return fmt.Errorf("内部错误:%v", err)
	}
	if len(tl) != 0 {
		return fmt.Errorf("在节点%v下已经有节点名为%v的节点了", *data.ParentId, data.Name)
	}

	// 判断 parent_id 是否存在
	_, has := model.TreeOne(map[string]interface{}{"id": data.ParentId})
	if !has {
		return fmt.Errorf("父节点不存在")
	}

	// 添加节点
	tree := &model.Tree{
		Name:        data.Name,
		Description: data.Description,
		Type:        *data.Type,
		ParentId:    *data.ParentId,
	}
	err = model.TreeAdd(tree)
	if err != nil {
		return fmt.Errorf("创建失败")
	}

	ut := &model.UserTree{
		UserId:     user.Id,
		TreeId:     tree.Id,
		Rights:     6,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = model.UserTreeAdd(ut)
	if err != nil {
		// 删除上面创建的树节点
		_ = model.TreeDelete(tree.Id, tree)
		return fmt.Errorf("创建失败")
	}

	return nil
}

// GetAllTreeNodeByUser 获取用户所有有权限的节点信息，组织称树状
func GetAllTreeNodeByUser(user *model.User) ([]*entity.TreeNodeInfo, error) {
	uts, err := model.UserTreeList(map[string]interface{}{"user_id": user.Id})
	if err != nil {
		return nil, fmt.Errorf("所有有权限的服务树失败:err=%s", err)
	}
	rdata := []*entity.TreeNodeInfo{}
	for _, ut := range uts {
		uti := ut.UserTree2UserTreeInfo()
		if uti.Tree.Un == "" {
			uti.Tree.Un = uti.Name
			_ = model.TreeUpdate(uti.Tree.Id, uti.Tree)
		}
		if uti.ParentId == 0 {
			rdata = append(rdata, &entity.TreeNodeInfo{
				UserTreeInfo: uti,
				Children:     buildTree(uti, user),
			})
		}
	}
	return rdata, nil
}

func buildTree(uti *model.UserTreeInfo, user *model.User) []*entity.TreeNodeInfo {
	ts, err := model.TreeList(map[string]interface{}{"parent_id": uti.Id})
	if err != nil {
		return []*entity.TreeNodeInfo{}
	}
	rdata := []*entity.TreeNodeInfo{}
	for _, t := range ts {
		ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": t.Id})
		if !has {
			tmpUt := &model.UserTree{
				UserId:     user.Id,
				TreeId:     t.Id,
				Rights:     uti.Rights,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			}
			tmpUdi := tmpUt.UserTree2UserTreeInfo()
			_, _ = model.GetMaster().Insert(tmpUt)
			bt := buildTree(tmpUdi, user)
			nodes := &entity.TreeNodeInfo{UserTreeInfo: tmpUdi, Children: bt}
			rdata = append(rdata, nodes)
			continue
		}

		if t.Un == "" {
			t.Un = t.Name + "." + uti.Un
			_ = model.TreeUpdate(t.Id, &t)
		}

		uti := ut.UserTree2UserTreeInfo()
		rdata = append(rdata, &entity.TreeNodeInfo{
			UserTreeInfo: uti,
			Children:     buildTree(uti, user),
		})
	}
	return rdata
}

// GetTreeNodesDetailById 根据 id 获取 tree 的详细信息
func GetTreeNodesDetailById(nodeId int64, user *model.User) (*entity.TreeNodeDetail, error) {
	res := new(entity.TreeNodeDetail)
	utl, err := model.UserTreeList(map[string]interface{}{"tree_id": nodeId})
	if err != nil {
		return nil, fmt.Errorf("获取用户-服务树信息失败，err=%v", err)
	}
	for _, ut := range utl {
		res.Users = append(res.Users, *ut.UserTree2UserTreeInfo())
	}

	uti, has := model.UserTreeOne(map[string]interface{}{"tree_id": nodeId, "user_id": user.Id})
	if !has {
		return nil, fmt.Errorf("获取服务树信息失败，err=%v", err)
	}
	res.UserTreeInfo = uti.UserTree2UserTreeInfo()
	res.Children = buildTree(res.UserTreeInfo, user)
	return res, nil
}

// GetTreeNodeInfoByName 根据服务树 name 获取服务树节点信息，模糊匹配
func GetTreeNodeInfoByName(name string, user *model.User) ([]model.Tree, error) {
	return model.TreeInfoByNodeNameOrDesc(name)
}

// AddUserTree 给一个用户添加权限
func AddUserTree(user *model.User, userId, treeId int64, rights int) error {
	ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": treeId})
	if !has {
		return fmt.Errorf("你没有权限")
	}
	if ut.Rights != 6 {
		return fmt.Errorf("你没有权限")
	}

	if rights < 0 || rights > 6 {
		return fmt.Errorf("权限错误:0-6")
	}

	_, has = model.UserOne(map[string]interface{}{"id": userId})
	if !has {
		return fmt.Errorf("没有这个用户")
	}
	_, has = model.TreeOne(map[string]interface{}{"id": treeId})
	if !has {
		return fmt.Errorf("没有这个用户")
	}

	ut, has = model.UserTreeOne(map[string]interface{}{"user_id": userId, "tree_id": treeId})
	if has {
		ut.Rights = rights
		if err := model.UserTreeUpdate(ut.Id, ut); err != nil {
			return fmt.Errorf("这个用户已有权限:%v，但是更新权限失败了:%v", ut.RightsMsg(), err)
		}
		return nil
	}

	if err := model.UserTreeAdd(&model.UserTree{
		UserId: userId,
		TreeId: treeId,
		Rights: rights,
	}); err != nil {
		return fmt.Errorf("添加失败:%v", err)
	}

	return nil
}

// DeleteUserTree 删除一个用户的节点权限
func DeleteUserTree(user *model.User, rd *entity.DeleteUserTreeRequestData) error {
	if user.Id == *rd.UserId {
		return fmt.Errorf("无权限删除自己")
	}

	ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": *rd.TreeId})
	if !has {
		return fmt.Errorf("没有权限删除")
	}
	if ut.Rights != model.PermAdmin {
		return fmt.Errorf("没有权限删除")
	}

	ut, has = model.UserTreeOne(map[string]interface{}{"user_id": *rd.UserId, "tree_id": *rd.TreeId})
	if !has {
		return nil
	}
	return model.UserTreeDelete(ut.Id, ut)
}

// DeleteTreeNode 标记删除一个节点
func DeleteTreeNode(user *model.User, rd *entity.DeleteTreeNodeRequestData) error {
	ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": *rd.TreeId})
	if !has {
		return fmt.Errorf("没有这个节点")
	}
	if ut.Rights < model.PermDelete || ut.Rights > model.PermAdmin {
		return fmt.Errorf("无权限删除")
	}

	t, has := model.TreeOne(map[string]interface{}{"id": *rd.TreeId})
	if !has {
		return fmt.Errorf("没有节点信息")
	}
	return model.TreeMarkDelete(t)
}
