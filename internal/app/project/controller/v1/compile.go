package v1

import (
	"lime/internal/app/project/logic"
	"lime/internal/app/project/requests"
	"lime/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Compile doc
//
// @tags 项目编译管理
// @Summary 编译项目
// @Description 编译项目
// @Produce json
// @Param body body requests.CompileRequest true "编译配置"
// @Success 200 {object} controller.Response "code: 200 成功"
// @Failure 500 {object} controller.Response "错误返回内容"
// @Router /api/v1/project/compile [post]
// @Security ApiKeyAuth
func Compile(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	// id 转换为 uint
	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	versionId := c.Query("version_id")
	if versionId == "" {
		ctl.Fail("version_id不能为空")
		return
	}

	// versionId 转换为 uint
	i2, err := strconv.Atoi(versionId)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CompileProject(uint(i), uint(i2))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// SaveCompileConfig doc
//
// @tags	项目编译管理
// @Summary	保存编译配置
// @Description	保存编译配置
// @Produce json
// @Param id path int true "项目ID"
// @Param body body requests.SaveCompileConfigRequest true "编译配置"
// @Success 200 {object} controller.Response "code: 200 成功"
func SaveCompileConfig(c *gin.Context) {
	ctl := controller.New(c)
	req := requests.NewSaveCompileConfigRequest()
	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.SaveCompileConfig(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// GetConfigByProjectId doc
//
// @tags	项目编译管理
// @Summary	获取编译配置
// @Description	获取编译配置
// @Produce json
// @Param id path int true "项目ID"
// @Success 200 {object} controller.Response "code: 200 成功"
func GetConfigByProjectId(c *gin.Context) {
	ctl := controller.New(c)
	id := c.Param("id")

	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	// id 转换为 uint
	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	config, err := logic.GetConfigByProjectId(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(config)
}
