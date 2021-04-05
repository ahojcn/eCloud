package model

import "time"

type UserTree struct {
	Id         int64     `json:"id" xorm:"pk autoincr notnull"`
	UserId     int64     `json:"user_id" xorm:"notnull default 0"`
	TreeId     int64     `json:"tree_id" xorm:"notnull default 0"`
	Rights     int       `json:"rights" xorm:"notnull default 0"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
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
	}
	return m[ut.Rights]
}
