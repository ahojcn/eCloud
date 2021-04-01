package entity

import "time"

type User struct {
	Id         int64     `json:"id" xorm:"pk autoincr notnull"`
	Username   string    `json:"username" xorm:"varchar(32) notnull default ''"`
	Password   string    `json:"password" xorm:"varchar(128) notnull default ''"`
	Email      string    `json:"email" xorm:"varchar(128) notnull default ''"`
	Phone      string    `json:"phone" xorm:"varchar(16) notnull default ''"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
	IsActive   int       `json:"is_active" xorm:"notnull default 0"`
	Extra      string    `json:"extra" xorm:"varchar(1024) notnull default '{}'"`
}

// IsActiveMsg 返回 user 对象中的 is_active 的文字描述
func (u User) IsActiveMsg() string {
	m := map[int]string{
		0: "未激活",
		1: "已激活",
		2: "已禁用",
	}
	return m[u.IsActive]
}
