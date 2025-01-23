package v1

import (
	"lime/internal/app/admin/logic"
	"lime/internal/common/controller"

	"github.com/gin-gonic/gin"
)

// GetHomeCount doc
//
//	@tags			主页
//	@Summary		数据统计
//	@Description	数据统计
//	@Produce		json
//	@Success		200	{object}	controller.Response{Data=model.User}	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/admin/homeCount [get]
//	@Security		ApiKeyAuth
func GetHomeCount(c *gin.Context) {
	ctl := controller.New(c)

	data, err := logic.GetHomeCount()
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(data)
}
