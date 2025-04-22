package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiManageController struct{}

func MakeApiManageController() ApiManageController {
	return ApiManageController{}
}

// GetApiInfos doc
//
//	@tags			接口信息
//	@Summary		获取接口信息列表
//	@Description	获取接口信息列表
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/api_manage [get]
//	@Security		ApiKeyAuth
func (control *ApiManageController) GetApiInfos(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewQueryApiInfo()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewApiManageLogic()
	ApiInfos, err := lc.GetList(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(ApiInfos)
}

// CreateApiInfo doc
//
//	@tags			接口信息
//	@Summary		创建接口信息
//	@Description	创建接口信息
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/api_manage [post]
//	@Security		ApiKeyAuth
func (control *ApiManageController) CreateApiInfo(c *gin.Context) {
	ctl := controller.New(c)

	data := requests.NewCreateApiInfo()
	err := ctl.Bind(data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewApiManageLogic()
	err = lc.Create(data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateApiInfo doc
//
//	@tags			接口信息
//	@Summary		修改接口信息
//	@Description	修改接口信息
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/api_manage [put]
//	@Security		ApiKeyAuth
func (control *ApiManageController) UpdateApiInfo(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewUpdateApiInfo()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	lc := logic.NewApiManageLogic()
	err = lc.Update(condition)
	if err != nil {
		ctl.FailWithError(err)
	}

	ctl.Success()
}

// DeleteApiInfo doc
//
//	@tags			接口信息
//	@Summary		删除接口信息
//	@Description	删除接口信息
//	@Produce		json
//	@Param			id		path	int	true	"接口信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/api_manage [delete]
//	@Security		ApiKeyAuth
func (control *ApiManageController) DeleteApiInfo(c *gin.Context) {
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

	lc := logic.NewApiManageLogic()
	err = lc.Delete(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
