package initialize

import (
	"base_k8s/cloud/global"
	"base_k8s/cloud/model/db"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// GormMysql 启动 mysql
func GormMysql() {
	// 根据配置文件中的内容尝试启动 mysql
	var err error
	m := global.CONFIG.MySQL
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.DBName + "?" + m.Config
	log.Println(dsn)
	global.DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("MySQL 启动失败:%v", err))
	}

	// 创建数据库

	// 自动生成数据表
	err = autoGenTables(&db.User{})
	if err != nil {
		panic(fmt.Errorf("MySQL 数据表自动生成失败:%v", err))
	}

	sqlDB, _ := global.DB.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
}

// autoGenTables 给定结构体自动生成mysql数据表
func autoGenTables(dst ...interface{}) error {
	err := global.DB.AutoMigrate(dst...)
	if err != nil {
		return err
	}
	return nil
}
