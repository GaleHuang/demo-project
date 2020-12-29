package main

import (
	"fmt"
	"github.com/galehuang/demo-project/config"
	"github.com/galehuang/demo-project/services"
	"github.com/galehuang/demo-project/services/log"
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

// 初始化日志服务
func init() {
	logger, err := log.NewLogger(config.GetConfig().Log.FileName)
	if err != nil {
		panic("init log service err=" + err.Error())
	}
	log.SetLogger(logger)
}

// 初始化数据库连接
func init() {
	log.GLogger.Critical("init database connection...")
	mainDB := config.GetConfig().DB.Main
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		mainDB.User, mainDB.Password, mainDB.Addr, mainDB.Port, mainDB.Name, mainDB.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("init main database connection err=" + err.Error())
		return
	}
	services.SetMainDB(db)
	log.GLogger.Critical("init database connection successfully...")
}
