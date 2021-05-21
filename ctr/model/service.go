package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Service struct {
	Id            int64  `json:"id" xorm:"pk autoincr notnull"`
	TreeId        int64  `json:"tree_id" xorm:"notnull"`
	RouterIp      string `json:"router_ip" xorm:"varchar(128) notnull"`
	RouterPort    int    `json:"router_port" xorm:"notnull"`
	FlowMap       string `json:"flow_map" xorm:"text notnull"`
	ConfigContent string `json:"config_content" xorm:"text notnull"`
}

type ServiceInfo struct {
	*Service
	TreeUnFlowMap map[string]int         `json:"tree_un_flow_map"`
	ChartOpt      map[string]interface{} `json:"chart_opt"`
}

func (s *Service) GetServiceInfo() (*ServiceInfo, error) {
	flowMap := map[string]int{}
	if err := json.Unmarshal([]byte(s.FlowMap), &flowMap); err != nil {
		return nil, fmt.Errorf("预案配置解析错误,flow_map=%v,err=%v", s.FlowMap, err)
	}

	svcTree, _ := TreeOne(map[string]interface{}{"id": s.TreeId})
	opt := map[string]interface{}{
		"name":  svcTree.Un,
		"value": 100,
	}

	treeUnFlowMap := map[string]int{}
	children := []map[string]interface{}{}
	for treeIdStr := range flowMap {
		treeId, _ := strconv.Atoi(treeIdStr)
		t, _ := TreeOne(map[string]interface{}{"id": treeId})
		treeUnFlowMap[t.Un] = flowMap[treeIdStr]
		child := map[string]interface{}{
			"name":  t.Un,
			"value": flowMap[treeIdStr],
		}
		children = append(children, child)
	}

	opt["children"] = children
	return &ServiceInfo{
		Service:       s,
		TreeUnFlowMap: treeUnFlowMap,
		ChartOpt:      opt,
	}, nil
}

func ServiceAdd(c *Service) error {
	orm := GetMaster()
	affected, err := orm.Insert(c)
	if affected == 0 {
		return fmt.Errorf("insert failed, affected = 0")
	}
	return err
}

func ServiceOne(cons map[string]interface{}) (*Service, bool) {
	orm := GetSlave()

	t := new(Service)
	has, err := orm.Where(cons).Get(t)
	if err != nil {
		return nil, false
	}
	return t, has
}

func ServiceList(cons map[string]interface{}) ([]Service, error) {
	orm := GetSlave()

	services := make([]Service, 0)
	err := orm.Where(cons).Find(&services)
	if err != nil {
		return nil, err
	}
	return services, err
}

func ServiceUpdate(id int64, s *Service) error {
	orm := GetMaster()
	_, err := orm.ID(id).Update(s)
	return err
}

func ServiceDelete(p *Service) error {
	orm := GetMaster()
	affected, err := orm.Delete(p)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}
