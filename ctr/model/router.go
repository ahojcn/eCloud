package model

import (
	"fmt"
	"time"
)

type Router struct {
	Id         int64     `json:"id" xorm:"pk autoincr notnull"`
	NsId       int64     `json:"ns_id" xorm:"notnull default 0"`
	HostId     int64     `json:"host_id" xorm:"notnull default 0"`
	Log        string    `json:"log" xorm:"text notnull"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
}

type RouterInfo struct {
	Router   *Router `json:"router"`
	HostInfo *Host   `json:"host_info"`
	NsInfo   *Tree   `json:"ns_info"`
}

func (r *Router) GetRouterInfo() *RouterInfo {
	hi, _ := r.GetHostInfo()
	ni, _ := r.GetNsInfo()
	return &RouterInfo{
		Router:   r,
		HostInfo: hi,
		NsInfo:   ni,
	}
}

func (r *Router) GetNsInfo() (*Tree, error) {
	t, has := TreeOne(map[string]interface{}{"id": r.NsId})
	if !has {
		return nil, fmt.Errorf("get ns info failed, has=false")
	}
	return t, nil
}

func (r *Router) GetHostInfo() (*Host, error) {
	h, has := HostOne(map[string]interface{}{"id": r.HostId})
	if !has {
		return nil, fmt.Errorf("get host info failed, has=false")
	}
	return h, nil
}

func RouterAdd(r *Router) error {
	orm := GetMaster()
	_, err := orm.Insert(r)
	return err
}

func RouterUpdate(id int64, r *Router) error {
	orm := GetMaster()
	_, err := orm.ID(id).Update(r)
	return err
}

func RouterOne(cons map[string]interface{}) (*Router, bool) {
	orm := GetSlave()
	r := new(Router)
	has, err := orm.Where(cons).Get(r)
	if err != nil {
		return nil, false
	}
	return r, has
}

func RouterList(cons map[string]interface{}) ([]Router, error) {
	orm := GetSlave()

	r := make([]Router, 0)
	err := orm.Where(cons).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, err
}

func RouterDelete(r *Router) error {
	orm := GetMaster()
	affected, err := orm.Delete(r)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}
