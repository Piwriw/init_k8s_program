package global

import (
	"base_k8s/cloud/config"
	"gorm.io/gorm"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
)
