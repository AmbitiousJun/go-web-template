// 在第三方库 zap 的基础上再封装的 logger
package logger

import (
	"fmt"
	"go_web_template/internal/config"
	"go_web_template/internal/constant"
	"go_web_template/internal/util/array"
	"go_web_template/internal/util/datetime"
	"log"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// dailyRefreshMutex 每日生成的第一条日志时, 都需要进行线程安全控制
var dailyRefreshMutex = sync.Mutex{}

// cacheDay 记录当前缓存中的 logger 是在哪一天生成的
var cacheDay = ""

// loggerCache 每天生成一个 logger 并进行缓存
var loggerCache *zap.Logger = nil

// 每个级别日志的存放前缀
var (
	infoPrefix  = fmt.Sprintf("%s%s%s", "logs", string(filepath.Separator), config.LogLevelInfo)
	errorPrefix = fmt.Sprintf("%s%s%s", "logs", string(filepath.Separator), config.LogLevelError)
	debugPrefix = fmt.Sprintf("%s%s%s", "logs", string(filepath.Separator), config.LogLevelDebug)
	logSuffix   = ".log"
)

// Get 获取 logger
func Get() *zap.Logger {
	// 1 判断是否是新的一天, 不是则返回旧的 logger
	currentDay := datetime.Today("yyyy-MM-dd")
	if !isNewDay(currentDay) {
		return loggerCache
	}
	// 2 初始化一个新的 logger 并更新缓存
	createNewLogger(currentDay)
	return loggerCache
}

// createNewLogger 根据 currentDay 创建一个新的 logger, 并更新缓存
func createNewLogger(currentDay string) {
	dailyRefreshMutex.Lock()
	defer dailyRefreshMutex.Unlock()
	if !isNewDay(currentDay) {
		return
	}
	cfg := initBaseConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(cfg)
	debug, info, err, console := initCore(encoder, currentDay)
	// 合并多个 core, 过滤掉 nil
	core := zapcore.NewTee(
		array.Filter(
			[]zapcore.Core{debug, info, err, console},
			func(elm zapcore.Core) bool { return elm != nil },
		)...,
	)
	loggerCache = zap.New(core, zap.AddCaller())
	cacheDay = currentDay
}

// initCore 初始化各个输出级别日志的 core
func initCore(encoder zapcore.Encoder, currentDay string) (zapcore.Core, zapcore.Core, zapcore.Core, zapcore.Core) {
	lvl := config.C.Log.Level
	numLevel := config.C.Log.NumLevel
	var debug, info, err, console zapcore.Core

	// 默认情况下初始化 info 级别的 console 日志
	consoleLevel := zap.InfoLevel
	if numLevel(lvl) >= numLevel(config.LogLevelError) {
		err = zapcore.NewCore(encoder, zapcore.AddSync(openFileByPrefix(errorPrefix, currentDay)), zap.ErrorLevel)
		consoleLevel = zap.ErrorLevel
	}
	if numLevel(lvl) >= numLevel(config.LogLevelInfo) {
		info = zapcore.NewCore(encoder, zapcore.AddSync(openFileByPrefix(infoPrefix, currentDay)), zap.InfoLevel)
		consoleLevel = zap.InfoLevel
	}
	if numLevel(lvl) >= numLevel(config.LogLevelDebug) {
		debug = zapcore.NewCore(encoder, zapcore.AddSync(openFileByPrefix(debugPrefix, currentDay)), zap.DebugLevel)
		consoleLevel = zap.DebugLevel
	}

	console = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), consoleLevel)
	return debug, info, err, console
}

// openFileByPrefix 根据日志前缀打开文件, 忽略错误
func openFileByPrefix(prefix, currentDay string) *os.File {
	pathName := fmt.Sprintf("%s%s%s%s", prefix, string(filepath.Separator), currentDay, logSuffix)
	if err := os.MkdirAll(filepath.Dir(pathName), os.ModePerm); err != nil {
		log.Println("创建日志目录失败", err)
		return nil
	}
	file, err := os.OpenFile(pathName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("打开日志文件失败", err)
		return nil
	}
	return file
}

// initBaseConfig 根据程序运行环境初始化一个基础配置
func initBaseConfig() zapcore.EncoderConfig {
	profiles := config.ActiveProfiles
	for i := len(profiles) - 1; i >= 0; i-- {
		switch profiles[i] {
		case constant.ProfileDev:
			return zap.NewDevelopmentEncoderConfig()
		case constant.ProfileProd:
			return zap.NewProductionEncoderConfig()
		case constant.ProfileTest:
			return zap.NewDevelopmentEncoderConfig()
		default:
		}
	}
	return zap.NewDevelopmentEncoderConfig()
}

// isNewDay 返回是否是新的一天
func isNewDay(currentDay string) bool {
	return currentDay != cacheDay
}
