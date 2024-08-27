// User 接口
package controller

import (
	"go_web_template/internal/businesserr"
	"go_web_template/internal/controller/response"
	"go_web_template/internal/model/dto"
	"go_web_template/internal/model/entity"
	"go_web_template/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController { return new(UserController) }

// Info 获取登录用户信息
//	@Summary		获取登录用户信息
//	@Description	从 session 中获取已登录用户的信息
//	@Tags			user
//	@Produce		json
//	@Success		200	{object}	response.R[vo.User]
//	@Failure		400	{object}	response.R[string]
//	@Router			/user/info [get]
func (u *UserController) Info(c *gin.Context) {
	user, err := service.User().GetLoginUser(c)
	if err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.OkWithData(service.User().GetLoginUserVO(user)))
}

// Register 用户注册
//	@Summary		用户注册
//	@Description	接收用户账号密码进行注册, 注册成功返回用户的 id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userRegisterDTO	body		dto.UserRegister	true	"用户注册参数"
//	@Success		200				{object}	response.R[uint]
//	@Failure		400				{object}	response.R[string]
//	@Router			/user/register [post]
func (u *UserController) Register(c *gin.Context) {
	params := new(dto.UserRegister)
	if err := c.ShouldBindJSON(params); err != nil {
		json(c, response.Error(businesserr.EnumParamsError))
		return
	}
	if id, err := service.User().Register(params); err != nil {
		json(c, response.NewByError(err))
	} else {
		json(c, response.OkWithData(id))
	}
}

// Login 用户登录
//	@Summary		用户登录
//	@Description	接收用户账号密码进行登录, 并保存登录态
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userLoginDTO	body		dto.UserLogin	true	"用户登录参数"
//	@Success		200				{object}	response.R[vo.User]
//	@Failure		400				{object}	response.R[string]
//	@Router			/user/login [post]
func (u *UserController) Login(c *gin.Context) {
	params := new(dto.UserLogin)
	if err := c.ShouldBindJSON(params); err != nil {
		json(c, response.Error(businesserr.EnumParamsError))
		return
	}
	uvo, err := service.User().Login(c, params)
	if err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.OkWithData(uvo))
}

// Logout 注销
//	@Summary		注销
//	@Description	用户注销登录
//	@Tags			user
//	@Produce		json
//	@Success		200	{object}	response.R[string]
//	@Router			/user/logout [post]
func (u *UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	json(c, response.OkWithData("已注销"))
}

// CreateUser 管理员接口, 使用默认密码 12345678 创建一个用户, 返回用户 id
//	@Summary		创建新用户
//	@Description	管理员接口, 使用默认密码 12345678 创建一个用户, 返回用户 id
//	@Tags			user,admin
//	@Accept			json
//	@Produce		json
//	@Param			userCreateDTO	body		dto.UserCreate	true	"创建用户参数"
//	@Success		200				{object}	response.R[uint]
//	@Failure		400				{object}	response.R[string]
//	@Router			/user/admin/create [post]
func (u *UserController) CreateUser(c *gin.Context) {
	params := new(dto.UserCreate)
	if err := c.ShouldBindJSON(params); err != nil {
		json(c, response.Error(businesserr.EnumParamsError))
		return
	}
	userId, err := service.User().CreateUserUseDftPwd(params)
	if err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.OkWithData(userId))
}

// DeleteUser 管理员接口, 删除用户
//	@Summary		删除用户
//	@Description	管理员接口, 删除用户
//	@Tags			user,admin
//	@Produce		json
//	@Param			id	path		string	true	"要删除的用户 id"
//	@Success		200	{object}	response.R[string]
//	@Failure		400	{object}	response.R[string]
//	@Router			/user/admin/delete/{id} [delete]
func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.User().DeleteUserById(id); err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.Ok())
}

// UpdateUser 管理员接口, 更新用户信息
//	@Summary		更新用户信息
//	@Description	管理员接口, 更新用户信息
//	@Tags			user,admin
//	@Accept			json
//	@Produce		json
//	@Param			userUpdateDTO	body		dto.UserUpdate	true	"用户更新参数"
//	@Success		200				{object}	response.R[string]
//	@Failure		400				{object}	response.R[string]
//	@Router			/user/admin/update [put]
func (u *UserController) UpdateUser(c *gin.Context) {
	param := new(dto.UserUpdate)
	if err := c.ShouldBindJSON(param); err != nil {
		json(c, response.Error(businesserr.EnumParamsError))
		return
	}
	if err := service.User().UpdateUserById(param); err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.Ok())
}

// UpdateLoginUser 更新登录用户自己的信息
//	@Summary		更新用户信息
//	@Description	更新登录用户自己的信息
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userUpdateDTO	body		dto.UserUpdate	true	"用户更新参数"
//	@Success		200				{object}	response.R[string]
//	@Failure		400				{object}	response.R[string]
//	@Router			/user/update [put]
func (u *UserController) UpdateLoginUser(c *gin.Context) {
	param := new(dto.UserUpdate)
	if err := c.ShouldBindJSON(param); err != nil {
		json(c, response.Error(businesserr.EnumParamsError))
		return
	}
	param.Id, param.Role = 0, ""
	if err := service.User().UpdateLoginUser(c, param); err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.Ok())
}

// FindUserById 管理员接口, 根据 id 获取用户
//	@Summary		查询用户信息
//	@Description	管理员接口, 根据 id 获取用户
//	@Tags			user,admin
//	@Produce		json
//	@Param			id	path		string	true	"要查询的用户 id"
//	@Success		200	{object}	response.R[vo.User]
//	@Failure		400	{object}	response.R[string]
//	@Router			/user/admin/info/{id} [get]
func (u *UserController) FindUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := service.User().FindUserById(id)
	if err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.OkWithData(service.User().GetLoginUserVO(user)))
}

// GetUserPage 管理员接口, 分页获取用户列表
//	@Summary		查询用户列表
//	@Description	管理员接口, 分页获取用户列表
//	@Tags			user,admin
//	@Accept			json
//	@Produce		json
//	@Param			user		body		entity.User	false	"用户查询条件"
//	@Param			curPage		path		int			true	"当前页"
//	@Param			pageSize	path		int			true	"每页大小"
//	@Success		200			{object}	response.R[entity.Page[vo.User]]
//	@Failure		400			{object}	response.R[string]
//	@Router			/user/admin/page/{curPage}/{pageSize} [post]
func (u *UserController) GetUserPage(c *gin.Context) {
	curPage, pageSize := c.Param("curPage"), c.Param("pageSize")
	user := new(entity.User)
	if err := c.ShouldBindJSON(user); err != nil {
		json(c, response.NewByError(err))
		return
	}
	page, err := service.User().GetUserPage(user, curPage, pageSize)
	if err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.OkWithData(page))
}
