package model

import (
	"fmt"
	"time"
)

type ICode struct {
	Id            int64     `json:"id" xorm:"pk autoincr notnull"`
	UserId        int64     `json:"user_id" xorm:"notnull"`
	HostId        int64     `json:"host_id" xorm:"notnull"`
	Name          string    `json:"name" xorm:"varchar(32) notnull"`
	Port          int       `json:"port" xorm:"notnull"`
	Password      string    `json:"password" xorm:"varchar(128) notnull"`
	ContainerId   string    `json:"container_id" xorm:"varchar(128) notnull"`
	ContainerIP   string    `json:"container_ip" xorm:"varchar(64) notnull"`
	ContainerPort int       `json:"container_port" xorm:"notnull"`
	CreateTime    time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime    time.Time `json:"update_time" xorm:"notnull updated"`
	Extra         string    `json:"extra" xorm:"text notnull"`
}

func (i *ICode) GetUser() *User {
	u, _ := UserOne(map[string]interface{}{"id": i.UserId})
	return u
}

func (i *ICode) GetHost() *Host {
	h, _ := HostOne(map[string]interface{}{"id": i.HostId})
	return h
}

func ICodeAdd(ic *ICode) error {
	orm := GetMaster()
	affected, err := orm.Insert(ic)
	if affected == 0 {
		return fmt.Errorf("insert failed, affected = 0")
	}
	return err
}

func ICodeOne(cons map[string]interface{}) (*ICode, bool) {
	orm := GetSlave()
	i := new(ICode)
	has, err := orm.Where(cons).Get(i)
	if err != nil {
		return nil, false
	}
	return i, has
}

func ICodeList(cons map[string]interface{}) ([]ICode, error) {
	orm := GetSlave()
	iCodes := make([]ICode, 0)
	err := orm.Where(cons).Find(&iCodes)
	if err != nil {
		return nil, err
	}
	return iCodes, err
}

func ICodeDelete(i *ICode) error {
	orm := GetMaster()
	affected, err := orm.Delete(i)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}
