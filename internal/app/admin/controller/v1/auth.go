package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/response"
	"lime/internal/common"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login doc
//
//	@tags			基础接口
//	@Summary		用户登录
//	@Description	用户登录
//	@Produce		json
//	@Param			data	body		requests.LoginRequest					true	"用户名、密码"
//	@Success		200		{object}	controller.Response{Data=response.Auth}	"code: 200 成功"
//	@Failure		500		{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/auth/login [post]
func Login(c *gin.Context) {
	ctl := controller.New(c)

	var userLogin requests.LoginRequest
	if err := ctl.BindJSON(&userLogin); err != nil {
		ctl.FailWithError(err)
		return
	}

	user, token, err := logic.Login(userLogin)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	res := new(response.Auth)
	res.ID = user.ID
	res.User = user.Username
	res.Token = token

	ctl.SuccessData(res)
}

// RefreshToken doc
//
//	@tags			基础接口
//	@Summary		刷新Token
//	@Description	刷新Token
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=response.Auth}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/auth/refresh [get]
//	@Security		ApiKeyAuth
func RefreshToken(c *gin.Context) {
	ctl := controller.New(c)

	tokenString := c.GetHeader("Authorization")
	token, err := common.RefreshToken(tokenString)

	if err != nil {
		ctl.FailWithError(err)
		return
	}

	res := new(response.Auth)
	userid, err := strconv.Atoi(token.UserID)
	if err != nil {
		ctl.FailWithError(err)
		return
	}
	res.ID = uint(userid)
	res.User = token.Username
	res.Token = token.Token

	ctl.SuccessData(res)
}

// Logout doc
//
//	@tags			基础接口
//	@Summary		退出登录
//	@Description	退出登录
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/auth/logout [get]
//	@Security		ApiKeyAuth
func Logout(c *gin.Context) {
	ctl := controller.New(c)
	ctl.Success()
}
