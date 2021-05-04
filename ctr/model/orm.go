package model

import (
	"math/rand"
	"time"

	"github.com/ahojcn/ecloud/ctr/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	client "github.com/influxdata/influxdb1-client/v2"
)

var dbMaster *xorm.Engine
var dbSlave []*xorm.Engine
var dbInflux client.Client

func init() {
	var err error

	// 主库添加
	masterConnStr := util.Config.Section("mysql_master").Key("master").String()
	dbMaster, err = xorm.NewEngine("mysql", masterConnStr)
	if err != nil {
		panic(err)
	}
	dbMaster.SetMaxIdleConns(10)
	dbMaster.SetMaxOpenConns(200)
	dbMaster.ShowSQL(true)
	dbMaster.ShowExecTime(true)

	// 从库添加
	slavesConnStr := util.Config.Section("mysql_slave").Keys()
	for _, k := range slavesConnStr {
		_dbSlave, err := xorm.NewEngine("mysql", k.String())
		if err != nil {
			panic(err)
		} else {
			_dbSlave.SetMaxIdleConns(10)
			_dbSlave.SetMaxOpenConns(200)
			_dbSlave.ShowSQL(true)
			_dbSlave.ShowExecTime(true)
			dbSlave = append(dbSlave, _dbSlave)
		}
	}

	dbInflux, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     util.Config.Section("influx_db").Key("addr").String(),
		Username: util.Config.Section("influx_db").Key("username").String(),
		Password: util.Config.Section("influx_db").Key("password").String(),
	})
	if err != nil {
		panic(err)
	}
}

func GetMaster() *xorm.Engine {
	return dbMaster
}

func GetSlave() *xorm.Engine {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(dbSlave))
	return dbSlave[n]
}

func GetInfluxDB() client.Client {
	return dbInflux
}
