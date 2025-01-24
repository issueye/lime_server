package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProject doc
//
//	@tags			项目信息管理
//	@Summary		添加项目信息信息
//	@Description	添加项目信息信息
//	@Produce		json
//	@Param			body	body		requests.CreateProject	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project [post]
//	@Security		ApiKeyAuth
func CreateProject(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateProject()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateProject(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateProject doc
//
//	@tags			项目信息管理
//	@Summary		修改项目信息信息
//	@Description	修改项目信息信息
//	@Produce		json
//	@Param			body	body		requests.UpdateProject	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project [put]
//	@Security		ApiKeyAuth
func UpdateProject(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateProject()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateProject(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteProject doc
//
//	@tags			项目信息管理
//	@Summary		删除项目信息信息
//	@Description	删除项目信息信息
//	@Produce		json
//	@Param			id		path	int	true	"项目信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteProject(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.DeleteProject(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// ProjectList doc
//
//	@tags			项目信息管理
//	@Summary		项目信息列表
//	@Description	项目信息列表
//	@Produce		json
//	@Param			body	body		requests.QueryProject	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/list [post]
//	@Security		ApiKeyAuth
func ProjectList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryProject()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.ProjectList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetProject doc
//
//	@tags			项目信息管理
//	@Summary		项目信息详情
//	@Description	项目信息详情
//	@Produce		json
//	@Param			id		path	int	true	"项目信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/{id} [get]
//	@Security		ApiKeyAuth
func GetProject(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	info, err := logic.GetProject(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// SyncProject doc
//
//	@tags			项目信息管理
//	@Summary		同步项目信息
//	@Description	同步项目信息
//	@Produce		json
//	@Param			id		path	int	true	"项目信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/sync/{id} [get]
//	@Security		ApiKeyAuth
func SyncProject(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.SyncProject(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
