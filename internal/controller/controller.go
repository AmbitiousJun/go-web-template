// 定义 web api 接口
package controller

import (
	"go_web_template/internal/controller/response"
	"go_web_template/internal/util/singleton"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	User = func() *UserController { return singleton.Get(NewUserController) }
)

// json 返回 json 响应
func json[T any](c *gin.Context, resp *response.R[T]) {
	c.JSON(http.StatusOK, resp)
}
