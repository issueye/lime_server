package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

func MakeMenuController() MenuController {
	return MenuController{}
}

// GetMenus doc
//
//	@tags			菜单
//	@Summary		获取菜单列表
//	@Description	获取菜单列表
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/list [get]
//	@Security		ApiKeyAuth
func (control *MenuController) GetMenus(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewQueryMenu()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	roles, err := logic.NewMenuLogic().ListMenu(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(roles)
}

// GetAll doc
//
//	@tags			菜单
//	@Summary		获取所有菜单
//	@Description	获取所有菜单
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/getAll [get]
//	@Security		ApiKeyAuth
func (control *MenuController) GetAll(c *gin.Context) {
	ctl := controller.New(c)

	menus, err := logic.NewMenuLogic().GetMenus()
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(menus)
}

// GetRoleMenus doc
//
//	@tags			菜单
//	@Summary		获取菜单列表
//	@Description	获取菜单列表
//	@Produce		json
//	@Param			code		path	string	true	"角色编码"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/roleMenus/{code} [get]
//	@Security		ApiKeyAuth
func (control *MenuController) GetRoleMenus(c *gin.Context) {
	ctl := controller.New(c)
	code := c.Param("code")
	if code == "" {
		ctl.Fail("code不能为空")
		return
	}

	menus, err := logic.NewMenuLogic().GetRoleMenus(code)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(menus)
}

// SaveRoleMenus doc
//
//	@tags			菜单
//	@Summary		保存角色菜单
//	@Description	保存角色菜单
//	@Produce		json
//	@Param			code		path	string	true	"角色编码"
//	@Param			menus		body	[]string	true	"菜单id列表"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/saveRoleMenus/{code} [post]
//	@Security		ApiKeyAuth
func (control *MenuController) SaveRoleMenus(c *gin.Context) {
	ctl := controller.New(c)
	code := c.Param("code")
	if code == "" {
		ctl.Fail("code不能为空")
		return
	}

	menus := make([]string, 0)
	err := c.BindJSON(&menus)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewMenuLogic().SaveRoleMenus(code, menus)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// Create doc
//
//	@tags			菜单
//	@Summary		创建菜单
//	@Description	创建菜单
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/create [post]
//	@Security		ApiKeyAuth
func (control *MenuController) Create(c *gin.Context) {
	ctl := controller.New(c)

	data := requests.NewCreateMenu()
	err := ctl.Bind(data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewMenuLogic().Create(data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// Update doc
//
//	@tags			菜单
//	@Summary		修改菜单
//	@Description	修改菜单
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/update [put]
//	@Security		ApiKeyAuth
func (control *MenuController) Update(c *gin.Context) {
	ctl := controller.New(c)

	menu := requests.NewUpdateMenu()
	err := ctl.Bind(menu)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.NewMenuLogic().Update(menu)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// Delete doc
//
//	@tags			菜单
//	@Summary		删除菜单
//	@Description	删除菜单
//	@Produce		json
//	@Param			id		path	int	true	"菜单id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/delete/{id} [delete]
//	@Security		ApiKeyAuth
func (control *MenuController) Delete(c *gin.Context) {
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

	err = logic.NewMenuLogic().Del(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
