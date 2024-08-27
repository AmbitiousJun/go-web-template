// application 配置, 此配置只能在根配置文件下生效
package config

import "strings"

// ActiveProfiles 是 Profiles 配置按照 “,” 进行分割后的数组
var ActiveProfiles []string

type ApplicationConfig struct {
	// 应用程序名称
	Name string `yaml:"name"`
	// 运行环境
	Profiles string `yaml:"profiles"`
}

func (ac *ApplicationConfig) Check() error {
	// 1 处理 Profiles
	ac.Profiles = strings.TrimSpace(ac.Profiles)
	ActiveProfiles = strings.Split(ac.Profiles, ",")
	return nil
}
