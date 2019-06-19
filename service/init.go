package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"im_chat/model"
	"log"
)

var DbEngine *xorm.Engine

func init() {
	dirverNmae := "mysql"
	dataSourceName := "root:@(127.0.0.1:3306)/go_im_chat?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(dirverNmae, dataSourceName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//显示SQL语句
	DbEngine.ShowSQL(true)
	//数据库最大连接数
	DbEngine.SetMaxOpenConns(10)
	//自动建表
	DbEngine.Sync2(new(model.User), new(model.Community), new(model.Contact))
	fmt.Println("init database finished")

}
