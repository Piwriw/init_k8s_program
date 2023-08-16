package main

import (
	"base_k8s/cloud/initialize"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 初始化 viper 以从配置文件中加载所需配置
	initialize.ViperConfig()
	// 初始化 mysql
	initialize.GormMysql()
}
func main() {
	// 初始化路由
	r := initialize.Router()
	// 启动 web 服务
	router.StartServer(r)
}
