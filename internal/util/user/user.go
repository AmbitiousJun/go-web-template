package user

import "go_web_template/internal/constant"

var roles = map[string]struct{}{
	constant.RoleUser:  {},
	constant.RoleAdmin: {},
	constant.RoleBan:   {},
}

// RoleLegal 判断角色是否合法
func RoleLegal(role string) bool {
	_, ok := roles[role]
	return ok
}
