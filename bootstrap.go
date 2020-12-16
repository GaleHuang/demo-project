package main

import (
	"fmt"
	"github.com/galehuang/demo-project/config"
	"github.com/galehuang/demo-project/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化配置文件
func init() {
	filePath := "./conf/config.yaml"
	err := config.GetConfig().LoadConfigFromProfile(filePath)
	if err != nil {
		panic("init server configuration err=" + err.Error())
		return
	}
}

// 初始化数据库连接
func init() {
	mainDB := config.GetConfig().DB.Main
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		mainDB.User, mainDB.Password, mainDB.Addr, mainDB.Port, mainDB.Name, mainDB.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("init main database connection err=" + err.Error())
		return
	}
	services.SetMainDB(db)
}
