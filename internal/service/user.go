// 用户业务层
package service

import (
	"encoding/json"
	"go_web_template/internal/businesserr"
	"go_web_template/internal/constant"
	"go_web_template/internal/dao"
	"go_web_template/internal/logger"
	"go_web_template/internal/model/dto"
	"go_web_template/internal/model/entity"
	"go_web_template/internal/model/vo"
	"go_web_template/internal/util/encrypt"
	"go_web_template/internal/util/errors"
	"go_web_template/internal/util/strs"
	"go_web_template/internal/util/syncs"
	"go_web_template/internal/util/user"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
}

func NewUserService() *UserService { return new(UserService) }

// Register 用户注册
func (us *UserService) Register(p *dto.UserRegister) (uint, error) {
	// 1 参数校验
	if p == nil || strs.AnyEmpty(p.Account, p.Password, p.CheckPass) {
		return 0, errors.Wrap(businesserr.EnumParamsError, "参数为空")
	}
	if len(p.Account) < constant.UserAccountMinLen {
		return 0, errors.Wrap(businesserr.EnumParamsError, "用户账号过短")
	}
	if len(p.Password) < constant.UserPwdMinLen || len(p.CheckPass) < constant.UserPwdMinLen {
		return 0, errors.Wrap(businesserr.EnumParamsError, "用户密码过短")
	}
	if !strings.EqualFold(p.Password, p.CheckPass) {
		return 0, errors.Wrap(businesserr.EnumParamsError, "两次输入的密码不一致")
	}
	// 2 并发控制
	mu := syncs.Mutex(syncs.KeyUserRegister + ":" + p.Account)
	mu.Lock()
	defer mu.Unlock()
	// 3 密码加密
	encrPass := encrypt.Md5Hash(p.Password)
	// 4 重复账号判断
	if dao.User().AccountExist(p.Account) {
		return 0, errors.Wrap(businesserr.EnumParamsError, "账号已存在")
	}
	// 5 添加用户
	user := entity.User{
		Account:  p.Account,
		Password: encrPass,
	}
	if err := dao.User().Add(&user); err != nil {
		logger.Get().Error("添加新用户异常", zap.Error(err))
		return 0, businesserr.EnumSystemError
	}
	return user.Id, nil
}

// CreateUserUseDftPwd 使用默认密码创建一个新用户
func (us *UserService) CreateUserUseDftPwd(p *dto.UserCreate) (uint, error) {
	// 1 参数校验
	if p == nil || strs.Empty(p.Account) {
		return 0, businesserr.EnumParamsError
	}

	// 2 账号重复判断
	mu := syncs.Mutex(syncs.KeyUserRegister + ":" + p.Account)
	mu.Lock()
	defer mu.Unlock()
	if dao.User().AccountExist(p.Account) {
		return 0, errors.Wrap(businesserr.EnumParamsError, "账号已存在")
	}
	if !user.RoleLegal(p.Role) {
		p.Role = constant.RoleUser
	}
	// 3 新增用户
	u := &entity.User{
		Account:  p.Account,
		Password: encrypt.Md5Hash("12345678"),
		Avatar:   p.Avatar,
		UserName: p.UserName,
		Role:     p.Role,
	}
	if err := dao.User().Add(u); err != nil {
		logger.Get().Error("添加新用户异常", zap.Error(err))
		return 0, businesserr.EnumSystemError
	}
	return u.Id, nil
}

// Login 用户登录
func (us *UserService) Login(c *gin.Context, p *dto.UserLogin) (*vo.User, error) {
	// 1 校验参数
	if p == nil || strs.AnyEmpty(p.Account, p.Password) {
		return nil, errors.Wrap(businesserr.EnumParamsError, "参数为空")
	}
	if len(p.Account) < constant.UserAccountMinLen {
		return nil, errors.Wrap(businesserr.EnumParamsError, "账号错误")
	}
	if len(p.Password) < constant.UserPwdMinLen {
		return nil, errors.Wrap(businesserr.EnumParamsError, "密码错误")
	}

	// 2 根据加密后的用户账号密码查询用户
	encrPass := encrypt.Md5Hash(p.Password)
	u, err := dao.User().InfoByActPwd(p.Account, encrPass)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrap(businesserr.EnumParamsError, "账号或密码错误")
		}
		logger.Get().Error("用户登录异常 => 查询用户信息失败", zap.Error(err))
		return nil, businesserr.EnumSystemError
	}

	// 3 保存登录态，返回脱敏后的用户信息
	bytes, _ := json.Marshal(&u)
	session := sessions.Default(c)
	session.Set(constant.SessionKeyUserInfo, string(bytes))
	session.Save()
	return us.GetLoginUserVO(u), nil
}

