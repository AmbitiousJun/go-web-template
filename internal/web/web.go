// web 服务包
package web

import (
	"fmt"
	"go_web_template/internal/config"
	"go_web_template/internal/constant"
	"go_web_template/internal/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// profile2Mode 将本地运行环境映射为 gin 的运行模式
var profile2Mode = map[string]string{
	constant.ProfileDev:  gin.DebugMode,
	constant.ProfileProd: gin.ReleaseMode,
	constant.ProfileTest: gin.TestMode,
}

// Listen 启动 web 监听服务
func Listen() error {
	port := config.C.Server.Port
	log := logger.Get()
	log.Info("在指定端口上开启并监听 web 服务", zap.Int("port", port))
	settingGinMode()
	r := gin.Default()
	initSession(r)
	initRoutes(r)
	return r.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

// initSession 初始化 session
func initSession(r *gin.Engine) {
	store := cookie.NewStore([]byte("Ambitious"))
	store.Options(sessions.Options{
		MaxAge: config.C.Web.CookieMaxAge,
		Path:   "/",
	})
	r.Use(sessions.Sessions("mysession", store))
}

// settingGinMode 设置 Gin 的运行模式, 根据 profile 配置
func settingGinMode() {
	profiles := config.ActiveProfiles
	// 从后往前遍历 profiles, 匹配第一个有效的运行模式
	for i := len(profiles) - 1; i >= 0; i-- {
		if mode, ok := profile2Mode[profiles[i]]; ok {
			gin.SetMode(mode)
			return
		}
	}
}
