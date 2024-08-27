// 鉴权函数包
package auth

import (
	"go_web_template/internal/businesserr"
	"go_web_template/internal/constant"
	"go_web_template/internal/controller/response"
	"go_web_template/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// User 包装 controller, 需要普通用户角色才能访问
func User(handler gin.HandlerFunc) gin.HandlerFunc {
	return NeedRole(handler, constant.RoleUser)
}

// Admin 包装 controller, 需要管理员角色才能访问
func Admin(handler gin.HandlerFunc) gin.HandlerFunc {
	return NeedRole(handler, constant.RoleAdmin)
}

// NeedRole 包装 controller, 实现角色认证
func NeedRole(handler gin.HandlerFunc, role string) gin.HandlerFunc {
	// 不需要权限, 直接放行
	if role = strings.TrimSpace(role); role == "" {
		return handler
	}
	return func(c *gin.Context) {
		// 1 获取用户角色
		user, err := service.User().GetLoginUser(c)
		if err != nil {
			c.JSON(http.StatusOK, response.NewByError(err))
			return
		}
		// 2 用户没有权限, 或者是被封号
		if user.Role == "" || user.Role == constant.RoleBan {
			c.JSON(http.StatusOK, response.Error(businesserr.EnumNoAuthError))
			return
		}
		// 3 需要管理员才能访问, 但不是管理员
		if role == constant.RoleAdmin && user.Role != constant.RoleAdmin {
			c.JSON(http.StatusOK, response.Error(businesserr.EnumNoAuthError))
			return
		}
		// 校验通过, 正常走业务
		handler(c)
	}
}
