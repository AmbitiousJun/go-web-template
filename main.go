package main

import (
	_ "go_web_template/docs"
	"go_web_template/internal/config"
	"go_web_template/internal/dao"
	"go_web_template/internal/web"
)

func main() {
	// 1 加载配置
	if err := config.Load(); err != nil {
		panic(err)
	}

	// 2 初始化数据库
	if err := dao.InitDB(); err != nil {
		panic(err)
	}

	// 3 初始化 web 服务
	if err := web.Listen(); err != nil {
		panic(err)
	}
}
