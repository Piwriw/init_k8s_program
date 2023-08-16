package db

import "gorm.io/gorm"

// User 任务信息在 mysql 数据库中的表格结构
type User struct {
	gorm.Model
	ID   string `gorm:"index;comment:'任务的ID'"`
	Name string `gorm:"comment:'名字'"`
}
