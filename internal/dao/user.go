package dao

import (
	"errors"
	"fmt"
	"go_web_template/internal/logger"
	"go_web_template/internal/model/entity"
	"go_web_template/internal/util/strs"

	"go.uber.org/zap"
)

type UserDao struct {
	u *entity.User
}

func NewUserDao() *UserDao { return &UserDao{u: new(entity.User)} }

// InfoById 根据 id 查询用户信息
func (ud *UserDao) InfoById(id int) (*entity.User, error) {
	var u = new(entity.User)
	if err := InspectDbError(db.Model(u).First(u, "id", id)); err != nil {
		return nil, err
	}
	return u, nil
}

// InfoByAccount 根据账号查询用户信息
func (ud *UserDao) InfoByAccount(account string) (*entity.User, error) {
	u := &entity.User{
		Account: account,
	}
	if err := InspectDbError(db.Model(u).First(u, u)); err != nil {
		return nil, err
	}
	return u, nil
}

// InfoByActPwd 根据账号密码查询用户信息
func (ud *UserDao) InfoByActPwd(account, password string) (*entity.User, error) {
	u := &entity.User{
		Account:  account,
		Password: password,
	}
	if err := InspectDbError(db.Model(u).First(u, u)); err != nil {
		return nil, err
	}
	return u, nil
}

// Page 分页查询用户数据
func (ud *UserDao) Page(cond *entity.User, current, size int64) (*entity.Page[*entity.User], error) {
	query := db.Model(ud.u)
	if cond == nil {
		return Page[*entity.User](query, current, size)
	}
	if !strs.Empty(cond.Account) {
		query = query.Where("account like ?", "%"+cond.Account+"%")
	}
	if !strs.Empty(cond.UserName) {
		query = query.Where("user_name like ?", "%"+cond.UserName+"%")
	}
	if !strs.Empty(cond.Role) {
		query = query.Where("role = ?", cond.Role)
	}
	return Page[*entity.User](query, current, size)
}

// DelById 根据 id 删除用户信息
func (ud *UserDao) DelById(id int) error {
	res := db.Model(ud.u).Unscoped().Delete(ud.u, id)
	return InspectDbError(res, 1)
}

// Update 更新用户信息
func (ud *UserDao) Update(u *entity.User) error {
	if u == nil || u.Id == 0 {
		return fmt.Errorf("无法更新用户: %v", u)
	}
	return InspectDbError(db.Model(u).Updates(u), 1)
}

// Add 新增用户
func (ud *UserDao) Add(u *entity.User) error {
	if u == nil {
		return errors.New("未传递待新增的用户信息")
	}
	u.Id = 0
	return InspectDbError(db.Model(u).Save(u), 1)
}

// AccountExist 检查账号是否已经存在
func (ud *UserDao) AccountExist(account string) bool {
	var cnt int64
	if err := db.Model(ud.u).Where("account = ?", account).Count(&cnt).Error; err != nil {
		logger.Get().Error("数据库调用异常", zap.Error(err))
		return true
	}
	return cnt > 0
}
