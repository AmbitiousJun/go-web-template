// 数据库配置
package config

// DatabaseConfig 定义数据库相关配置
type DatabaseConfig struct {
	// dsn 数据库连接信息
	Dsn string `yaml:"dsn"`
}

func (dc *DatabaseConfig) Check() error {
	return nil
}
