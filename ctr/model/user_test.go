package model

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"testing"
)

func TestUserAdd(t *testing.T) {
	u1 := entity.User{
		//Username:   "ahoj2",
		Password:   "ahoj2",
		Email:      "ahoj2@qq.com",
		Phone:      "",
		IsActive:   0,
	}
	t.Log(UserAdd(&u1), u1.Id)
}

func TestUserOne(t *testing.T) {
	u, has := UserOne(map[string]interface{}{})
	t.Log(has, u.Username)
}

func TestUserList(t *testing.T) {
	users, err := UserList(map[string]interface{}{"username": "ahoj1"})
	t.Log(len(users), err)
	t.Log(users)
}
