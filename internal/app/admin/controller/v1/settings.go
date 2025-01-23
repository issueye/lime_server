package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/requests"
	"lime/internal/common/controller"

	"github.com/gin-gonic/gin"
)

// GetSystemSetting doc
//
//	@tags			系统设置
//	@Summary		系统设置
//	@Description	系统设置
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=[]response.Settings}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/settings/system [get]
//	@Security		ApiKeyAuth
func GetSystemSettings(c *gin.Context) {
	ctl := controller.New(c)
	list := logic.GetSystemSettings()
	ctl.SuccessData(list)
}

// SetSystemSetting doc
//
//	@tags			系统设置
//	@Summary		设置系统设置
//	@Description	设置系统设置
//	@Produce		json
//	@Param			body	body		requests.SetSystemSetting	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/settings/system [put]
//	@Security		ApiKeyAuth
func SetSystemSettings(c *gin.Context) {
	ctl := controller.New(c)
	data := requests.NewSettings()
	err := ctl.Bind(&data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	logic.SetSystemSetting(data)

	ctl.Success()
}

// GetLoggerSetting doc
//
//	@tags			系统设置
//	@Summary		日志设置
//	@Description	日志设置
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=[]response.Settings}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/settings/logger [get]
//	@Security		ApiKeyAuth
func GetLoggerSettings(c *gin.Context) {
	ctl := controller.New(c)
	list := logic.GetLoggerSettings()
	ctl.SuccessData(list)
}

// SetLoggerSetting doc
//
//	@tags			系统设置
//	@Summary		设置日志设置
//	@Description	设置日志设置
//	@Produce		json
//	@Param			body	body		requests.SetLoggerSetting	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/settings/logger [put]
//	@Security		ApiKeyAuth
func SetLoggerSettings(c *gin.Context) {
	ctl := controller.New(c)
	data := requests.NewSettings()
	err := ctl.Bind(&data)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	logic.SetLoggerSetting(data)
	ctl.Success()
}
