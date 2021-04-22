package main

import (
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", util.Config.Section("mysql_master").Key("master").String())
	if err != nil {
		panic(err)
	}

	engine.ShowSQL(true)

	// sync
	err = engine.Sync(
		new(model.Tree), new(model.User), new(model.UserTree),
		new(model.HostUser), new(model.Host), new(model.ICode),
	)
	if err != nil {
		panic(err)
	}
}
