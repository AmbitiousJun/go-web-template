package businesserr

import "fmt"

// 系统预定义的错误枚举
var (
	EnumSuccess        = New(0, "请求成功")
	EnumParamsError    = New(40000, "请求参数错误")
	EnumNoLoginError   = New(40100, "未登录")
	EnumNoAuthError    = New(40101, "无权限")
	EnumNotFoundError  = New(40400, "请求数据不存在")
	EnumForbiddenError = New(40300, "禁止访问")
	EnumSystemError    = New(50000, "系统内部异常")
	EnumOperationError = New(50001, "操作失败")
)

// E 实现了 Error 接口的自定义业务错误类型
type E struct {
	Code    int
	Message string
}

func New(code int, msg string) *E {
	return &E{Code: code, Message: msg}
}

func (e *E) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}
