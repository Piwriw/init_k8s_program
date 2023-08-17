package main

import (
	"base_k8s/edge/cmd/agent/app"
	"base_k8s/edge/cmd/agent/app/config"
	"base_k8s/edge/pkg/utils"
	"github.com/spf13/viper"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
)

func init() {
	config.CMViper = viper.New()
	config.AgentViperConfig()

	// 创建db的文件夹
	if _, err := os.Stat(config.DbPath); err != nil {
		if err2 := os.Remove(config.DbPath); err2 != nil {
			klog.Exit("Remove old db file failed:", err2)
		}
	}
	// 本地创建sqlite目录
	err := utils.MakeDirAll(filepath.Dir(config.DbPath))
	if err != nil {
		klog.Exitf("创建 Sql 的文件夹失败，原因：%v", err)
	}

	//// Get is case-insensitive for a key.
	//if config.CMViper.Get("Debug").(bool) {
	//	utils.DebugMakeRes()
	//	err := utils.DebugResViper()
	//	if err != nil {
	//		klog.Exit(err)
	//	}
	//}
}

func main() {
	command := app.NewAgentCommand() // 主要逻辑

	defer logs.FlushLogs() // 程序退出前，将log中的缓冲 flush
	//go serf.ListenSerfEvent() // Serf 事件回调处理函数
	//go serf.ListenAndServerForSerfEvent()

	if err := command.Execute(); err != nil {
		klog.Infoln("err: ", err)
		os.Exit(1)
	}

	// 优雅关闭容器
	err := app.GracefulShutdown()
	if err != nil {
		klog.Errorf("a error has occurred in shut program, Error: ", err)
		return
	}
}
