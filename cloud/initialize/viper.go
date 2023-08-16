package initialize

import (
	"base_k8s/cloud/global"
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// DefaultConfigDir 配置文件在容器里的默认路径
const DefaultConfigDir = "./build/config/"

// SerfAgentConfigFileName 配置文件在容器里的默认名称
const SerfAgentConfigFileName = "cloud.yaml"

// ViperConfig 从配置文件中读取程序所需的参数
func ViperConfig() {
	viper.AutomaticEnv()
	viper.AddConfigPath(DefaultConfigDir)
	cfgFilePath := filepath.Join(DefaultConfigDir, SerfAgentConfigFileName)
	viper.SetConfigFile(cfgFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败:%v", err))
	}

	err = viper.Unmarshal(&global.CONFIG)
	fmt.Println(global.CONFIG)
	if err != nil {
		log.Println(err)
	}
}
