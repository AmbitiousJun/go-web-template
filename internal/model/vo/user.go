package vo

import "go_web_template/internal/model/entity"

// User 登录用户信息（脱敏）
type User struct {
	entity.E
	Account  string `json:"account"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	Profile  string `json:"profile"`
	Role     string `json:"role"`
}
