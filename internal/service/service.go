// service 业务层
package service

import "go_web_template/internal/util/singleton"

var (
	User = func() *UserService { return singleton.Get(NewUserService) }
)
