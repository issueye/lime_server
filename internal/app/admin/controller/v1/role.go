package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func MakeRoleController() RoleController {
	return RoleController{}
}

// GetRoles doc
//
//	@tags			角色
//	@Summary		获取角色列表
//	@Description	获取角色列表
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/list [get]
//	@Security		ApiKeyAuth
func (control *RoleController) Roles(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewQueryRole()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	roles, err := lc.ListRole(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(roles)
}

// CreateRole doc
//
//	@tags			角色
//	@Summary		创建角色
//	@Description	创建角色
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/create [post]
//	@Security		ApiKeyAuth
func (control *RoleController) Create(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewCreateRole()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	err = lc.CreateRole(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateRole doc
//
//	@tags			角色
//	@Summary		修改角色
//	@Description	修改角色
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/update [put]
//	@Security		ApiKeyAuth
func (control *RoleController) Update(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewUpdateRole()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	err = lc.UpdateRole(condition)
	if err != nil {
		ctl.FailWithError(err)
	}

	ctl.Success()
}

// DeleteRole doc
//
//	@tags			角色
//	@Summary		删除角色
//	@Description	删除角色
//	@Produce		json
//	@Param			id		path	int	true	"角色id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/delete [delete]
//	@Security		ApiKeyAuth
func (control *RoleController) Delete(c *gin.Context) {
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

	lc := logic.NewRoleLogic()
	err = lc.DeleteRole(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// GetApis doc
//
//	@tags			角色
//	@Summary		获取角色接口权限
//	@Description	获取角色接口权限
//	@Produce		json
//	@Param			code		path	string	true	"角色识别码"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/getApis [post]
//	@Security		ApiKeyAuth
func (control *RoleController) GetApis(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewRoleQryApi()
	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	apis, err := lc.GetRoleApis(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(apis)
}

// GetNoHaveApis doc
//
//	@tags			角色
//	@Summary		获取角色接口权限
//	@Description	获取角色接口权限
//	@Produce		json
//	@Param			code		path	string	true	"角色识别码"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/getNoHaveApis [post]
//	@Security		ApiKeyAuth
func (control *RoleController) GetNoHaveApis(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewRoleQryApi()
	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	apis, err := lc.GetNoHaveApis(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(apis)
}

// RemoveApi doc
//
//	@tags			角色
//	@Summary		移除角色接口权限
//	@Description	移除角色接口权限
//	@Produce		json
//	@Param			data		body	requests.RemoveRoleApi	true	"移除信息"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/removeApi [put]
//	@Security		ApiKeyAuth
func (control *RoleController) RemoveApi(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewRoleApi()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	err = lc.RemoveApi(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// AddApi doc
//
//	@tags			角色
//	@Summary		添加角色接口权限
//	@Description	添加角色接口权限
//	@Produce		json
//	@Param			data		body	requests.RemoveRoleApi	true	"移除信息"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/addApi [post]
//	@Security		ApiKeyAuth
func (control *RoleController) AddApi(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewRoleApi()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewRoleLogic()
	err = lc.AddApi(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// FreshCasbin doc
//
//	@tags			角色
//	@Summary		添加角色接口权限
//	@Description	添加角色接口权限
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/role/freshCasbin [get]
//	@Security		ApiKeyAuth
func (control *RoleController) FreshCasbin(c *gin.Context) {
	ctl := controller.New(c)

	lc := logic.NewRoleLogic()
	err := lc.FreshCasbin()
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
