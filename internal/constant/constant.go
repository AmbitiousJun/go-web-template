// 常量包
package constant

const (
	// 用户登录信息
	SessionKeyUserInfo = "userInfo"
)

// 程序运行环境
const (
	ProfileDev  = "dev"
	ProfileProd = "prod"
	ProfileTest = "test"
)

// 用户属性常量
const (
	UserAccountMinLen = 4 // 用户账号最少长度
	UserPwdMinLen     = 8 // 用户密码最少长度
)

// 角色常量
const (
	RoleUser  = "user"  // 普通用户
	RoleAdmin = "admin" // 管理员
	RoleBan   = "ban"   // 账号禁用
)
