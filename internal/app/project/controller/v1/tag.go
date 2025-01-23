package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTag doc
//
//	@tags			标签管理
//	@Summary		创建标签
//	@Description	创建标签
//	@Produce		json
//	@Param			body	body		requests.CreateTag	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/tag [post]
//	@Security		ApiKeyAuth
func CreateTag(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateTag()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateTag(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateTag doc
//
//	@tags			标签管理
//	@Summary		更新标签
//	@Description	更新标签
//	@Produce		json
//	@Param			body	body		requests.UpdateTag	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/tag [put]
//	@Security		ApiKeyAuth
func UpdateTag(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateTag()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateTag(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateTagReleaseStatus doc
//
//	@tags			标签管理
//	@Summary		更新标签发布状态
//	@Description	更新标签发布状态
//	@Produce		json
//	@Param			id		path	int	true	"标签ID"
//	@Param			status	body	string	true	"状态"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/tag/release [put]
//	@Security		ApiKeyAuth
func UpdateTagReleaseStatus(c *gin.Context) {
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

	var status struct {
		Status string `json:"status"`
	}

	err = ctl.Bind(&status)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateTagReleaseStatus(uint(i), status.Status)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteTag doc
//
//	@tags			标签管理
//	@Summary		删除标签
//	@Description	删除标签
//	@Produce		json
//	@Param			id		path	int	true	"标签ID"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/tag/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteTag(c *gin.Context) {
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

	err = logic.DeleteTag(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// TagList doc
//
//	@tags			标签管理
//	@Summary		标签列表
//	@Description	标签列表
//	@Produce		json
//	@Param			body	body		requests.QueryTag	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/tag/list [post]
//	@Security		ApiKeyAuth
func TagList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryTag()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.TagList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}