// GetLoginUserVO 获取用户实体的脱敏信息
func (us *UserService) GetLoginUserVO(u *entity.User) *vo.User {
	if u == nil {
		return nil
	}
	return &vo.User{
		Account:  u.Account,
		UserName: u.UserName,
		Avatar:   u.Avatar,
		Profile:  u.Profile,
		Role:     u.Role,
		E: entity.E{
			Id:         u.Id,
			CreateTime: u.CreateTime,
			UpdateTime: u.UpdateTime,
		},
	}
}

// GetLoginUser 从 session 获取登录用户信息
// 获取不到返回 error
func (us *UserService) GetLoginUser(c *gin.Context) (*entity.User, error) {
	session := sessions.Default(c)
	if userJson, ok := session.Get(constant.SessionKeyUserInfo).(string); ok {
		user := new(entity.User)
		if err := json.Unmarshal([]byte(userJson), user); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, businesserr.EnumNoLoginError
}

// DeleteUserById 根据 id 删除用户
func (us *UserService) DeleteUserById(id string) error {
	// 1 参数校验
	if strs.Empty(id) {
		return businesserr.EnumParamsError
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return errors.Wrap(businesserr.EnumParamsError, "id 必须是整型")
	}
	// 2 删除
	return dao.User().DelById(idNum)
}

// UpdateUserById 根据 id 更新用户信息
func (us *UserService) UpdateUserById(p *dto.UserUpdate) error {
	// 1 参数校验
	if p == nil || p.Id <= 0 {
		return errors.Wrap(businesserr.EnumParamsError, "参数不足, 至少提供用户 id")
	}
	// 2 更新用户信息
	u := &entity.User{
		E:        entity.E{Id: p.Id},
		Profile:  p.Profile,
		UserName: p.UserName,
		Avatar:   p.Avatar,
		Role:     p.Role,
	}
	return dao.User().Update(u)
}

// FindUserById 根据 id 查找用户
func (us *UserService) FindUserById(id string) (*entity.User, error) {
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrap(businesserr.EnumParamsError, "id 参数错误")
	}
	return dao.User().InfoById(idNum)
}

// GetUserPage 分页获取用户数据
func (us *UserService) GetUserPage(cond *entity.User, curPage, pageSize string) (*entity.Page[*vo.User], error) {
	cp, err := strconv.Atoi(curPage)
	if err != nil || cp < 0 {
		return nil, errors.Wrap(businesserr.EnumParamsError, "curPage 不合法")
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil || ps < 0 {
		return nil, errors.Wrap(businesserr.EnumParamsError, "pageSize 不合法")
	}

	page, err := dao.User().Page(cond, int64(cp), int64(ps))
	if err != nil {
		logger.Get().Error("查询用户信息失败", zap.Error(err))
		return nil, errors.Wrap(businesserr.EnumSystemError, "查询失败")
	}
	resPage := &entity.Page[*vo.User]{
		Current:      page.Current,
		TotalPages:   page.TotalPages,
		Size:         page.Size,
		TotalRecords: page.TotalPages,
	}
	records := make([]*vo.User, 0)
	for _, u := range page.Records {
		records = append(records, us.GetLoginUserVO(u))
	}
	resPage.Records = records
	return resPage, nil
}

// UpdateLoginUser 更新登录用户信息
func (us *UserService) UpdateLoginUser(c *gin.Context, p *dto.UserUpdate) error {
	// 1 获取当前登录用户 id
	user, err := us.GetLoginUser(c)
	if err != nil {
		return errors.Wrap(err, "获取登录用户信息失败")
	}

	// 2 调用更新接口
	p.Id = user.Id
	return us.UpdateUserById(p)
}
