package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBranch doc
//
//	@tags			分支管理
//	@Summary		创建分支
//	@Description	创建分支
//	@Produce		json
//	@Param			body	body		requests.CreateBranch	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/branch [post]
//	@Security		ApiKeyAuth
func CreateBranch(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateBranch()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateBranch(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateBranch doc
//
//	@tags			分支管理
//	@Summary		更新分支
//	@Description	更新分支
//	@Produce		json
//	@Param			body	body		requests.UpdateBranch	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/branch [put]
//	@Security		ApiKeyAuth
func UpdateBranch(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateBranch()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateBranch(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateBranchStatus doc
//
//	@tags			分支管理
//	@Summary		更新分支状态
//	@Description	更新分支状态
//	@Produce		json
//	@Param			id		path	int	true	"分支ID"
//	@Param			status	body	string	true	"状态"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/branch/status [put]
//	@Security		ApiKeyAuth
func UpdateBranchStatus(c *gin.Context) {
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

	err = logic.UpdateBranchStatus(uint(i), status.Status)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteBranch doc
//
//	@tags			分支管理
//	@Summary		删除分支
//	@Description	删除分支
//	@Produce		json
//	@Param			id		path	int	true	"分支ID"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/branch/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteBranch(c *gin.Context) {
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

	err = logic.DeleteBranch(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// BranchList doc
//
//	@tags			分支管理
//	@Summary		分支列表
//	@Description	分支列表
//	@Produce		json
//	@Param			body	body		requests.QueryBranch	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/branch/list [post]
//	@Security		ApiKeyAuth
func BranchList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryBranch()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.BranchList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}
