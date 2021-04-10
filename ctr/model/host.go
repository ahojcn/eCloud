package model

import (
	"errors"
	"github.com/ahojcn/ecloud/ctr/util"
	"time"
)

type Host struct {
	Id          int64     `json:"id" xorm:"pk autoincr notnull"`
	UserId      int64     `json:"user_id" xorm:"notnull default 0"`
	Description string    `json:"description" xorm:"varchar(1024) notnull default ''"`
	IP          string    `json:"ip" xorm:"varchar(20) notnull default ''"`
	Username    string    `json:"username" xorm:"varchar(128) notnull default ''"`
	Password    string    `json:"password" xorm:"varchar(128) notnull default ''"`
	Port        int       `json:"port" xorm:"notnull default 22"`
	CreateTime  time.Time `json:"create_time" xorm:"notnull created"`
	UpdateTime  time.Time `json:"update_time" xorm:"notnull updated"`
	Extra       string    `json:"extra" xorm:"text notnull"`
}

func (h *Host) RunCmd(cmd string, timeout time.Duration) (string, error) {
	cli := &util.SSHClient{
		IP:       h.IP,
		Username: h.Username,
		Password: h.Password,
		Port:     h.Port,
	}
	return cli.RunCmd(cmd, timeout)
}

func HostAdd(host *Host) error {
	orm := GetMaster()
	affected, err := orm.Insert(host)
	if affected == 0 {
		return errors.New("insert failed, affected = 0")
	}
	return err
}

func HostUpdate(id int64, host *Host) error {
	orm := GetMaster()
	affected, err := orm.ID(id).Update(host)
	if affected == 0 {
		return errors.New("update failed, affected = 0")
	}
	return err
}

func HostOne(cons map[string]interface{}) (*Host, bool) {
	orm := GetSlave()
	host := new(Host)
	has, err := orm.Where(cons).Get(&host)
	if err != nil {
		return nil, false
	}
	return host, has
}

func HostList(cons map[string]interface{}) ([]Host, error) {
	orm := GetSlave()

	host := make([]Host, 0)
	err := orm.Where(cons).Find(&host)
	if err != nil {
		return nil, err
	}
	return host, err
}
