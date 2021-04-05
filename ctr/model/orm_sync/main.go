package main

import (
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	// todo !
	engine, err = xorm.NewEngine("mysql", util.Config.Section("mysql").Key("master").String())
	if err != nil {
		panic(err)
	}

	engine.ShowSQL(true)

	// sync
	err = engine.Sync(new(model.Tree), new(entity.User), new(model.UserTree))
	if err != nil {
		panic(err)
	}
}

