package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeletePackage doc
//
//	@tags			打包文件管理
//	@Summary		删除打包文件
//	@Description	删除打包文件
//	@Produce		json
//	@Param			id	path		int	true	"文件ID"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/project/package/{id} [delete]
//	@Security		ApiKeyAuth
func DeletePackage(c *gin.Context) {
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

	err = logic.DeletePackage(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// PackageList doc
//
//	@tags			打包文件管理
//	@Summary		打包文件列表
//	@Description	打包文件列表
//	@Produce		json
//	@Param			body	body		requests.QueryPackage	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/project/package/list [post]
//	@Security		ApiKeyAuth
func PackageList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryPackage()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.PackageList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// DownloadPackage doc
//
//	@tags			打包文件管理
//	@Summary		下载打包文件
//	@Description	下载打包文件
//	@Produce		octet-stream
//	@Param			id	path		int	true	"文件ID"
//	@Success		200	{file}		文件流
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/project/package/download/{id} [get]
//	@Security		ApiKeyAuth
func DownloadPackage(c *gin.Context) {
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

	path, err := logic.DownloadPackage(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	c.File(path)
}
