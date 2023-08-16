package config

// MySQL 连接到一个 MySQL 所需要的字段
type MySQL struct {
	Path         string `mapstructure:"path"`
	Config       string `mapstructure:"config"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}
