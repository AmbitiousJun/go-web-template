package dto

// UserRegister 用户注册请求参数
type UserRegister struct {
	Account   string `form:"account"`   // 登录账号
	Password  string `form:"password"`  // 登录密码
	CheckPass string `form:"checkPass"` // 二次输入密码
}

// UserLogin 用户登录请求参数
type UserLogin struct {
	Account  string `form:"account"`  // 登录账号
	Password string `form:"password"` // 登录密码
}

// UserCreate 创建用户需要的参数
type UserCreate struct {
	Account  string `form:"account"`  // 登录账号
	UserName string `form:"userName"` // 用户昵称
	Avatar   string `form:"avatar"`   // 用户头像
	Role     string `form:"role"`     // 用户角色
}

// UserUpdate 更新用户需要的参数
type UserUpdate struct {
	Id       uint   `form:"id"`       // 用户 id
	Profile  string `form:"profile"`  // 用户简介
	UserName string `form:"userName"` // 用户昵称
	Avatar   string `form:"avatar"`   // 用户头像
	Role     string `form:"role"`     // 用户角色
}
