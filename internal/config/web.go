// web 服务配置
package config

import (
	"errors"
	"strings"
)

// WebConfig 定义 web 服务相关配置参数
type WebConfig struct {
	// 接口上下文路径
	ContextPath string `yaml:"context-path"`
	// session 过期时间
	CookieMaxAge int `yaml:"cookie-max-age"`
}

func (wc *WebConfig) Check() error {
	if wc.ContextPath = strings.TrimSpace(wc.ContextPath); wc.ContextPath == "" {
		wc.ContextPath = "/"
	}
	if wc.CookieMaxAge < 0 {
		return errors.New("无效的配置: web.cookie-max-age 配置值取值范围 [0, MaxInt32]")
	}
	return nil
}
