package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"net/http"
)

func ClusterRetrieve(user *model.User, rd *entity.ClusterRetrieveRequestData) (int, *model.ClusterInfo, error) {
	c, has := model.ClusterOne(map[string]interface{}{"tree_id": *rd.TreeID})
	if !has {
		return http.StatusNotFound, nil, fmt.Errorf("不存在的集群配置")
	}

	return http.StatusOK, c.GetClusterInfo(), nil
}

func ClusterDelete(user *model.User, rd *entity.ClusterRetrieveRequestData) (int, error) {
	c, has := model.ClusterOne(map[string]interface{}{"tree_id": *rd.TreeID})
	if !has {
		return http.StatusNotFound, fmt.Errorf("不存在的集群配置")
	}

	err := model.ClusterDelete(c)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("删除集群配置失败,err=%v", err)
	}

	return http.StatusOK, nil
}

func ClusterCreate(user *model.User, rd *entity.ClusterCreateRequestData) (int, *model.Cluster, error) {
	_, has := model.ClusterOne(map[string]interface{}{"tree_id": *rd.TreeID})
	if has {
		return http.StatusBadRequest, nil, fmt.Errorf("此集群已存在配置")
	}
	c := &model.Cluster{
		TreeId:        *rd.TreeID,
		ClusterNum:    *rd.ClusterNum,
		ContainerPort: *rd.ContainerPort,
	}
	if err := model.ClusterAdd(c); err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("创建集群配置失败, err=%v", err)
	}
	return http.StatusOK, c, nil
}

func ClusterList(user *model.User, rd *entity.ClusterListRequestData) (int, []*entity.TreeNodeInfo, error) {
	ut, has := model.UserTreeOne(map[string]interface{}{"user_id": user.Id, "tree_id": *rd.TreeID})
	if !has {
		return http.StatusUnauthorized, nil, fmt.Errorf("没有此节点权限")
	}
	treeNodeInfoList := buildTree(ut.UserTree2UserTreeInfo(), user)
	fmt.Println(treeNodeInfoList)
	return http.StatusOK, treeNodeInfoList, nil
}
