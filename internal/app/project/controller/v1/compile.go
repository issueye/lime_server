package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"

	"github.com/gin-gonic/gin"
)

// Compile doc
//
// @tags项目编译管理
// @Summary编译项目
// @Description编译项目
// @Produce json
// @Param body body requests.CompileRequest true "编译配置"
// @Success 200 {object} controller.Response "code: 200 成功"
// @Failure 500 {object} controller.Response "错误返回内容"
// @Router /api/v1/project/compile [post]
// @Security ApiKeyAuth
func Compile(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCompileRequest()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CompileProject(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
