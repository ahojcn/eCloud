package model

import (
	"errors"
	"github.com/ahojcn/ecloud/ctr/entity"
)

func UserAdd(user *entity.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("username||password||email can't be empty")
	}

	orm := GetMaster()

	affected, err := orm.Insert(user)
	if err != nil {
		// todo log
		return err
	}
	if affected == 0 {
		// todo log
		return errors.New("insert user failed, affected = 0")
	}

	return nil
}

func UserOne(cons map[string]interface{}) (*entity.User, bool) {
	orm := GetSlave()

	user := new(entity.User)
	has, err := orm.Where(cons).Get(user)
	if err != nil {
		// todo log
	}
	return user, has
}

func UserList(cons map[string]interface{}) ([]entity.User, error) {
	orm := GetSlave()

	users := make([]entity.User, 0)
	err := orm.Where(cons).Find(&users)
	if err != nil {
		// todo log
	}
	return users, err
}
