package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	// DefaultConfigDir 配置文件在容器里的默认路径
	DefaultConfigDir = "/etc/program"
	// SerfAgentConfigFileName 配置文件在容器里的默认名称
	SerfAgentConfigFileName = "config.yaml"

	ImagePath = "/root/program/face_images"

	//serf-agent pod 在创建的时候会将主机的 /root 文件夹挂载到 serf-agent pod 上，随后它会从配置文件中读取放置 AI 待处理图片的文件目录的地址
	//（该地址是由创建 agent-cloud 的时候在配置文件中写入的，并且在创建 serf-agent pod 的时候这个图片目录也被挂载到了 serf pod 中）。
	//读取了地址后 serf pod 会在其挂载的 /root/uavedge/images/task-id/nodeName 目录创建一个此图片目录的软链接，然后这个软链接地址
	//（/root/uavedge/images/task-id/nodeName）会作为 AI 容器的宿主机路径然后被挂载到容器的目录中。

	// VolumePath AI 容器在宿主机中挂载的目录以及容器中的目标目录
	VolumePath = "/root/program"
	// ImagesMountPath AI 容器的图片在容器中的挂载地址。这和 AI 的代码逻辑密切相关 TODO：后续是否考虑把它写到 cloud 的配置文件中？
	ImagesMountPath = "/root/program/images"

	ImagesMountHostPath = "/root/program/images"

	// MemUseMaxPercent 内存的最大使用率阀值
	MemUseMaxPercent = 90
	// CPUUseMaxPercent CPU 时间片的最大使用率阀值
	CPUUseMaxPercent = 90
	// BatteryMinPercent 电池的最低电量阀值
	BatteryMinPercent = 20
)

var (
	// CMViper 用于读取位于 /etc/program/config.yaml 的 viper 对象
	CMViper *viper.Viper
	// ImageVolumePath 图片在宿主机上的挂载目录(程序会通过创建一个软链接的方式创建该目录，指向的是图片真实的路径即 ImagePath)
	ImageVolumePath string

	// AIModelSaveHostPath AI 模型文件在宿主机上的保存目录
	AIModelSavePath string
	// AIModelPath AI 模型文件在容器上的保存路径（绝对路径）
	AIModelPath string

	// DbPath 数据库文件的路径
	DbPath string
	// DebugResPath debug 模式时模拟的资源路径
	DebugResPath string
	// TaskCommand 容器初始化时的指令
	Command []string
)

// AgentViperConfig 使用 viper 读取位于 DefaultConfigDir 目录下的配置文件内容，以及读取 Pod 中的环境变量
func AgentViperConfig() {
	// 自动读取环境变量
	CMViper.AutomaticEnv()
	CMViper.AddConfigPath(DefaultConfigDir)
	cfgFilePath := filepath.Join(DefaultConfigDir, SerfAgentConfigFileName)
	CMViper.SetConfigFile(cfgFilePath)

	err := CMViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败:%v", err))
	}
	NewSerfAgentConfig()

	// ImageVolumePath 是 serf-agent 创建的宿主机上 AI 待识别图片目录的软链接目录地址
	ImageVolumePath = fmt.Sprintf("/uavedge/images/%v/%v", CMViper.Get("TaskID").(string), CMViper.Get("MY_NODE_NAME").(string))
	AIModelSavePath = filepath.Join("/root/program/cache/ai-model", fmt.Sprintf("%v-%v", CMViper.Get("TaskName"), CMViper.Get("TaskID")))

	DbPath = fmt.Sprintf("/root/uavedge/cache/dbs/%v/resource.db", CMViper.Get("TaskID").(string))
	DebugResPath = fmt.Sprintf("/root/uavedge/cache/debug-res/%v/res.yaml", CMViper.Get("TaskID").(string))
}

// SerfAgentConfig 从容器的文件中读取的配置文件的结构体
type SerfAgentConfig struct {
	Debug           string            `yaml:"Debug"`
	TaskID          string            `yaml:"TaskID"`
	TaskName        string            `yaml:"TaskName"`
	ImageAddress    string            `yaml:"ImageAddress"`
	AIModelAddress  string            `yaml:"AIModelAddress"`
	TaskStatus      string            `yaml:"TaskStatus"`
	NodeNameAndIP   map[string]string `yaml:"NodeNameAndIP"`
	CloudServerAddr string            `yaml:"CloudServerAddr"`
	Command         []string          `yaml:"Command"`
}

func NewSerfAgentConfig() *SerfAgentConfig {
	err := CMViper.Unmarshal(&SerfAgentSetting)
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败:%v", err))
	}
	return SerfAgentSetting
}

func GetImagesMountPath() string {
	return fmt.Sprintf("/program/images/%v/%v", CMViper.Get("TaskID").(string), CMViper.Get("MY_NODE_NAME").(string))
}

func GetExtImagesMountPath() string {
	return fmt.Sprintf("/program/ext_images/%v/%v", CMViper.Get("TaskID").(string), CMViper.Get("MY_NODE_NAME").(string))
}

func GetExtImagesWorkDirPath() string {
	return fmt.Sprintf("/program/ext_images/%v/workdir", CMViper.Get("TaskID").(string))
}

func GetFaceImagesMountPath() string {
	return "/program/face_images"
}
