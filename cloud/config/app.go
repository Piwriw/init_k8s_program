package config

type App struct {
	Debug     bool   `mapstructure:"debug"`
	EdgeImage string `mapstructure:"edge_image"`
	Port      string `mapstructure:"port"`
}
