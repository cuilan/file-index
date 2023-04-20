package orm

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/file-index/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	SqliteDb *gorm.DB
	MysqlDb  *gorm.DB
	err      error
)

func StartSqlite(dbPath string) {
	logs.Info("start init sqlite...")
	SqliteDb, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = SqliteDb.AutoMigrate(&entity.File{})
	if err != nil {
		return
	}
	logs.Info("end init sqlite.")
}

func StartMysql(host, user, password, db, charset string, maxActive, maxIdle, port int) {
	logs.Info("connect mysql start...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&charset=%s",
		user, password, host, port, db, charset)
	logs.Info("mysql dsn: %s", dsn)
	MysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Error("connect common mysql err:" + err.Error())
		return
	}
	if err != nil {
		logs.Error("connect mysql err:" + err.Error())
	}
	logs.Info("connect mysql end.")
}
