package model

import (
	"errors"
	"time"
)

type HostUser struct {
	Id         int64     `json:"id" xorm:"pk autoincr notnull"`
	UserId     int64     `json:"user_id" xorm:"notnull default 0"`
	HostId     int64     `json:"host_id" xorm:"notnull default 0"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
}

func (hu *HostUser) GetHost() *Host {
	h, _ := HostOne(map[string]interface{}{"id": hu.HostId})
	return h
}

func (hu *HostUser) GetUser() *User {
	u, _ := UserOne(map[string]interface{}{"id": hu.UserId})
	return u
}

func HostUserAdd(hu *HostUser) error {
	orm := GetMaster()
	affected, err := orm.Insert(hu)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("insert failed, affected = 0")
	}
	return nil
}

func HostUserOne(cons map[string]interface{}) (*HostUser, bool) {
	orm := GetSlave()
	hu := new(HostUser)
	has, err := orm.Where(cons).Get(hu)
	if err != nil {
		return nil, false
	}
	return hu, has
}

func HostUserList(cons map[string]interface{}) ([]HostUser, error) {
	orm := GetSlave()
	hus := make([]HostUser, 0)
	err := orm.Where(cons).Find(&hus)
	if err != nil {
		return nil, err
	}
	return hus, err
}

func HostUserDelete(hu *HostUser) error {
	orm := GetMaster()
	affected, err := orm.Delete(hu)
	if affected == 0 {
		return errors.New("delete failed, affected = 0")
	}
	return err
}
