// @title           Go_Web_Template
// @version         1.0
// @description     API 接口文档, 一个适用于 Java 程序员的 Golang web 模板
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:10086
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package web

import (
	"go_web_template/internal/config"
	"go_web_template/internal/controller"
	"go_web_template/internal/controller/auth"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// initRoutes 初始化 api 路由
func initRoutes(r *gin.Engine) {
	base := r.Group(config.C.Web.ContextPath)
	base.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := base.Group("/user")
	admin := user.Group("/admin")
	admin.GET("/info/:id", auth.Admin(controller.User().FindUserById))
	admin.POST("/page/:curPage/:pageSize", auth.Admin(controller.User().GetUserPage))
	admin.POST("/create", auth.Admin(controller.User().CreateUser))
	admin.DELETE("/delete/:id", auth.Admin(controller.User().DeleteUser))
	admin.PUT("/update", auth.Admin(controller.User().UpdateUser))
	user.GET("/info", auth.User(controller.User().Info))
	user.PUT("/update", auth.User(controller.User().UpdateLoginUser))
	user.POST("/login", controller.User().Login)
	user.POST("/register", controller.User().Register)
	user.POST("/logout", auth.User(controller.User().Logout))
}
