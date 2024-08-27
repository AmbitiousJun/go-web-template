// 日志配置
package config

import (
	"errors"
	"go_web_template/internal/util/strs"
)

// LogLevel 标记日志输出级别枚举
type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug LogLevel = "debug"
	LogLevelError LogLevel = "error"
)

// LogConfig 定义程序的日志相关配置
type LogConfig struct {
	// level 日志输出级别
	Level LogLevel `yaml:"level"`
}

func (lc *LogConfig) NumLevel(lvl LogLevel) int {
	switch lvl {
	case LogLevelError:
		return 1
	case LogLevelInfo:
		return 2
	case LogLevelDebug:
		return 3
	default:
		return 2
	}
}

func (lc *LogConfig) Check() error {
	if strs.Empty(string(lc.Level)) {
		// 默认 info 级别日志
		lc.Level = LogLevelInfo
	}
	if lc.Level != LogLevelInfo &&
		lc.Level != LogLevelDebug &&
		lc.Level != LogLevelError {
		return errors.New("log.level: 无效的配置: " + (string)(lc.Level))
	}
	return nil
}
