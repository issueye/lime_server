package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func MakeUserController() UserController {
	return UserController{}
}

// GetUserInfo doc
//
//	@tags			用户
//	@Summary		获取用户信息
//	@Description	获取用户信息
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=model.User}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/admin/userInfo [get]
//	@Security		ApiKeyAuth
func (control *UserController) GetUserInfo(c *gin.Context) {
	ctl := controller.New(c)

	// 从token 中解析出用户信息
	user, err := logic.GetUser(c)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	// 查询用户角色
	fullUser, err := logic.FindUserByName(user.Username)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	// 密码置空
	fullUser.Password = ""

	ctl.SuccessData(fullUser)
}

// UpdateUserinfo doc
//
//	@tags			用户
//	@Summary		更新用户信息
//	@Description	更新用户信息
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=model.User}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/admin/updateuserinfo [post]
//	@Security		ApiKeyAuth
func (control *UserController) UpdateUserinfo(c *gin.Context) {
	ctl := controller.New(c)

	// 从token 中解析出用户信息
	user, err := logic.GetUser(c)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = ctl.BindJSON(&user)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewUserLogic().UpdateUser(&user)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdatePassword doc
//
//	@tags			用户
//	@Summary		更新用户密码
//	@Description	更新用户密码
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/admin/updatepassword [post]
//	@Security		ApiKeyAuth
func (control *UserController) UpdatePassword(c *gin.Context) {
	ctl := controller.New(c)
	condition := requests.NewUpdatePassword()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	// 从token 中解析出用户信息
	user, err := logic.GetUser(c)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewUserLogic().UpdatePassword(user, condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// GetUsers doc
//
//	@tags			用户
//	@Summary		获取用户列表
//	@Description	获取用户列表
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=model.ResPage{list=[]model.User}}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/user/list [get]
//	@Security		ApiKeyAuth
func (control *UserController) GetUsers(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewQueryUser()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	users, err := logic.NewUserLogic().ListUser(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(users)
}

// UpdateUser doc
//
//	@tags			用户
//	@Summary		更新用户信息
//	@Description	更新用户信息
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/user/update [put]
//	@Security		ApiKeyAuth
func (control *UserController) UpdateUser(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewUpdateUser()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewUserLogic().UpdateUserInfo(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteUser doc
//
//	@tags			用户
//	@Summary		删除用户
//	@Description	删除用户
//	@Produce		json
//	@Param			id		path	int	true	"用户id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/user/update [delete]
//	@Security		ApiKeyAuth
func (control *UserController) DeleteUser(c *gin.Context) {
	ctl := controller.New(c)

	id := ctl.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewUserLogic().DeleteUser(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// CreateUser doc
//
//	@tags			用户
//	@Summary		创建用户
//	@Description	创建用户
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=[]model.User}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/user/create [post]
//	@Security		ApiKeyAuth
func (control *UserController) CreateUser(c *gin.Context) {
	ctl := controller.New(c)

	data := requests.NewCreateUser()
	err := ctl.Bind(data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewUserLogic().CreateUser(data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// ResetPwd doc
//
//	@tags			用户
//	@Summary		重置密码
//	@Description	重置密码
//	@Produce		json
//	@Success		200	{object}	controller.Response{}					"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/user/resetPwd/{id} [put]
//	@Security		ApiKeyAuth
func (control *UserController) ResetPwd(c *gin.Context) {
	ctl := controller.New(c)

	id := ctl.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewUserLogic().ResetPwd(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
