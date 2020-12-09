package services

import "gorm.io/gorm"

var gSqlConn *gorm.DB

func SetSqlConn(sqlConn *gorm.DB)  {
	gSqlConn = sqlConn
}

func GetSqlConn() *gorm.DB {
	return gSqlConn
}
