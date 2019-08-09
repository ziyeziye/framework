package models

import (
	"fmt"
	"framework/pkg/config"
	"framework/pkg/logger"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var tablePrefix string
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host string
	)

	dbType = config.GetDatabase("TYPE").String()
	dbName = config.GetDatabase("NAME").String()
	user = config.GetDatabase("USER").String()
	password = config.GetDatabase("PASSWORD").String()
	host = config.GetDatabase("HOST").String()
	tablePrefix = config.GetDatabase("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		fmt.Println(err)
	}

	db.LogMode(true)
	if config.RumMode == "debug" {
		//调试模式输出在命令行
		sqlLogger := &logger.SqlLogger{}
		db.SetLogger(sqlLogger)
	} else {
		//输出到日志文件
		sqlLogger := &logger.SqlLogger{}
		db.SetLogger(sqlLogger)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}
