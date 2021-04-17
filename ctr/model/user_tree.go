package model

import (
	"errors"
	"time"
)

type UserTree struct {
	Id         int64     `json:"id" xorm:"pk autoincr notnull"`
	UserId     int64     `json:"user_id" xorm:"notnull default 0"`
	TreeId     int64     `json:"tree_id" xorm:"notnull default 0"`
	Rights     int       `json:"rights" xorm:"notnull default 0"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
}

type UserTreeInfo struct {
	*Tree
	UserInfo   *UserInfo `json:"user_info"`
	Rights     int       `json:"rights"`
	RightMsg   string    `json:"right_msg"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
}

func (ut UserTree) UserTree2UserTreeInfo() *UserTreeInfo {
	t, _ := TreeOne(map[string]interface{}{"id": ut.TreeId})
	u, _ := UserOne(map[string]interface{}{"id": ut.UserId})
	return &UserTreeInfo{
		Tree:       t,
		UserInfo:   u.User2UserInfo(),
		Rights:     ut.Rights,
		RightMsg:   ut.RightsMsg(),
		CreateTime: ut.CreateTime,
		UpdateTime: ut.UpdateTime,
	}
}

// RightsMsg 返回 user_tree 中 rights 的文字描述
func (ut UserTree) RightsMsg() string {
	m := map[int]string{
		0: "n 无权限",
		1: "r 只读",
		2: "w 只写",
		3: "rw 可读写",
		4: "c 可新增",
		5: "d 可删除",
		6: "a 管理",
	}
	return m[ut.Rights]
}

func UserTreeAdd(ut *UserTree) error {
	orm := GetMaster()
	affected, err := orm.Insert(ut)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("insert failed, affected = 0")
	}
	return nil
}

func UserTreeOne(cons map[string]interface{}) (*UserTree, bool) {
	orm := GetSlave()
	ut := new(UserTree)
	has, err := orm.Where(cons).Get(ut)
	if err != nil {
		return nil, false
	}
	return ut, has
}

func UserTreeList(cons map[string]interface{}) ([]UserTree, error) {
	orm := GetSlave()
	uts := make([]UserTree, 0)
	err := orm.Where(cons).Find(&uts)
	if err != nil {
		return nil, err
	}
	return uts, err
}

func UserTreeUpdate(id int64, ut *UserTree) error {
	orm := GetMaster()
	affected, err := orm.ID(id).Update(ut)
	if affected == 0 {
		return errors.New("update failed, affected = 0")
	}
	return err
}