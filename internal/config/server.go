// server 服务器配置
package config

import (
	"errors"
)

// ServerConfig 定义服务器相关配置参数
type ServerConfig struct {
	// 监听端口
	Port int `yaml:"port"`
}

func (sc *ServerConfig) Check() error {
	if sc.Port < 1 || sc.Port > 65535 {
		return errors.New("server.port: 不合法的端口号, 需配置在区间 [1, 65535]")
	}
	return nil
}
