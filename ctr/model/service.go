package model

import "fmt"

type Service struct {
	Id      int64  `json:"id" xorm:"pk autoincr notnull"`
	TreeID  int64  `json:"tree_id" xorm:"notnull"`
	FlowMap string `json:"flow_map" xorm:"text notnull"`
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

func ServiceDelete(p *Service) error {
	orm := GetMaster()
	affected, err := orm.Delete(p)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}
