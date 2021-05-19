package model

import (
	"fmt"
	"time"
)

type Container struct {
	Id            int64     `json:"id" xorm:"pk autoincr notnull"`
	HostId        int64     `json:"host_id" xorm:"notnull"`
	ContainerId   string    `json:"container_id" xorm:"varchar(128) notnull"`
	ContainerIp   string    `json:"container_ip" xorm:"varchar(64) notnull"`
	ContainerPort int       `json:"container_port" xorm:"notnull"`
	HostPort      int       `json:"host_port" xorm:"notnull"`
	CreateTime    time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime    time.Time `json:"update_time" xorm:"notnull updated"`
}

func (c *Container) GetHost() (*Host, error) {
	h, has := HostOne(map[string]interface{}{"id": c.HostId})
	if !has {
		return nil, fmt.Errorf("获取主机失败,err:不存在的主机%d,containerId:%v", c.HostId, c.Id)
	}
	return h, nil
}

func ContainerAdd(c *Container) error {
	orm := GetMaster()
	affected, err := orm.Insert(c)
	if affected == 0 {
		return fmt.Errorf("insert failed, affected = 0")
	}
	return err
}

func ContainerOne(cons map[string]interface{}) (*Container, bool) {
	orm := GetSlave()
	c := new(Container)
	has, err := orm.Where(cons).Get(c)
	if err != nil {
		return nil, false
	}
	return c, has
}

func ContainerList(cons map[string]interface{}) ([]Container, error) {
	orm := GetSlave()

	containers := make([]Container, 0)
	err := orm.Where(cons).Find(&containers)
	if err != nil {
		return nil, err
	}
	return containers, err
}

func ContainerDelete(host *Host) error {
	orm := GetMaster()
	affected, err := orm.Delete(host)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}
