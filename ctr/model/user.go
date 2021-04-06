package model

import (
	"errors"
	"time"
)

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

type UserInfo struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	IsActive   string    `json:"is_active"`
	Extra      string    `json:"extra"`
}

func (u User) User2UserInfo() *UserInfo {
	return &UserInfo{
		Id:         u.Id,
		Username:   u.Username,
		Email:      u.Email,
		Phone:      u.Phone,
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
		IsActive:   u.IsActiveMsg(),
		Extra:      u.Extra,
	}
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

func UserAdd(user *User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("username||password||email can't be empty")
	}

	orm := GetMaster()

	affected, err := orm.Insert(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("insert failed, affected = 0")
	}

	return nil
}

func UserOne(cons map[string]interface{}) (*User, bool) {
	orm := GetSlave()

	user := new(User)
	has, err := orm.Where(cons).Get(user)
	if err != nil {
		return nil, false
	}
	return user, has
}

func UserList(cons map[string]interface{}) ([]User, error) {
	orm := GetSlave()

	users := make([]User, 0)
	err := orm.Where(cons).Find(&users)
	if err != nil {
		// todo log
	}
	return users, err
}
