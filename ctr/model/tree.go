package model

import (
	"errors"
	"fmt"
	"time"
)

type Tree struct {
	Id          int64     `json:"id" xorm:"pk autoincr notnull"`
	Name        string    `json:"name" xorm:"varchar(128) notnull default ''"`
	Un          string    `json:"un" xorm:"text notnull"`
	Description string    `json:"description" xorm:"varchar(1024) notnull default ''"`
	Type        int       `json:"type" xorm:"notnull default 0"`
	ParentId    int64     `json:"parent_id" xorm:"notnull default 0"`
	CreateTime  time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime  time.Time `json:"update_time" xorm:"notnull updated"`
	Extra       string    `json:"extra" xorm:"varchar(1024) notnull default '{}'"`
	IsDeleted   bool      `json:"is_deleted" xorm:"default false"`
}

const (
	TreeTypeNamespace   = 4
	TreeTypeProductLine = 3
	TreeTypeGroup       = 2
	TreeTypeService     = 1
	TreeTypeCluster     = 0
)

// TypeMsg 获取 tree 中 type 的文字描述
func (t Tree) TypeMsg() string {
	m := map[int]string{
		TreeTypeCluster:     "cluster",
		TreeTypeService:     "service",
		TreeTypeGroup:       "group",
		TreeTypeProductLine: "pdl",
		TreeTypeNamespace:   "namespace",
	}
	return m[t.Type]
}

func TreeOne(cons map[string]interface{}) (*Tree, bool) {
	orm := GetSlave()

	t := new(Tree)
	has, err := orm.Where(cons).Get(t)
	if err != nil {
		return nil, false
	}
	return t, has
}

func TreeAdd(tree *Tree) error {
	orm := GetMaster()
	affected, err := orm.Insert(tree)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("insert failed, affected = 0")
	}
	return nil
}

func TreeList(cons map[string]interface{}) ([]Tree, error) {
	orm := GetSlave()
	ts := make([]Tree, 0)
	err := orm.Where(cons).Find(&ts)
	if err != nil {
		return nil, err
	}
	return ts, err
}

func TreeInfoByNodeNameOrDesc(name string) ([]Tree, error) {
	orm := GetSlave()
	tl := make([]Tree, 0)
	err := orm.Where("name like ?", "%"+name+"%").Or("description like ?", "%"+name+"%").Find(&tl)
	if err != nil {
		return nil, err
	}
	return tl, nil
}

func TreeUpdate(id int64, t *Tree) error {
	orm := GetMaster()
	affected, err := orm.ID(id).Update(t)
	if affected == 0 {
		return fmt.Errorf("update failed, affected = 0")
	}
	return err
}

func TreeDelete(id int64, t *Tree) error {
	orm := GetSlave()
	affected, err := orm.ID(id).Delete(t)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("update failed, affected = 0")
	}
	return nil
}

func TreeMarkDelete(t *Tree) error {
	orm := GetMaster()
	t.IsDeleted = true
	affected, err := orm.ID(t.Id).Cols("is_deleted").Update(t)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("update failed, affected = 0")
	}
	return nil
}
