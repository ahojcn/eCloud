package model

import "fmt"

type Cluster struct {
	Id                int64 `json:"id" xorm:"pk autoincr notnull"`
	TreeID            int64 `json:"tree_id" xorm:"notnull"`
	ClusterNum        int   `json:"cluster_num" xorm:"notnull default 0"` // 集群数量，用户指定
	CurrentClusterNum int   `json:"current_cluster_num" xorm:"notnull default 0"`
	ContainerPort     int   `json:"container_port" xorm:"notnull"` // 容器端口，用户指定
}

func (c *Cluster) GetTree() (*Tree, error) {
	t, has := TreeOne(map[string]interface{}{"id": c.TreeID})
	if !has {
		return nil, fmt.Errorf("没有这个服务树节点信息,treeid=%d", c.TreeID)
	}
	return t, nil
}

func ClusterAdd(c *Cluster) error {
	orm := GetMaster()
	affected, err := orm.Insert(c)
	if affected == 0 {
		return fmt.Errorf("insert failed, affected = 0")
	}
	return err
}

func ClusterOne(cons map[string]interface{}) (*Cluster, bool) {
	orm := GetSlave()

	t := new(Cluster)
	has, err := orm.Where(cons).Get(t)
	if err != nil {
		return nil, false
	}
	return t, has
}

func ClusterList(cons map[string]interface{}) ([]Cluster, error) {
	orm := GetSlave()

	clusters := make([]Cluster, 0)
	err := orm.Where(cons).Find(&clusters)
	if err != nil {
		return nil, err
	}
	return clusters, err
}

func ClusterDelete(c *Cluster) error {
	orm := GetMaster()
	affected, err := orm.Delete(c)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}

type ClusterContainer struct {
	Id          int64 `json:"id" xorm:"pk autoincr notnull"`
	ClusterID   int64 `json:"cluster_id" xorm:"notnull"`
	ContainerID int64 `json:"container_id" xorm:"notnull"`
}

func (c *ClusterContainer) GetContainers(cons map[string]interface{}) ([]*Container, error) {
	orm := GetSlave()

	clusterContainers := make([]ClusterContainer, 0)
	err := orm.Where(cons).Find(&clusterContainers)
	if err != nil {
		return nil, err
	}

	containers := make([]*Container, 0)
	for _, clusterContainer := range clusterContainers {
		c, _ := ContainerOne(map[string]interface{}{"id": clusterContainer.ContainerID})
		containers = append(containers, c)
	}

	return containers, err
}

func ClusterContainerAdd(c *ClusterContainer) error {
	orm := GetMaster()
	affected, err := orm.Insert(c)
	if affected == 0 {
		return fmt.Errorf("insert failed, affected = 0")
	}
	return err
}

func ClusterContainerList(cons map[string]interface{}) ([]ClusterContainer, error) {
	orm := GetSlave()

	clusters := make([]ClusterContainer, 0)
	err := orm.Where(cons).Find(&clusters)
	if err != nil {
		return nil, err
	}
	return clusters, err
}

func ClusterContainerDelete(c *ClusterContainer) error {
	orm := GetMaster()
	affected, err := orm.Delete(c)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}
