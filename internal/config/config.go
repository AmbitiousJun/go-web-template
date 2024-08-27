// config 包读取程序根目录下的 yml 配置文件到程序中
package config

import (
	"errors"
	"go_web_template/internal/util/structs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	// 配置文件路径
	PathName = "config"
	// 配置文件名称, 如 config.yml config-dev.yml config-prod.yml
	BaseName = "config"
	// 配置文件扩展名
	FileExt = ".yml"
)

// Config 应用程序配置
type Config struct {
	// Server 服务器配置
	Server *ServerConfig `yaml:"server"`
	// Application 应用程序配置
	Application *ApplicationConfig `yaml:"application"`
	// Web web 服务配置
	Web *WebConfig `yaml:"web"`
	// Log 日志配置
	Log *LogConfig `yaml:"log"`
	// DB 数据库配置
	DB *DatabaseConfig `yaml:"database"`
}

// Checker 检查配置正确性的接口
type Checker interface {
	// Check 检查配置, 出现错误则返回 error
	Check() error
}

// C 是全局唯一的配置对象
var C = &Config{}

// Load 加载配置文件, 出错时返回 error
func Load() error {
	// 1 加载根配置
	rootCfg, err := loadFile(genPath(""))
	if err != nil {
		return errors.Join(errors.New("加载根配置异常"), err)
	}
	C = structs.Merge(C, rootCfg)
	// 2 加载其他配置
	if err = C.Application.Check(); err != nil {
		return err
	}
	if len(ActiveProfiles) != 0 {
		log.Println("当前程序运行环境: ", ActiveProfiles)
		for _, profile := range ActiveProfiles {
			cfg, err := loadFile(genPath(profile))
			if err != nil {
				return errors.Join(errors.New("配置【"+profile+"】读取失败"), err)
			}
			cfg.Application = nil
			C = structs.Merge(C, cfg)
		}
	}
	// 3 校验配置
	cValue := reflect.ValueOf(C).Elem()
	for i := 0; i < cValue.NumField(); i++ {
		value := cValue.Field(i)
		if checker, ok := value.Interface().(Checker); ok {
			if err := checker.Check(); err != nil {
				return errors.Join(errors.New("配置校验异常"), err)
			}
		}
	}
	return nil
}

// genPath 生成配置文件路径
// profile 不传递, 生成 config.yml
// profile 传递, 生成 config-${profile}.yml
func genPath(profile string) string {
	builder := strings.Builder{}
	builder.WriteString(PathName)
	builder.WriteByte(filepath.Separator)
	builder.WriteString(BaseName)
	if profile = strings.TrimSpace(profile); profile != "" {
		builder.WriteString("-" + profile)
	}
	builder.WriteString(FileExt)
	return builder.String()
}

// loadFile 读取指定的配置文件
func loadFile(fp string) (*Config, error) {
	// 1 读取文件
	fileBytes, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	// 2 转换配置
	c := new(Config)
	if err = yaml.Unmarshal(fileBytes, c); err != nil {
		return nil, err
	}
	return c, nil
}
