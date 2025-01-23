package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDicts doc
//
//	@tags			字典信息管理
//	@Summary		添加字典信息信息
//	@Description	添加字典信息信息
//	@Produce		json
//	@Param			body	body		requests.CreateDicts	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana [post]
//	@Security		ApiKeyAuth
func CreateDicts(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateDicts()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateDicts(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateDicts doc
//
//	@tags			字典信息管理
//	@Summary		修改字典信息信息
//	@Description	修改字典信息信息
//	@Produce		json
//	@Param			body	body		requests.UpdateDicts	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana [put]
//	@Security		ApiKeyAuth
func UpdateDicts(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateDicts()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateDicts(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteDicts doc
//
//	@tags			字典信息管理
//	@Summary		删除字典信息信息
//	@Description	删除字典信息信息
//	@Produce		json
//	@Param			id		path	int	true	"字典信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteDicts(c *gin.Context) {
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

	err = logic.DeleteDicts(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DictsList doc
//
//	@tags			字典信息管理
//	@Summary		字典信息列表
//	@Description	字典信息列表
//	@Produce		json
//	@Param			body	body		requests.QueryDicts	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/list [post]
//	@Security		ApiKeyAuth
func DictsList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryDicts()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.DictsList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetDicts doc
//
//	@tags			字典信息管理
//	@Summary		字典信息详情
//	@Description	字典信息详情
//	@Produce		json
//	@Param			id		path	int	true	"字典信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/{id} [get]
//	@Security		ApiKeyAuth
func GetDicts(c *gin.Context) {
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

	info, err := logic.GetDicts(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// ListDetail doc
//
//	@tags			字典信息管理
//	@Summary		字典明细列表
//	@Description	字典明细列表
//	@Produce		json
//	@Param			body	body		requests.SaveDetail	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/details [post]
//	@Security		ApiKeyAuth
func ListDetail(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryDictsDetail()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.ListDetail(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// SaveDetail doc
//
//	@tags			字典信息管理
//	@Summary		保存字典明细
//	@Description	保存字典明细
//	@Produce		json
//	@Param			body	body		requests.QueryDictsDetail	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/detail [post]
//	@Security		ApiKeyAuth
func SaveDetail(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewSaveDetail()
	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.SaveDetail(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DelDetail doc
//
//	@tags			字典信息管理
//	@Summary		删除字典明细
//	@Description	删除字典明细
//	@Produce		json
//	@Param			id		path	int	true	"字典信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/detail/{id} [delete]
//	@Security		ApiKeyAuth
func DelDetail(c *gin.Context) {
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

	err = logic.DelDetail(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// GetDictDetails doc
//
//	@tags			字典信息管理
//	@Summary		获取字典明细
//	@Description	根据字典ID获取字典明细
//	@Produce		json
//	@Param			id		path	int	true	"字典信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/dict_mana/{id}/details [get]
//	@Security		ApiKeyAuth
func GetDictDetails(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	dictInfo, err := logic.GetDictsByCode(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	condition := requests.NewQueryDictsDetail()
	condition.Condition.DictCode = dictInfo.Code

	list, err := logic.ListDetail(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}
