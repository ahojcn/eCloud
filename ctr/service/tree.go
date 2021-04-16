package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"time"
)

// GetAllTreeNodeByUser 获取用户所有有权限的节点信息，组织称树状
func GetAllTreeNodeByUser(user *model.User) ([]*entity.TreeNodeInfo, error) {
	uts, err := model.UserTreeList(map[string]interface{}{"user_id": user.Id})
	if err != nil {
		return nil, fmt.Errorf("所有有权限的服务树失败:err=%s", err)
	}
	rdata := []*entity.TreeNodeInfo{}
	for _, ut := range uts {
		uti := ut.UserTree2UserTreeInfo()
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
		u, has := model.UserOne(map[string]interface{}{"id": ut.UserId})
		if !has {
			continue
		}
		res.Users = append(res.Users, u.User2UserInfo())
	}

	uti, has := model.UserTreeOne(map[string]interface{}{"tree_id": nodeId, "user_id": user.Id})
	if !has {
		return nil, fmt.Errorf("获取服务树信息失败，err=%v", err)
	}
	res.UserTreeInfo = uti.UserTree2UserTreeInfo()
	res.Children = buildTree(res.UserTreeInfo, user)
	return res, nil
}
