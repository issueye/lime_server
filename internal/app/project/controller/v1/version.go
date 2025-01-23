package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateVersion doc
//
//	@tags			版本管理
//	@Summary		创建版本
//	@Description	创建版本
//	@Produce		json
//	@Param			body	body		requests.CreateVersion	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/version [post]
//	@Security		ApiKeyAuth
func CreateVersion(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateVersion()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateVersion(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateVersion doc
//
//	@tags			版本管理
//	@Summary		更新版本
//	@Description	更新版本
//	@Produce		json
//	@Param			body	body		requests.UpdateVersion	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/version [put]
//	@Security		ApiKeyAuth
func UpdateVersion(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateVersion()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateVersion(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateVersionBuildStatus doc
//
//	@tags			版本管理
//	@Summary		更新版本构建状态
//	@Description	更新版本构建状态
//	@Produce		json
//	@Param			id		path	int	true	"版本ID"
//	@Param			status	body	string	true	"状态"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/version/build [put]
//	@Security		ApiKeyAuth
func UpdateVersionBuildStatus(c *gin.Context) {
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

	var req struct {
		Status model.BuildStatus `json:"status"`
	}

	err = ctl.Bind(&req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateVersionBuildStatus(uint(i), req.Status)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteVersion doc
//
//	@tags			版本管理
//	@Summary		删除版本
//	@Description	删除版本
//	@Produce		json
//	@Param			id		path	int	true	"版本ID"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/version/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteVersion(c *gin.Context) {
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

	err = logic.DeleteVersion(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// VersionList doc
//
//	@tags			版本管理
//	@Summary		版本列表
//	@Description	版本列表
//	@Produce		json
//	@Param			body	body		requests.QueryVersion	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/version/list [post]
//	@Security		ApiKeyAuth
func VersionList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryVersion()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.VersionList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}
