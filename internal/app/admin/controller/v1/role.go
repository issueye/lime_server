package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
func GetRoles(c *gin.Context) {
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
func CreateRole(c *gin.Context) {
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
func UpdateRole(c *gin.Context) {
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
func DeleteRole(c *gin.Context) {
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
