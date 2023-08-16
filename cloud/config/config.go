package config

// Config 项目中的一些配置文件
type Config struct {
	MySQL MySQL `mapstructure:"mysql"`
	K8s   K8s   `mapstructure:"k8s"`
	App   App   `mapstructure:"app"`
}

// ConfigMapTemplate configMap 配置文件的模板
type ConfigMapTemplate struct {
	Debug          bool   `yaml:"Debug"`
	TaskID         string `yaml:"TaskID"`
	TaskName       string `yaml:"TaskName"`
	ImageAddress   string `yaml:"ImageAddress"`
	AIModelAddress string `yaml:"AIModelAddress"`
	TaskStatus     string `yaml:"TaskStatus"`
	// 待部署节点的名称以及其对应的 IP
	NodeNameAndIP map[string]string `yaml:"NodeNameAndIP"`
	// Agent-cloud 服务地址
	CloudServerAddr string `yaml:"CloudServerAddr"`
	// TaskCommand 待部署节点容器的运行指令
	Command string `yaml:"Command"`
}
