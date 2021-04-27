package model

import (
	"errors"
	"fmt"
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
	Router      bool      `json:"router" xorm:"notnull default false"`
	Extra       string    `json:"extra" xorm:"text notnull"`
}

type HostInfo struct {
	Id          int64       `json:"id"`
	CreateUser  *UserInfo   `json:"create_user"`
	Description string      `json:"description"`
	IP          string      `json:"ip"`
	Username    string      `json:"username"`
	Port        int         `json:"port"`
	CreateTime  time.Time   `json:"create_time"`
	UpdateTime  time.Time   `json:"update_time"`
	Router      bool        `json:"router"`
	Extra       string      `json:"extra"`
	UserList    []*UserInfo `json:"user_list"`
}

func (h *Host) GetHostInfo() *HostInfo {
	u, _ := UserOne(map[string]interface{}{"id": h.UserId})
	hul, _ := HostUserList(map[string]interface{}{"host_id": h.Id})
	ul := make([]*UserInfo, 0)
	for _, hu := range hul {
		ul = append(ul, hu.GetUser().User2UserInfo())
	}

	return &HostInfo{
		Id:          h.Id,
		CreateUser:  u.User2UserInfo(),
		Description: h.Description,
		IP:          h.IP,
		Username:    h.Username,
		Port:        h.Port,
		CreateTime:  h.CreateTime,
		UpdateTime:  h.UpdateTime,
		Router:      h.Router,
		Extra:       h.Extra,
		UserList:    ul,
	}
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

func (h *Host) GetUnusedPort() int {
	for i := 10000; i < 60000; i++ {
		_, err := h.RunCmd(fmt.Sprintf("lsof -i:%d", i), time.Second*10)
		if err != nil {
			return i
		}
	}
	return 0
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

func HostUpdateRouter(r bool) {
	GetMaster().Cols("router").Update(&Host{
		Router: r,
	})
}

func HostOne(cons map[string]interface{}) (*Host, bool) {
	orm := GetSlave()
	host := new(Host)
	has, err := orm.Where(cons).Get(host)
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

func HostDelete(host *Host) error {
	orm := GetMaster()
	affected, err := orm.Delete(host)
	if affected == 0 {
		return errors.New("delete failed, affected = 0")
	}
	return err
}
