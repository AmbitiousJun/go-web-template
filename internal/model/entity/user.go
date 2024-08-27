// 用户相关实体
package entity

// User 用户
type User struct {
	E
	Account  string `json:"account" gorm:"type:VARCHAR(64);unique;not null;comment:用户账号"`
	Password string `json:"password" gorm:"type:VARCHAR(256);not null;comment:用户密码"`
	UserName string `json:"userName" gorm:"type:VARCHAR(256);comment:用户名称"`
	Avatar   string `json:"avatar" gorm:"type:VARCHAR(512);comment:用户头像"`
	Profile  string `json:"profile" gorm:"type:TEXT;comment:用户简介"`
	Role     string `json:"role" gorm:"type:VARCHAR(16);default:user;comment:用户角色 user/admin/ban"`
}

func (User) TableName() string { return "tb_user" }
