package model

import (
	"errors"
	"time"
)

type Tree struct {
	Id          int64     `json:"id" xorm:"pk autoincr notnull"`
	Name        string    `json:"name" xorm:"varchar(128) notnull default ''"`
	Description string    `json:"description" xorm:"varchar(1024) notnull default ''"`
	Type        int       `json:"type" xorm:"notnull default 0"`
	ParentId    int64     `json:"parent_id" xorm:"notnull default 0"`
	CreateTime  time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime  time.Time `json:"update_time" xorm:"notnull updated"`
	Extra       string    `json:"extra" xorm:"varchar(1024) notnull default '{}'"`
}

// TypeMsg 获取 tree 中 type 的文字描述
func (t Tree) TypeMsg() string {
	m := map[int]string{
		0: "cluster",
		1: "service",
		2: "group",
		3: "pdl",
		4: "namespace",
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